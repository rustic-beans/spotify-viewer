package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/ent"
	"github.com/rustic-beans/spotify-viewer/generated"
	"github.com/rustic-beans/spotify-viewer/internal/models"
	lspotify "github.com/rustic-beans/spotify-viewer/internal/spotify"
	spotify "github.com/zmb3/spotify/v2"
)

// DiscNumber is the resolver for the disc_number field.
func (r *fullTrackResolver) DiscNumber(ctx context.Context, obj *spotify.FullTrack) (int, error) {
	return int(obj.DiscNumber), nil
}

// DurationMs is the resolver for the duration_ms field.
func (r *fullTrackResolver) DurationMs(ctx context.Context, obj *spotify.FullTrack) (int, error) {
	return int(obj.Duration), nil
}

// ExternalIds is the resolver for the external_ids field.
func (r *fullTrackResolver) ExternalIds(ctx context.Context, obj *spotify.FullTrack) (models.StringMap, error) {
	return obj.ExternalURLs, nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *fullTrackResolver) ExternalUrls(ctx context.Context, obj *spotify.FullTrack) (models.StringMap, error) {
	return obj.ExternalURLs, nil
}

// Href is the resolver for the href field.
func (r *fullTrackResolver) Href(ctx context.Context, obj *spotify.FullTrack) (string, error) {
	return obj.Endpoint, nil
}

// ID is the resolver for the id field.
func (r *fullTrackResolver) ID(ctx context.Context, obj *spotify.FullTrack) (string, error) {
	return string(obj.ID), nil
}

// Popularity is the resolver for the popularity field.
func (r *fullTrackResolver) Popularity(ctx context.Context, obj *spotify.FullTrack) (int, error) {
	return int(obj.Popularity), nil
}

// TrackNumber is the resolver for the track_number field.
func (r *fullTrackResolver) TrackNumber(ctx context.Context, obj *spotify.FullTrack) (int, error) {
	return int(obj.TrackNumber), nil
}

// URI is the resolver for the uri field.
func (r *fullTrackResolver) URI(ctx context.Context, obj *spotify.FullTrack) (string, error) {
	return string(obj.URI), nil
}

// Href is the resolver for the href field.
func (r *playbackContextResolver) Href(ctx context.Context, obj *spotify.PlaybackContext) (string, error) {
	return obj.Endpoint, nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *playbackContextResolver) ExternalUrls(ctx context.Context, obj *spotify.PlaybackContext) (models.StringMap, error) {
	return obj.ExternalURLs, nil
}

// URI is the resolver for the uri field.
func (r *playbackContextResolver) URI(ctx context.Context, obj *spotify.PlaybackContext) (string, error) {
	return string(obj.URI), nil
}

// ID is the resolver for the id field.
func (r *playerDeviceResolver) ID(ctx context.Context, obj *spotify.PlayerDevice) (string, error) {
	return string(obj.ID), nil
}

// IsActive is the resolver for the is_active field.
func (r *playerDeviceResolver) IsActive(ctx context.Context, obj *spotify.PlayerDevice) (bool, error) {
	return obj.Active, nil
}

// IsRestricted is the resolver for the is_restricted field.
func (r *playerDeviceResolver) IsRestricted(ctx context.Context, obj *spotify.PlayerDevice) (bool, error) {
	return obj.Restricted, nil
}

// VolumePercent is the resolver for the volume_percent field.
func (r *playerDeviceResolver) VolumePercent(ctx context.Context, obj *spotify.PlayerDevice) (int, error) {
	return int(obj.Volume), nil
}

// Context is the resolver for the context field.
func (r *playerStateResolver) Context(ctx context.Context, obj *spotify.PlayerState) (*spotify.PlaybackContext, error) {
	return &obj.PlaybackContext, nil
}

// ProgressMs is the resolver for the progress_ms field.
func (r *playerStateResolver) ProgressMs(ctx context.Context, obj *spotify.PlayerState) (int, error) {
	return int(obj.Progress), nil
}

// IsPlaying is the resolver for the is_playing field.
func (r *playerStateResolver) IsPlaying(ctx context.Context, obj *spotify.PlayerState) (bool, error) {
	return obj.Playing, nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *simpleAlbumResolver) ExternalUrls(ctx context.Context, obj *spotify.SimpleAlbum) (models.StringMap, error) {
	return obj.ExternalURLs, nil
}

// Href is the resolver for the href field.
func (r *simpleAlbumResolver) Href(ctx context.Context, obj *spotify.SimpleAlbum) (string, error) {
	return obj.Endpoint, nil
}

// ID is the resolver for the id field.
func (r *simpleAlbumResolver) ID(ctx context.Context, obj *spotify.SimpleAlbum) (string, error) {
	return string(obj.ID), nil
}

// Images is the resolver for the images field.
func (r *simpleAlbumResolver) Images(ctx context.Context, obj *spotify.SimpleAlbum) ([]*ent.Image, error) {
	models := make([]*models.Image, len(obj.Images))
	for i, image := range obj.Images {
		models[i] = lspotify.ImageToModel(&image)
	}

	return models, nil
}

// URI is the resolver for the uri field.
func (r *simpleAlbumResolver) URI(ctx context.Context, obj *spotify.SimpleAlbum) (string, error) {
	return string(obj.URI), nil
}

// ExternalUrls is the resolver for the external_urls field.
func (r *simpleArtistResolver) ExternalUrls(ctx context.Context, obj *spotify.SimpleArtist) (models.StringMap, error) {
	return obj.ExternalURLs, nil
}

// Href is the resolver for the href field.
func (r *simpleArtistResolver) Href(ctx context.Context, obj *spotify.SimpleArtist) (string, error) {
	return obj.Endpoint, nil
}

// ID is the resolver for the id field.
func (r *simpleArtistResolver) ID(ctx context.Context, obj *spotify.SimpleArtist) (string, error) {
	return string(obj.ID), nil
}

// URI is the resolver for the uri field.
func (r *simpleArtistResolver) URI(ctx context.Context, obj *spotify.SimpleArtist) (string, error) {
	return string(obj.URI), nil
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