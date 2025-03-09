package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"
)

// PlayerState is the resolver for the playerState field.
func (r *queryResolver) PlayerState(ctx context.Context) (*models.PlayerState, error) {
	return r.SharedService.GetPlayerState(ctx)
}

// PlayerState is the resolver for the playerState field.
func (r *subscriptionResolver) PlayerState(ctx context.Context) (<-chan *models.PlayerState, error) {
	utils.Logger.Info("Adding PlayerState subscription", zap.Any("context", ctx))
	id, ch := r.PlayerStateWebsocketHandler.AddConnection()
	utils.Logger.Info("PlayerState subscription added", zap.Any("context", ctx))

	go func(id string) {
		<-ctx.Done()
		r.PlayerStateWebsocketHandler.RemoveConnection(id)
		utils.Logger.Info("PlayerState subscription closed", zap.Any("context", ctx))
	}(id)

	utils.Logger.Debug("Getting PlayerState")
	playerState, err := r.SharedService.GetPlayerState(ctx)
	if err != nil {
		return nil, err
	}
	utils.Logger.Debug("Got PlayerState, broadcasting", zap.Any("playerState", playerState))
	ch <- playerState

	return ch, nil
}
