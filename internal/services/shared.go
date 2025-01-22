package services

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"
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

func (s *Shared) GetArtistsById(ctx context.Context, ids []string) ([]*models.Artist, []error) {
	artists, err := s.databaseService.GetArtistsById(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	if len(artists) == len(ids) {
		return artists, nil
	}

	errs := make([]error, 0, len(ids))

	for _, id := range ids {
		found := false
		for _, artist := range artists {
			if artist.ID == id {
				found = true
				break
			}
		}

		if found {
			continue
		}

		artistParams, imageParams, err := s.spotifyService.GetArtist(ctx, id)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		images, err := s.databaseService.CreateImages(ctx, imageParams)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		artist, err := s.databaseService.CreateArtist(ctx, artistParams, getImageUrls(images))
		if err != nil {
			errs = append(errs, err)
			continue
		}

		artists = append(artists, artist)
	}

	if len(errs) == 0 {
		return artists, nil
	}

	return artists, errs
}

func (s *Shared) GetAlbumsById(ctx context.Context, ids []string) ([]*models.Album, []error) {
	albums, err := s.databaseService.GetAlbumsById(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	if len(albums) == len(ids) {
		return albums, nil
	}

	errs := make([]error, 0, len(ids))

	for _, id := range ids {
		found := false
		for _, album := range albums {
			if album.ID == id {
				found = true
				break
			}
		}

		if found {
			continue
		}

		albumParams, imageParams, artistIDs, err := s.spotifyService.GetAlbum(ctx, id)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		images, err := s.databaseService.CreateImages(ctx, imageParams)

		_, errs := s.GetArtistsById(ctx, artistIDs)
		if errs != nil {
			errs = append(errs, errs...)
			continue
		}

		album, err := s.databaseService.CreateAlbum(ctx, albumParams, getImageUrls(images), artistIDs)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (s *Shared) GetTracksById(ctx context.Context, ids []string) ([]*models.Track, []error) {
	tracks, err := s.databaseService.GetTracksById(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	if len(tracks) == len(ids) {
		return tracks, nil
	}

	errs := make([]error, 0, len(ids))

	for _, id := range ids {
		found := false
		for _, track := range tracks {
			if track.ID == id {
				found = true
				break
			}
		}

		if found {
			continue
		}

		trackParams, artistIDs, err := s.spotifyService.GetTrack(ctx, id)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		_, errs := s.GetArtistsById(ctx, artistIDs)
		if errs != nil {
			errs = append(errs, errs...)
			continue
		}

		_, errs = s.GetAlbumsById(ctx, []string{trackParams.AlbumID})
		if err != nil {
			errs = append(errs, err)
			continue
		}

		track, err := s.databaseService.CreateTrack(ctx, trackParams, artistIDs)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
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

	tracks, errs := s.GetTracksById(ctx, []string{string(playerState.Item.ID)})
	if errs != nil {
		for _, err := range errs {
			utils.Logger.Error("Failed to get track", zap.Error(err))
		}

		return nil, errs[0]
	}

	model.Track = tracks[0]

	return model, nil
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

func (s *Shared) GetArtistAlbums(ctx context.Context, id string) ([]*models.Album, error) {
	return s.databaseService.GetArtistAlbums(ctx, id)
}

func (s *Shared) GetArtistImages(ctx context.Context, id string) ([]*models.Image, error) {
	return s.databaseService.GetArtistImages(ctx, id)
}

func (s *Shared) GetArtistTracks(ctx context.Context, id string) ([]*models.Track, error) {
	return s.databaseService.GetArtistTracks(ctx, id)
}

func (s *Shared) GetImages(ctx context.Context) ([]*models.Image, error) {
	return s.databaseService.GetImages(ctx)
}

func (s *Shared) GetTracks(ctx context.Context) ([]*models.Track, error) {
	return s.databaseService.GetTracks(ctx)
}

func (s *Shared) GetTrackAlbum(ctx context.Context, id string) (*models.Album, error) {
	return s.databaseService.GetTrackAlbum(ctx, id)
}

func (s *Shared) GetTrackArtists(ctx context.Context, id string) ([]*models.Artist, error) {
	return s.databaseService.GetTrackArtists(ctx, id)
}
