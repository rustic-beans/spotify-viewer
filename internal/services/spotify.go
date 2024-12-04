package services

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
)

type Spotify struct {
	client *spotify.Spotify
}

func NewSpotify(client *spotify.Spotify) *Spotify {
	return &Spotify{client: client}
}

func (s *Spotify) GetArtist(ctx context.Context, id string) (*models.Artist, []*models.Image, error) {
	artist, err := s.client.GetArtist(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	return spotify.FullArtistToModel(artist), spotify.ImageSliceToModelSlice(artist.Images), nil
}

// TODO: This is kind bad. Use a DTO or CreateInput from gqlgen instead
func (s *Spotify) GetAlbum(ctx context.Context, id string) (*models.Album, []*models.Image, []string, error) {
	album, err := s.client.GetAlbum(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	artistIDs := make([]string, 0, len(album.Artists))
	for _, artist := range album.Artists {
		artistIDs = append(artistIDs, string(artist.ID))
	}

	return spotify.FullAlbumToModel(album), spotify.ImageSliceToModelSlice(album.Images), artistIDs, nil
}

func (s *Spotify) GetTrack(ctx context.Context, id string) (*models.Track, []string, error) {
	track, err := s.client.GetTrack(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	artistIDs := make([]string, 0, len(track.Artists))
	for _, artist := range track.Artists {
		artistIDs = append(artistIDs, string(artist.ID))
	}

	return spotify.FullTrackToModel(track), artistIDs, nil
}
