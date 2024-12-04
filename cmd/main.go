package main

import (
	"fmt"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/services"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"

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

	dbClient := database.Connect(config)
	spotifyClient := spotify.NewSpotify(config)

	playerStateWebsocketHandler := httpLib.NewWebsocketHandler[*models.PlayerState]()

	databaseService := services.NewDatabase(dbClient)
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
