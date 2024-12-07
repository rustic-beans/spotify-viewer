package services

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
)

type Shared struct {
	databaseService IDatabase
	spotifyService  *Spotify
	client          *spotify.Spotify
}

func NewShared(databaseService IDatabase, spotifyService *Spotify, client *spotify.Spotify) *Shared {
	return &Shared{databaseService: databaseService, spotifyService: spotifyService, client: client}
}

func getImageUrls(images []*models.Image) []string {
	urls := make([]string, 0, len(images))
	for _, image := range images {
		urls = append(urls, image.Url)
	}

	return urls
}

func (s *Shared) GetArtist(ctx context.Context, id string) (*models.Artist, error) {
	artist, err := s.databaseService.GetArtistById(ctx, id)
	if err != nil {
		return nil, err
	}

	if artist == nil {
		artistParams, imageParams, err := s.spotifyService.GetArtist(ctx, id)
		if err != nil {
			return nil, err
		}

		images, err := s.databaseService.CreateImages(ctx, imageParams)
		if err != nil {
			return nil, err
		}

		artist, err = s.databaseService.CreateArtist(ctx, artistParams, getImageUrls(images))
		if err != nil {
			return nil, err
		}
	}

	return artist, nil
}

func (s *Shared) GetAlbum(ctx context.Context, id string) (*models.Album, error) {
	album, err := s.databaseService.GetAlbumById(ctx, id)
	if err != nil {
		return nil, err
	}

	if album == nil {
		albumParams, imageParams, artistIDs, err := s.spotifyService.GetAlbum(ctx, id)
		if err != nil {
			return nil, err
		}

		images, err := s.databaseService.CreateImages(ctx, imageParams)

		_, err = s.checkArtists(ctx, artistIDs)
		if err != nil {
			return nil, err
		}

		album, err = s.databaseService.CreateAlbum(ctx, albumParams, getImageUrls(images), artistIDs)
		if err != nil {
			return nil, err
		}
	}

	return album, nil
}

func (s *Shared) GetTrack(ctx context.Context, id string) (*models.Track, error) {
	track, err := s.databaseService.GetTrackById(ctx, id)
	if err != nil {
		return nil, err
	}

	if track == nil {
		trackParams, artistIDs, err := s.spotifyService.GetTrack(ctx, id)
		if err != nil {
			return nil, err
		}

		_, err = s.checkArtists(ctx, artistIDs)
		if err != nil {
			return nil, err
		}

		_, err = s.GetAlbum(ctx, trackParams.AlbumID)
		if err != nil {
			return nil, err
		}

		track, err = s.databaseService.CreateTrack(ctx, trackParams, artistIDs)
		if err != nil {
			return nil, err
		}
	}

	return track, nil
}

func (s *Shared) checkArtists(ctx context.Context, ids []string) ([]*models.Artist, error) {
	artists := make([]*models.Artist, 0, len(ids))
	for _, id := range ids {
		artist, err := s.GetArtist(ctx, id)
		if err != nil {
			return nil, err
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

func (s *Shared) GetPlayerState(ctx context.Context) (*models.PlayerState, error) {
	playerState, err := s.client.GetPlayerState(ctx)
	if err != nil {
		return nil, err
	}

	model := &models.PlayerState{
		ContextType: string(playerState.PlaybackContext.Type),
		ContextURI:  string(playerState.PlaybackContext.URI),

		Timestamp:  int64(playerState.Timestamp),
		ProgressMs: int64(playerState.Progress),
		IsPlaying:  playerState.Playing,
	}

	if playerState.Item == nil {
		utils.Logger.Info("No track playing")
		return model, nil
	}

	track, err := s.GetTrack(ctx, string(playerState.Item.ID))
	if err != nil {
		return nil, err
	}

	model.Track = track

	return model, nil
}

func (s *Shared) GetTracks(ctx context.Context) ([]*models.Track, error) {
	return s.databaseService.GetTracks(ctx)
}

func (s *Shared) GetAlbums(ctx context.Context) ([]*models.Album, error) {
	return s.databaseService.GetAlbums(ctx)
}

func (s *Shared) GetAlbumArtists(ctx context.Context, id string) ([]*models.Artist, error) {
	return s.databaseService.GetAlbumArtists(ctx, id)
}

func (s *Shared) GetAlbumImages(ctx context.Context, id string) ([]*models.Image, error) {
	return s.databaseService.GetAlbumImages(ctx, id)
}

func (s *Shared) GetAlbumTracks(ctx context.Context, id string) ([]*models.Track, error) {
	return s.databaseService.GetAlbumTracks(ctx, id)
}

func (s *Shared) GetArtists(ctx context.Context) ([]*models.Artist, error) {
	return s.databaseService.GetArtists(ctx)
}

func (s *Shared) GetImages(ctx context.Context) ([]*models.Image, error) {
	return s.databaseService.GetImages(ctx)
}
