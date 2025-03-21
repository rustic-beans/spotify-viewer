package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/generated"
	"github.com/rustic-beans/spotify-viewer/internal/database"
	"github.com/rustic-beans/spotify-viewer/internal/models"
)

// Tracks is the resolver for the tracks field.
func (r *queryResolver) Tracks(ctx context.Context) ([]*database.Track, error) {
	return r.SharedService.GetTracks(ctx)
}

// ExternalUrls is the resolver for the externalUrls field.
func (r *trackResolver) ExternalUrls(ctx context.Context, obj *database.Track) (models.StringMap, error) {
	return obj.ExternalUrls, nil
}

// Artists is the resolver for the artists field.
func (r *trackResolver) Artists(ctx context.Context, obj *database.Track) ([]*database.Artist, error) {
	return r.SharedService.GetTrackArtists(ctx, obj.ID)
}

// Album is the resolver for the album field.
func (r *trackResolver) Album(ctx context.Context, obj *database.Track) (*database.Album, error) {
	return r.SharedService.GetTrackAlbum(ctx, obj.ID)
}

// Track returns generated.TrackResolver implementation.
func (r *Resolver) Track() generated.TrackResolver { return &trackResolver{r} }

type trackResolver struct{ *Resolver }
