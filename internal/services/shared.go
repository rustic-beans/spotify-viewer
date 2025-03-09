package services

import (
	"context"
	"sync"

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

type HasID interface {
	GetID() string
}

func getImageUrls(images []*models.Image) []string {
	urls := make([]string, 0, len(images))
	for _, image := range images {
		urls = append(urls, image.Url)
	}

	return urls
}

func getToCreateIDs[T HasID](ids []string, existing []T) []string {
	existingIDs := make(map[string]struct{}, len(existing))
	for _, item := range existing {
		existingIDs[item.GetID()] = struct{}{}
	}

	toGetIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		if _, exists := existingIDs[id]; !exists {
			toGetIDs = append(toGetIDs, id)
		}
	}

	return toGetIDs
}

func (s *Shared) GetArtistsByID(ctx context.Context, ids []string) ([]*models.Artist, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	artists, err := s.databaseService.GetArtistsByID(ctx, ids)
	if err != nil {
		return nil, err
	}

	if len(artists) == len(ids) {
		return artists, nil
	}

	toGetIDs := getToCreateIDs(ids, artists)

	errs := utils.NewEmptyMultiError()
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, id := range toGetIDs {
		wg.Add(1)
		go func(artistID string) {
			defer wg.Done()

			artist, err := s.fetchAndCreateArtist(ctx, artistID)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errs.Add(err)
				return
			}

			artists = append(artists, artist)
		}(id)
	}

	wg.Wait()

	if errs.HasErrors() {
		return nil, errs
	}

	return artists, nil
}

func (s *Shared) fetchAndCreateArtist(ctx context.Context, id string) (*models.Artist, error) {
	artistParams, imageParams, err := s.spotifyService.GetArtist(ctx, id)
	if err != nil {
		utils.Logger.Error("Failed to get artist from Spotify", zap.String("artistID", id), zap.Error(err))
		return nil, err
	}

	images, err := s.databaseService.CreateImages(ctx, imageParams)
	if err != nil {
		utils.Logger.Error("Failed to create artist images", zap.String("artistID", id), zap.Error(err))
		return nil, err
	}

	artist, err := s.databaseService.CreateArtist(ctx, artistParams, getImageUrls(images))
	if err != nil {
		utils.Logger.Error("Failed to create artist", zap.String("artistID", id), zap.Error(err))
		return nil, err
	}

	return artist, nil
}

func (s *Shared) GetAlbumsByID(ctx context.Context, ids []string) ([]*models.Album, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	albums, err := s.databaseService.GetAlbumsByID(ctx, ids)
	if err != nil {
		return nil, err
	}

	if len(albums) == len(ids) {
		return albums, nil
	}

	toGetIDs := getToCreateIDs(ids, albums)

	errs := utils.NewEmptyMultiError()
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, id := range toGetIDs {
		wg.Add(1)
		go func(albumID string) {
			defer wg.Done()

			album, err := s.fetchAndCreateAlbum(ctx, albumID)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errs.Add(err)
				return
			}

			albums = append(albums, album)
		}(id)
	}

	wg.Wait()

	if errs.HasErrors() {
		return nil, errs
	}

	return albums, nil
}

func (s *Shared) fetchAndCreateAlbum(ctx context.Context, id string) (*models.Album, error) {
	albumParams, imageParams, artistIDs, err := s.spotifyService.GetAlbum(ctx, id)
	if err != nil {
		utils.Logger.Error("Failed to get album from Spotify", zap.String("albumID", id), zap.Error(err))
		return nil, err
	}

	fetchCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var images []*models.Image
	errs := utils.NewEmptyMultiError()

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()

		images, err = s.databaseService.CreateImages(fetchCtx, imageParams)
		if err != nil {
			utils.Logger.Error("Failed to create album images", zap.String("albumID", id), zap.Error(err))
			mu.Lock()
			errs.Add(err)
			mu.Unlock()
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := s.GetArtistsByID(fetchCtx, artistIDs)
		if err != nil {
			utils.Logger.Error("Failed to get artists for album", zap.String("albumID", id), zap.Error(err))
			mu.Lock()
			errs.Add(err)
			mu.Unlock()
			cancel()
		}
	}()

	wg.Wait()

	if errs.HasErrors() {
		return nil, errs
	}

	album, err := s.databaseService.CreateAlbum(ctx, albumParams, getImageUrls(images), artistIDs)
	if err != nil {
		utils.Logger.Error("Failed to create album", zap.String("albumID", id), zap.Error(err))
		return nil, err
	}

	return album, nil
}

