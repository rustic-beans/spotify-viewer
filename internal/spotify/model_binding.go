package spotify

import (
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/zmb3/spotify/v2"
)

func mapToStringMap(m map[string]string) *models.StringMap {
	if m == nil {
		return nil
	}

	sm := make(models.StringMap)
	for k, v := range m {
		sm[k] = v
	}

	return &sm
}

func ImageToModel(i *spotify.Image) *models.Image {
	return &models.Image{
		ID:     i.URL,
		URL:    i.URL,
		Width:  int(i.Width),
		Height: int(i.Height),
	}
}

func ImageSliceToModelSlice(images []spotify.Image) []*models.Image {
	if images == nil {
		return nil
	}

	models := make([]*models.Image, len(images))
	for i, image := range images {
		models[i] = ImageToModel(&image)
	}

	return models
}

func FullArtistToModel(a *spotify.FullArtist) *models.Artist {
	return &models.Artist{
		ExternalUrls: mapToStringMap(a.ExternalURLs),
		Href:         a.Endpoint,
		ID:           string(a.ID),
		Name:         a.Name,
		URI:          string(a.URI),
		Genres:       a.Genres,
	}
}

func FullAlbumToModel(a *spotify.FullAlbum) *models.Album {
	albumType, _ := models.StringToAlbumType(a.AlbumType)
	releaseDatePrecision, _ := models.StringToAlbumReleaseDatePrecision(a.ReleaseDatePrecision)

	return &models.Album{
		ID:                   string(a.ID),
		AlbumType:            albumType,
		TotalTracks:          int(a.Tracks.Total),
		ExternalUrls:         mapToStringMap(a.ExternalURLs),
		Href:                 a.Endpoint,
		Name:                 a.Name,
		ReleaseDate:          a.ReleaseDate,
		ReleaseDatePrecision: releaseDatePrecision,
		URI:                  string(a.URI),
		Genres:               a.Genres,
	}
}

func FullTrackToModel(t *spotify.FullTrack) *models.Track {
	return &models.Track{
		ID:           string(t.ID),
		AlbumID:      string(t.Album.ID),
		DiscNumber:   int(t.DiscNumber),
		DurationMs:   int(t.Duration),
		Explicit:     t.Explicit,
		ExternalUrls: mapToStringMap(t.ExternalURLs),
		Href:         t.Endpoint,
		Name:         t.Name,
		Popularity:   int(t.Popularity),
		PreviewURL:   t.PreviewURL,
		TrackNumber:  int(t.TrackNumber),
		URI:          string(t.URI),
	}
}
