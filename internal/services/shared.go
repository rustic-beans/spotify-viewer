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

func (s *Shared) GetArtistsByID(ctx context.Context, ids []string) ([]*models.Artist, []error) {
	artists, err := s.databaseService.GetArtistsByID(ctx, ids)
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

func (s *Shared) GetAlbumsByID(ctx context.Context, ids []string) ([]*models.Album, []error) {
	albums, err := s.databaseService.GetAlbumsByID(ctx, ids)
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
		if err != nil {
			errs = append(errs, err)
			continue
		}

		_, newErrs := s.GetArtistsByID(ctx, artistIDs)
		if newErrs != nil {
			errs = append(errs, newErrs...)
			continue
		}

		album, err := s.databaseService.CreateAlbum(ctx, albumParams, getImageUrls(images), artistIDs)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		albums = append(albums, album)
	}

	if len(errs) == 0 {
		return albums, nil
	}

	return albums, errs
}

func (s *Shared) GetTracksByID(ctx context.Context, ids []string) ([]*models.Track, []error) {
	tracks, err := s.databaseService.GetTracksByID(ctx, ids)
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

		_, errs = s.GetArtistsByID(ctx, artistIDs)
		if errs != nil {
			errs = append(errs, errs...)
			continue
		}

		_, newErrs := s.GetAlbumsByID(ctx, []string{trackParams.AlbumID})
		if newErrs != nil {
			errs = append(errs, newErrs...)
			continue
		}

		track, err := s.databaseService.CreateTrack(ctx, trackParams, artistIDs)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		tracks = append(tracks, track)
	}

	if len(errs) == 0 {
		return tracks, nil
	}

	return tracks, errs
}

func (s *Shared) GetPlaylistByID(ctx context.Context, ids []string) ([]*models.Playlist, []error) {
	playlists, err := s.databaseService.GetPlaylistsByID(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	if len(playlists) == len(ids) {
		return playlists, nil
	}

	errs := make([]error, 0, len(ids))

	for _, id := range ids {
		found := false

		for _, playlist := range playlists {
			if playlist.ID == id {
				found = true
				break
			}
		}

		if found {
			continue
		}

		playlistParams, imageParams, err := s.spotifyService.GetPlaylist(ctx, id)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if playlistParams == nil {
			continue
		}

		images, err := s.databaseService.CreateImages(ctx, imageParams)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		playlist, err := s.databaseService.CreatePlaylist(ctx, playlistParams, getImageUrls(images))
		if err != nil {
			errs = append(errs, err)
			continue
		}

		playlists = append(playlists, playlist)
	}

	if len(errs) == 0 {
		return playlists, nil
	}

	return playlists, errs
}

func (s *Shared) GetPlayerState(ctx context.Context) (*models.PlayerState, error) {
	playerState, err := s.spotifyService.GetPlayerState(ctx)
	if err != nil {
		return nil, err
	}

	if playerState.TrackID == "" {
		utils.Logger.Info("No track playing")
		return playerState, nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//nolint:mnd // There can only be 2 channels
	channel := make(chan error, 2)

	go func(ctx context.Context) {
		tracks, errs := s.GetTracksByID(ctx, []string{playerState.TrackID})
		if errs != nil {
			for _, err := range errs {
				utils.Logger.Error("Failed to get track", zap.Error(err))
			}
			channel <- errs[0]

			return
		}

		playerState.Track = tracks[0]
		channel <- nil
	}(ctx)

	go func() {
		playerStateContext, err := s.getContext(ctx, playerState.Context)
		if err != nil {
			utils.Logger.Error("Failed to get context", zap.Error(err))
			channel <- err

			return
		}

		playerState.Context = playerStateContext
		channel <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-channel; err != nil {
			return nil, err
		}
	}

	return playerState, nil
}

func (s *Shared) getContext(ctx context.Context, contextModel *models.PlayerStateContext) (*models.PlayerStateContext, error) {
	if contextModel.ID == "" {
		return contextModel, nil
	}

	switch contextModel.Type {
	case "artist":
		artists, errs := s.GetArtistsByID(ctx, []string{contextModel.ID})
		if errs != nil {
			for _, err := range errs {
				utils.Logger.Error("Failed to get artist", zap.Error(err))
			}

			return nil, errs[0]
		}

		images, err := s.GetArtistImages(ctx, contextModel.ID)
		if err != nil {
			utils.Logger.Error("Failed to get artist images", zap.Error(err))
			return nil, err
		}

		contextModel.Name = artists[0].Name
		contextModel.ImageURL = images[0].Url
	case "album":
		albums, errs := s.GetAlbumsByID(ctx, []string{contextModel.ID})
		if errs != nil {
			for _, err := range errs {
				utils.Logger.Error("Failed to get album", zap.Error(err))
			}

			return nil, errs[0]
		}

		images, err := s.GetAlbumImages(ctx, contextModel.ID)
		if err != nil {
			utils.Logger.Error("Failed to get album images", zap.Error(err))
			return nil, err
		}

		contextModel.Name = albums[0].Name
		contextModel.ImageURL = images[0].Url
	case "playlist":
		playlists, errs := s.GetPlaylistByID(ctx, []string{contextModel.ID})
		if errs != nil {
			for _, err := range errs {
				utils.Logger.Error("Failed to get playlist", zap.Error(err))
			}

			return nil, errs[0]
		}

		if len(playlists) == 0 {
			contextModel.Name = "Probably Radio"
		} else {
			images, err := s.GetPlaylistImages(ctx, contextModel.ID)
			if err != nil {
				utils.Logger.Error("Failed to get playlist images", zap.Error(err))
				return nil, err
			}

			contextModel.Name = playlists[0].Name
			contextModel.ImageURL = images[0].Url
		}
	}

	return contextModel, nil
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

func (s *Shared) GetPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	return s.databaseService.GetPlaylists(ctx)
}

func (s *Shared) GetPlaylistImages(ctx context.Context, id string) ([]*models.Image, error) {
	return s.databaseService.GetPlaylistImages(ctx, id)
}
