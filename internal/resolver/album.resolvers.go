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

// ExternalUrls is the resolver for the externalUrls field.
func (r *albumResolver) ExternalUrls(ctx context.Context, obj *database.Album) (models.StringMap, error) {
	return obj.ExternalUrls, nil
}

// Images is the resolver for the images field.
func (r *albumResolver) Images(ctx context.Context, obj *database.Album) ([]*database.Image, error) {
	return r.SharedService.GetAlbumImages(ctx, obj.ID)
}

// Artists is the resolver for the artists field.
func (r *albumResolver) Artists(ctx context.Context, obj *database.Album) ([]*database.Artist, error) {
	return r.SharedService.GetAlbumArtists(ctx, obj.ID)
}

// Tracks is the resolver for the tracks field.
func (r *albumResolver) Tracks(ctx context.Context, obj *database.Album) ([]*database.Track, error) {
	return r.SharedService.GetAlbumTracks(ctx, obj.ID)
}

// Albums is the resolver for the albums field.
func (r *queryResolver) Albums(ctx context.Context) ([]*database.Album, error) {
	return r.SharedService.GetAlbums(ctx)
}

// Album returns generated.AlbumResolver implementation.
func (r *Resolver) Album() generated.AlbumResolver { return &albumResolver{r} }

type albumResolver struct{ *Resolver }
