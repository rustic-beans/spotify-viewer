package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/services"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/rustic-beans/spotify-viewer/internal/infrastructure/database"
	"github.com/rustic-beans/spotify-viewer/internal/infrastructure/graphql"
	httpLib "github.com/rustic-beans/spotify-viewer/internal/infrastructure/http"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
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

	playerStateWebsocketHandler := httpLib.NewWebsocketHandler[*models.PlayerState]()

	dbClient := database.Connect(config)
	databaseService := services.NewDatabase(dbClient)

	token, err := getAuthToken(config, databaseService)
	if err != nil {
		utils.Logger.Error("failed getting token. Falling back to auth", zap.Error(err))
	}

	spotifyClient := spotify.NewSpotify(config, token, getTokenSaveFunc(config, databaseService))
	spotifyService := services.NewSpotify(spotifyClient)

	sharedService := services.NewShared(databaseService, spotifyService, spotifyClient)

	watcherService := services.NewWatcher(sharedService, playerStateWebsocketHandler)

	graphqlServer := graphql.NewServer(sharedService, playerStateWebsocketHandler)
	e := httpLib.NewServer(graphqlServer)

	spotifyClient.SetupRoutes(e)
	spotifyClient.Authenticate()

	go watcherService.StartPlayerStateLoop()

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}

func getAuthToken(config *utils.Config, databaseService services.IDatabase) (*oauth2.Token, error) {
	if config.Spotify.TokenLocation != "database" {
		return config.ReadToken()
	}

	token, err := databaseService.GetToken(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed getting token from database: %w", err)
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
				return fmt.Errorf("failed marshalling token: %w", err)
			}

			//nolint:mnd // 0o600 is the file permission
			err = os.WriteFile(config.Spotify.TokenLocation, jsonData, 0o600)
			if err != nil {
				return fmt.Errorf("failed writing token to file: %w", err)
			}

			return nil
		}
	}

	return func(token *oauth2.Token) error {
		_, err := databaseService.UpsertToken(context.Background(), models.FromTokenToUpsertParams(token))
		if err != nil {
			return fmt.Errorf("failed upserting token: %w", err)
		}

		return nil
	}
}
