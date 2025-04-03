package spotify

import (
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/zmb3/spotify/v2"
)

func getImageURL(images []spotify.Image) string {
	if images == nil {
		return ""
	}

	largestImage := images[0]
	for i, image := range images {
		if image.Width > largestImage.Width {
			largestImage = images[i]
		}
	}

	return largestImage.URL
}

func FullArtistToParams(a *spotify.FullArtist) *models.CreateArtistParams {
	return &models.CreateArtistParams{
		ExternalUrls: a.ExternalURLs,
		Href:         a.Endpoint,
		ID:           string(a.ID),
		Name:         a.Name,
		Uri:          string(a.URI),
		Genres:       a.Genres,
		ImageUrl:     getImageURL(a.Images),
	}
}

func FullAlbumToParams(a *spotify.FullAlbum) *models.CreateAlbumParams {
	albumType, _ := models.StringToAlbumType(a.AlbumType)
	releaseDatePrecision, _ := models.StringToAlbumReleaseDatePrecision(a.ReleaseDatePrecision)

	return &models.CreateAlbumParams{
		ID:                   string(a.ID),
		AlbumType:            albumType,
		TotalTracks:          int64(a.Tracks.Total),
		ExternalUrls:         a.ExternalURLs,
		Href:                 a.Endpoint,
		Name:                 a.Name,
		ReleaseDate:          a.ReleaseDate,
		ReleaseDatePrecision: releaseDatePrecision,
		Uri:                  string(a.URI),
		Genres:               a.Genres,
		ImageUrl:             getImageURL(a.Images),
	}
}

func FullTrackToParams(t *spotify.FullTrack) *models.CreateTrackParams {
	var previewURL *string
	if t.PreviewURL != "" {
		previewURL = &t.PreviewURL
	}

	return &models.CreateTrackParams{
		ID:           string(t.ID),
		AlbumID:      string(t.Album.ID),
		DurationMs:   int64(t.Duration),
		Explicit:     t.Explicit,
		ExternalUrls: t.ExternalURLs,
		Href:         t.Endpoint,
		Name:         t.Name,
		Popularity:   int64(t.Popularity),
		PreviewUrl:   previewURL,
		TrackNumber:  int64(t.TrackNumber),
		Uri:          string(t.URI),
	}
}

func FullPlaylistToParams(p *spotify.FullPlaylist) *models.CreatePlaylistParams {
	return &models.CreatePlaylistParams{
		ID:           string(p.ID),
		ExternalUrls: p.ExternalURLs,
		Href:         p.Endpoint,
		Name:         p.Name,
		Uri:          string(p.URI),
		ImageUrl:     getImageURL(p.Images),
	}
}
