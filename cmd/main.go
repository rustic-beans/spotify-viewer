package main

import (
	"fmt"

	"github.com/rustic-beans/spotify-viewer/utils"
	"go.uber.org/zap"

	"github.com/rustic-beans/spotify-viewer/lib/infrastructure/database"
	"github.com/rustic-beans/spotify-viewer/lib/infrastructure/graphql"
	httpLib "github.com/rustic-beans/spotify-viewer/lib/infrastructure/http"
	"github.com/rustic-beans/spotify-viewer/lib/spotify"
)

const (
	state = "state" // TODO: unique state string to identify the session, should be random
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

	graphqlServer := graphql.NewServer(spotifyClient)
	e := httpLib.NewServer(graphqlServer)

	spotifyClient.SetupRoutes(e)
	spotifyClient.Authenticate()

	go spotify.PlayerStateLoop(spotifyClient, dbClient)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}
