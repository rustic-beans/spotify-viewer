package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/albe2669/spotify-viewer/ent"
	"github.com/albe2669/spotify-viewer/generated"
	"github.com/albe2669/spotify-viewer/resolver"
	"github.com/albe2669/spotify-viewer/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"

	"github.com/albe2669/spotify-viewer/lib/spotify"
	spotifyLib "github.com/albe2669/spotify-viewer/lib/spotify"
)

const (
	QueryPath      = "/query"
	PlaygroundPath = "/playground"
	state          = "state" // TODO: unique state string to identify the session, should be random
)

type PlayerState struct {
	Progress int `json:"progress_ms"`
}

func configureLogger(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			utils.Logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
}

func graphqlServer() *handler.Server {
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{}},
		),
	)

	return server
}

func getPlayer(sa *spotify.Spotify) echo.HandlerFunc {
	return func(c echo.Context) error {
		player, err := sa.GetPlayerState(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, player)
	}
}

func playerStateLoop(sa *spotify.Spotify) {
	var c echo.Context

	//TODO: start a loop here to
	for {

	player, err := sa.GetPlayerState(c.Request().Context())
	if err != nil {
		panic(err.Error())
	}

	utils.Logger.Info(string(rune(player.Progress)))
	time.Sleep(5 * time.Second)
}

}

func httpServer(graphqlServer *handler.Server) *echo.Echo {
	e := echo.New()

	configureLogger(e)

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST(QueryPath, echo.WrapHandler(graphqlServer))
	e.GET(PlaygroundPath, func(c echo.Context) error {
		playground.Handler("GraphQL", QueryPath).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return e
}

func connectDatabase() *ent.Client {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

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

  log.Println("hi")
	dbClient := connectDatabase()
	tracks, _ := dbClient.Track.Query().All(context.Background())
	log.Println("Tracks:", tracks)

	graphqlServer := graphqlServer()
	e := httpServer(graphqlServer)

	spotify := spotifyLib.NewSpotify(config)
	spotify.SetupRoutes(e)
	spotify.Authenticate()

	e.GET("/player", getPlayer(spotify))

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}