func (s *Shared) GetTracksByID(ctx context.Context, ids []string) ([]*models.Track, error) {
	tracks, err := s.databaseService.GetTracksByID(ctx, ids)
	if err != nil {
		return nil, err
	}

	if len(tracks) == len(ids) {
		return tracks, nil
	}

	toGetIDs := getToCreateIDs(ids, tracks)

	errs := utils.NewEmptyMultiError()

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, id := range toGetIDs {
		wg.Add(1)

		go func(trackID string) {
			defer wg.Done()

			track, err := s.fetchAndCreateTrack(ctx, trackID)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errs.Add(err)
				return
			}

			tracks = append(tracks, track)
		}(id)
	}

	wg.Wait()

	if errs.HasErrors() {
		return nil, errs
	}

	return tracks, nil
}

func (s *Shared) fetchAndCreateTrack(ctx context.Context, id string) (*models.Track, error) {
	trackParams, artistIDs, err := s.spotifyService.GetTrack(ctx, id)
	if err != nil {
		return nil, err
	}

	_, err = s.GetArtistsByID(ctx, artistIDs)
	if err != nil {
		return nil, err
	}

	_, err = s.GetAlbumsByID(ctx, []string{trackParams.AlbumID})
	if err != nil {
		return nil, err
	}

	track, err := s.databaseService.CreateTrack(ctx, trackParams, artistIDs)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (s *Shared) GetPlaylistByID(ctx context.Context, ids []string) ([]*models.Playlist, error) {
	playlists, err := s.databaseService.GetPlaylistsByID(ctx, ids)
	if err != nil {
		return nil, err
	}

	if len(playlists) == len(ids) {
		return playlists, nil
	}

	toGetIDs := getToCreateIDs(ids, playlists)

	errs := utils.NewEmptyMultiError()

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, id := range toGetIDs {
		wg.Add(1)
		go func(playlistID string) {
			defer wg.Done()

			playlist, err := s.fetchAndCreatePlaylist(ctx, playlistID)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errs.Add(err)
				return
			}

			playlists = append(playlists, playlist)
		}(id)
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return playlists, nil
}

func (s *Shared) fetchAndCreatePlaylist(ctx context.Context, id string) (*models.Playlist, error) {
	playlistParams, imageParams, err := s.spotifyService.GetPlaylist(ctx, id)
	if err != nil {
		return nil, err
	}

	if playlistParams == nil {
		return nil, nil
	}

	images, err := s.databaseService.CreateImages(ctx, imageParams)
	if err != nil {
		return nil, err
	}

	playlist, err := s.databaseService.CreatePlaylist(ctx, playlistParams, getImageUrls(images))
	if err != nil {
		return nil, err
	}

	return playlist, nil
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
		tracks, err := s.GetTracksByID(ctx, []string{playerState.TrackID})
		if err != nil {
			channel <- err

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
		artists, err := s.GetArtistsByID(ctx, []string{contextModel.ID})
		if err != nil {
			return nil, err
		}

		images, err := s.GetArtistImages(ctx, contextModel.ID)
		if err != nil {
			utils.Logger.Error("Failed to get artist images", zap.Error(err))
			return nil, err
		}

		contextModel.Name = artists[0].Name
		contextModel.ImageURL = images[0].Url
	case "album":
		albums, err := s.GetAlbumsByID(ctx, []string{contextModel.ID})
		if err != nil {
			return nil, err
		}

		images, err := s.GetAlbumImages(ctx, contextModel.ID)
		if err != nil {
			utils.Logger.Error("Failed to get album images", zap.Error(err))
			return nil, err
		}

		contextModel.Name = albums[0].Name
		contextModel.ImageURL = images[0].Url
	case "playlist":
		playlists, err := s.GetPlaylistByID(ctx, []string{contextModel.ID})
		if err != nil {
			return nil, err
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
