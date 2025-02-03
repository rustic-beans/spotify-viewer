package spotify

import (
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/zmb3/spotify/v2"
)

func ImageToParams(i *spotify.Image) *models.CreateImageParams {
	return &models.CreateImageParams{
		Url:    i.URL,
		Width:  int64(i.Width),
		Height: int64(i.Height),
	}
}

func ImageSliceToModelParams(images []spotify.Image) []*models.CreateImageParams {
	if images == nil {
		return nil
	}

	imgModels := make([]*models.CreateImageParams, len(images))
	for i, image := range images {
		imgModels[i] = ImageToParams(&image)
	}

	return imgModels
}

func FullArtistToParams(a *spotify.FullArtist) *models.CreateArtistParams {
	return &models.CreateArtistParams{
		ExternalUrls: a.ExternalURLs,
		Href:         a.Endpoint,
		ID:           string(a.ID),
		Name:         a.Name,
		Uri:          string(a.URI),
		Genres:       a.Genres,
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
	}
}
