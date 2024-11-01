package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/generated"
	graphql1 "github.com/rustic-beans/spotify-viewer/lib/graphql"
	spotify "github.com/zmb3/spotify/v2"
)

// ExternalIds is the resolver for the external_ids field.
func (r *fullTrackResolver) ExternalIds(ctx context.Context, obj *spotify.FullTrack) (graphql1.StringMap, error) {
	return obj.ExternalURLs, nil
}

// DurationMs is the resolver for the duration_ms field.
func (r *fullTrackResolver) DurationMs(ctx context.Context, obj *spotify.FullTrack) (*int, error) {
	return &obj.Duration, nil
}

// Href is the resolver for the href field.
func (r *fullTrackResolver) Href(ctx context.Context, obj *spotify.FullTrack) (*string, error) {
	return getPointerString(obj.Endpoint), nil
}

// ID is the resolver for the id field.
func (r *fullTrackResolver) ID(ctx context.Context, obj *spotify.FullTrack) (*string, error) {
	return getPointerString(string(obj.ID)), nil
}

// URI is the resolver for the uri field.
func (r *fullTrackResolver) URI(ctx context.Context, obj *spotify.FullTrack) (*string, error) {
	return getPointerString(string(obj.URI)), nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *fullTrackResolver) ExternalUrls(ctx context.Context, obj *spotify.FullTrack) (graphql1.StringMap, error) {
	return obj.ExternalURLs, nil
}

// Href is the resolver for the href field.
func (r *playbackContextResolver) Href(ctx context.Context, obj *spotify.PlaybackContext) (*string, error) {
	return getPointerString(obj.Endpoint), nil
}

// URI is the resolver for the uri field.
func (r *playbackContextResolver) URI(ctx context.Context, obj *spotify.PlaybackContext) (*string, error) {
	return getPointerString(string(obj.URI)), nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *playbackContextResolver) ExternalUrls(ctx context.Context, obj *spotify.PlaybackContext) (graphql1.StringMap, error) {
	return obj.ExternalURLs, nil
}

// ID is the resolver for the id field.
func (r *playerDeviceResolver) ID(ctx context.Context, obj *spotify.PlayerDevice) (*string, error) {
	return getPointerString(string(obj.ID)), nil
}

// IsActive is the resolver for the is_active field.
func (r *playerDeviceResolver) IsActive(ctx context.Context, obj *spotify.PlayerDevice) (*bool, error) {
	return &obj.Active, nil
}

// IsRestricted is the resolver for the is_restricted field.
func (r *playerDeviceResolver) IsRestricted(ctx context.Context, obj *spotify.PlayerDevice) (*bool, error) {
	return &obj.Restricted, nil
}

// VolumePercent is the resolver for the volume_percent field.
func (r *playerDeviceResolver) VolumePercent(ctx context.Context, obj *spotify.PlayerDevice) (*int, error) {
	return &obj.Volume, nil
}

// Context is the resolver for the context field.
func (r *playerStateResolver) Context(ctx context.Context, obj *spotify.PlayerState) (*spotify.PlaybackContext, error) {
	return &obj.PlaybackContext, nil
}

// ProgressMs is the resolver for the progress_ms field.
func (r *playerStateResolver) ProgressMs(ctx context.Context, obj *spotify.PlayerState) (*int, error) {
	return &obj.Progress, nil
}

// IsPlaying is the resolver for the is_playing field.
func (r *playerStateResolver) IsPlaying(ctx context.Context, obj *spotify.PlayerState) (*bool, error) {
	return &obj.Playing, nil
}

// Href is the resolver for the href field.
func (r *simpleAlbumResolver) Href(ctx context.Context, obj *spotify.SimpleAlbum) (*string, error) {
	return getPointerString(obj.Endpoint), nil
}

// ID is the resolver for the id field.
func (r *simpleAlbumResolver) ID(ctx context.Context, obj *spotify.SimpleAlbum) (*string, error) {
	return getPointerString(string(obj.ID)), nil
}

// URI is the resolver for the uri field.
func (r *simpleAlbumResolver) URI(ctx context.Context, obj *spotify.SimpleAlbum) (*string, error) {
	return getPointerString(string(obj.URI)), nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *simpleAlbumResolver) ExternalUrls(ctx context.Context, obj *spotify.SimpleAlbum) (graphql1.StringMap, error) {
	return obj.ExternalURLs, nil
}

// Href is the resolver for the href field.
func (r *simpleArtistResolver) Href(ctx context.Context, obj *spotify.SimpleArtist) (*string, error) {
	return getPointerString(obj.Endpoint), nil
}

// ID is the resolver for the id field.
func (r *simpleArtistResolver) ID(ctx context.Context, obj *spotify.SimpleArtist) (*string, error) {
	return getPointerString(string(obj.ID)), nil
}

// URI is the resolver for the uri field.
func (r *simpleArtistResolver) URI(ctx context.Context, obj *spotify.SimpleArtist) (*string, error) {
	return getPointerString(string(obj.URI)), nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *simpleArtistResolver) ExternalUrls(ctx context.Context, obj *spotify.SimpleArtist) (graphql1.StringMap, error) {
	return obj.ExternalURLs, nil
}

// FullTrack returns generated.FullTrackResolver implementation.
func (r *Resolver) FullTrack() generated.FullTrackResolver { return &fullTrackResolver{r} }

// PlaybackContext returns generated.PlaybackContextResolver implementation.
func (r *Resolver) PlaybackContext() generated.PlaybackContextResolver {
	return &playbackContextResolver{r}
}

// PlayerDevice returns generated.PlayerDeviceResolver implementation.
func (r *Resolver) PlayerDevice() generated.PlayerDeviceResolver { return &playerDeviceResolver{r} }

// PlayerState returns generated.PlayerStateResolver implementation.
func (r *Resolver) PlayerState() generated.PlayerStateResolver { return &playerStateResolver{r} }

// SimpleAlbum returns generated.SimpleAlbumResolver implementation.
func (r *Resolver) SimpleAlbum() generated.SimpleAlbumResolver { return &simpleAlbumResolver{r} }

// SimpleArtist returns generated.SimpleArtistResolver implementation.
func (r *Resolver) SimpleArtist() generated.SimpleArtistResolver { return &simpleArtistResolver{r} }

type fullTrackResolver struct{ *Resolver }
type playbackContextResolver struct{ *Resolver }
type playerDeviceResolver struct{ *Resolver }
type playerStateResolver struct{ *Resolver }
type simpleAlbumResolver struct{ *Resolver }
type simpleArtistResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func getPointerString(s string) *string {
	return &s
}
