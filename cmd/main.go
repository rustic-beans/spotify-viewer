package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rustic-beans/spotify-viewer/ent"
	"github.com/rustic-beans/spotify-viewer/utils"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"

	"github.com/rustic-beans/spotify-viewer/lib/infrastructure/graphql"
	httpLib "github.com/rustic-beans/spotify-viewer/lib/infrastructure/http"
	"github.com/rustic-beans/spotify-viewer/lib/spotify"
)

const (
	state = "state" // TODO: unique state string to identify the session, should be random
)

func connectDatabase(config *utils.Config) *ent.Client {
	// Create an ent.Client with the configured database
	client, err := ent.Open(config.Database.Driver, config.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(context.Background()); err != nil {
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

	dbClient := connectDatabase(config)

	spotifyClient := spotify.NewSpotify(config)

	graphqlServer := graphql.NewServer(spotifyClient)
	e := httpLib.NewServer(graphqlServer)

	spotifyClient.SetupRoutes(e)
	spotifyClient.Authenticate()

	go spotify.PlayerStateLoop(spotifyClient, dbClient)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}
