package services

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
)

type Shared struct {
	databaseService *Database
	spotifyService  *Spotify
	client          *spotify.Spotify
}

func NewShared(databaseService *Database, spotifyService *Spotify, client *spotify.Spotify) *Shared {
	return &Shared{databaseService: databaseService, spotifyService: spotifyService, client: client}
}

func (s *Shared) GetArtist(ctx context.Context, id string) (*models.Artist, error) {
	artist, err := s.databaseService.GetArtist(ctx, id)
	if err != nil {
		return nil, err
	}

	if artist == nil {
		artist, images, err := s.spotifyService.GetArtist(ctx, id)
		if err != nil {
			return nil, err
		}

		artist, err = s.databaseService.SaveArtist(ctx, artist, images)
		if err != nil {
			return nil, err
		}
	}

	return artist, nil
}

func (s *Shared) GetAlbum(ctx context.Context, id string) (*models.Album, error) {
	album, err := s.databaseService.GetAlbum(ctx, id)
	if err != nil {
		return nil, err
	}

	if album == nil {
		album, images, artistIDs, err := s.spotifyService.GetAlbum(ctx, id)
		if err != nil {
			return nil, err
		}

		_, err = s.checkArtists(ctx, artistIDs)
		if err != nil {
			return nil, err
		}

		album, err = s.databaseService.SaveAlbum(ctx, album, images, artistIDs)
		if err != nil {
			return nil, err
		}
	}

	return album, nil
}

func (s *Shared) GetTrack(ctx context.Context, id string) (*models.Track, error) {
	track, err := s.databaseService.GetTrack(ctx, id)
	if err != nil {
		return nil, err
	}

	if track == nil {
		track, artistIDs, err := s.spotifyService.GetTrack(ctx, id)
		if err != nil {
			return nil, err
		}

		_, err = s.checkArtists(ctx, artistIDs)
		if err != nil {
			return nil, err
		}

		_, err = s.GetAlbum(ctx, track.AlbumID)
		if err != nil {
			return nil, err
		}

		track, err = s.databaseService.SaveTrack(ctx, track, artistIDs)
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

	track, err := s.GetTrack(ctx, string(playerState.Item.ID))
	if err != nil {
		return nil, err
	}

	model := &models.PlayerState{
		ContextType: string(playerState.PlaybackContext.Type),
		ContextURI:  string(playerState.PlaybackContext.URI),

		Timestamp:  int64(playerState.Timestamp),
		ProgressMs: int64(playerState.Progress),
		IsPlaying:  playerState.Playing,

		Track: track,
	}

	return model, nil
}

func (s *Shared) GetTracks(ctx context.Context) ([]*models.Track, error) {
	return s.databaseService.GetTracks(ctx)
}

func (s *Shared) GetAlbums(ctx context.Context) ([]*models.Album, error) {
	return s.databaseService.GetAlbums(ctx)
}

func (s *Shared) GetArtists(ctx context.Context) ([]*models.Artist, error) {
	return s.databaseService.GetArtists(ctx)
}

func (s *Shared) GetImages(ctx context.Context) ([]*models.Image, error) {
	return s.databaseService.GetImages(ctx)
}
