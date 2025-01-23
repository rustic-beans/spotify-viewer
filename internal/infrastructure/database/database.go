package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"
)

func Connect(config *utils.Config) *pgxpool.Pool {
	client, err := pgxpool.New(context.Background(), config.Database.Source)
	if err != nil {
		utils.Logger.Fatal("failed opening connection to database: %v", zap.Error(err))
	}

	return client
}
