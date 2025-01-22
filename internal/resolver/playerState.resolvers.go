package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"

	"github.com/rustic-beans/spotify-viewer/internal/models"
)

// PlayerState is the resolver for the playerState field.
func (r *queryResolver) PlayerState(ctx context.Context) (*models.PlayerState, error) {
	return r.SharedService.GetPlayerState(ctx)
}

// PlayerState is the resolver for the playerState field.
func (r *subscriptionResolver) PlayerState(ctx context.Context) (<-chan *models.PlayerState, error) {
	ch := make(chan *models.PlayerState, 1)
	id := r.PlayerStateWebsocketHandler.AddConnection(ch)

	go func() {
		<-ctx.Done()
		r.PlayerStateWebsocketHandler.RemoveConnection(id)
		close(ch)
	}()
	state, err := r.SharedService.GetPlayerState(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get player state: %w", err)
	}

	ch <- state

	return ch, nil
}
