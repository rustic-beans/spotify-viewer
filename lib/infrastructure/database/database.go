package database

import (
	"context"
	"log"

	"github.com/rustic-beans/spotify-viewer/ent"
	"github.com/rustic-beans/spotify-viewer/utils"

	_ "github.com/mattn/go-sqlite3" // sqlite driver
)

func Connect(config *utils.Config) *ent.Client {
	// Create an ent.Client with the configured database
	client, err := ent.Open(config.Database.Driver, config.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	// Run the automatic migration tool to create all schema resources.
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
