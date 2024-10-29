package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	spotify "github.com/zmb3/spotify/v2"
)

// PlayerState is the resolver for the playerState field.
func (r *queryResolver) PlayerState(ctx context.Context) (*spotify.PlayerState, error) {
	return r.SpotifyClient.GetPlayerState(ctx)
}