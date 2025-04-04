package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/services"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/rustic-beans/spotify-viewer/internal/infrastructure/database"
	"github.com/rustic-beans/spotify-viewer/internal/infrastructure/graphql"
	httpLib "github.com/rustic-beans/spotify-viewer/internal/infrastructure/http"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Logger
	defer func() {
		_ = utils.Logger.Sync()
	}()

	// Config
	config, err := utils.ReadConfig()
	if err != nil {
		utils.Logger.Fatal("failed reading config", zap.Error(err))
	}

	playerStateWebsocketHandler := httpLib.NewWebsocketHandler[*models.PlayerState](config.Server.QueueSize)

	dbClient := database.Connect(config)

	if err = migrateDB(config); err != nil {
		utils.Logger.Fatal("failed migrating database", zap.Error(err))
	}

	databaseService := services.NewDatabase(dbClient)

	token, err := getAuthToken(config, databaseService)
	if err != nil {
		utils.Logger.Fatal("failed getting token", zap.Error(err))
	}

	spotifyClient := spotify.NewSpotify(config, token, getTokenSaveFunc(config, databaseService))
	spotifyService := services.NewSpotify(spotifyClient)

	sharedService := services.NewShared(databaseService, spotifyService, spotifyClient)

	watcherService := services.NewWatcher(sharedService, playerStateWebsocketHandler)

	graphqlServer := graphql.NewServer(sharedService, playerStateWebsocketHandler)
	e := httpLib.NewServer(graphqlServer)

	spotifyClient.SetupRoutes(e)

	err = spotifyClient.Authenticate()
	if err != nil {
		utils.Logger.Fatal("failed authenticating", zap.Error(err))
	}

	ctx := context.Background()
	go watcherService.StartPlayerStateLoop(ctx)

	healthCheck(e, databaseService)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}

func getAuthToken(config *utils.Config, databaseService services.IDatabase) (*oauth2.Token, error) {
	if config.Spotify.TokenLocation != "database" {
		return config.ReadToken()
	}

	token, err := databaseService.GetToken(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed getting token from DB")
	}

	if token == nil {
		utils.Logger.Info("No token found in DB")
		return nil, nil
	}

	return models.IntoOauth2Token(token), nil
}

func getTokenSaveFunc(config *utils.Config, databaseService services.IDatabase) func(*oauth2.Token) error {
	if config.Spotify.TokenLocation != "database" {
		return func(token *oauth2.Token) error {
			jsonData, err := json.Marshal(token)
			if err != nil {
				return errors.Wrap(err, "failed marshalling token")
			}

			//nolint:mnd // 0o600 is the file permission
			err = os.WriteFile(config.Spotify.TokenLocation, jsonData, 0o600)
			if err != nil {
				return errors.Wrap(err, "failed writing token to file")
			}

			return nil
		}
	}

	return func(token *oauth2.Token) error {
		_, err := databaseService.UpsertToken(context.Background(), models.FromTokenToUpsertParams(token))
		if err != nil {
			return errors.Wrap(err, "failed upserting token")
		}

		return nil
	}
}

// TODO: move to a separate file (a little hard without getting circular imports)
func healthCheck(e *echo.Echo, databaseService services.IDatabase) {
	e.GET("/health", func(c echo.Context) error {
		err := databaseService.HealthCheck(context.Background())
		if err != nil {
			utils.Logger.Error("database health check failed", zap.Error(err))
			return c.String(http.StatusInternalServerError, "database health check failed")
		}

		return c.String(http.StatusOK, "healthy")
	})
}

func migrateDB(config *utils.Config) error {
	utils.Logger.Info("migrating database")

	m, err := migrate.New(
		"file://database/migrations",
		config.Database.Source,
	)
	if err != nil {
		return errors.Wrap(err, "failed creating driver")
	}
	defer m.Close()

	version, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return errors.Wrap(err, "failed getting version")
	}

	utils.Logger.Info("database version", zap.Uint("version", version), zap.Bool("dirty", dirty))

	err = m.Up()
	if err == nil || err.Error() == "no change" {
		utils.Logger.Info("database migrated")
		return nil
	}

	return errors.Wrap(err, "failed migrating")
}
