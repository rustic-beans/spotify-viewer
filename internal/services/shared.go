package services

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/multierr"
)

type Shared struct {
	databaseService IDatabase
	spotifyService  *Spotify
	client          *spotify.Spotify
}

func NewShared(databaseService IDatabase, spotifyService *Spotify, client *spotify.Spotify) *Shared {
	return &Shared{databaseService: databaseService, spotifyService: spotifyService, client: client}
}

func (s *Shared) GetArtistsByID(ctx context.Context, ids []string) ([]*models.Artist, error) {
	artists, err := s.databaseService.GetArtistsByID(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting artists")
	}

	if len(artists) == len(ids) {
		return artists, nil
	}

	var errs error

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

		artistParams, err := s.spotifyService.GetArtist(ctx, id)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed getting artist")) {
			continue
		}

		artist, err := s.databaseService.CreateArtist(ctx, artistParams)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed creating artist")) {
			continue
		}

		artists = append(artists, artist)
	}

	return artists, errors.Wrap(errs, "failed getting artists")
}

func (s *Shared) GetAlbumsByID(ctx context.Context, ids []string) ([]*models.Album, error) {
	albums, err := s.databaseService.GetAlbumsByID(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting albums")
	}

	if len(albums) == len(ids) {
		return albums, nil
	}

	var errs error

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

		albumParams, artistIDs, err := s.spotifyService.GetAlbum(ctx, id)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed getting album")) {
			continue
		}

		_, newErrs := s.GetArtistsByID(ctx, artistIDs)
		if multierr.AppendInto(&errs, errors.Wrap(newErrs, "failed getting artists")) {
			continue
		}

		album, err := s.databaseService.CreateAlbum(ctx, albumParams, artistIDs)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed creating album")) {
			continue
		}

		albums = append(albums, album)
	}

	return albums, errors.Wrap(errs, "failed getting albums")
}

func (s *Shared) GetTracksByID(ctx context.Context, ids []string) ([]*models.Track, error) {
	tracks, err := s.databaseService.GetTracksByID(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting tracks")
	}

	if len(tracks) == len(ids) {
		return tracks, nil
	}

	var errs error

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
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed getting track")) {
			continue
		}

		_, err = s.GetArtistsByID(ctx, artistIDs)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed getting artists")) {
			continue
		}

		_, err = s.GetAlbumsByID(ctx, []string{trackParams.AlbumID})
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed getting albums")) {
			continue
		}

		track, err := s.databaseService.CreateTrack(ctx, trackParams, artistIDs)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed creating track")) {
			continue
		}

		tracks = append(tracks, track)
	}

	return tracks, errors.Wrap(errs, "failed getting tracks")
}

func (s *Shared) GetPlaylistByID(ctx context.Context, ids []string) ([]*models.Playlist, error) {
	playlists, err := s.databaseService.GetPlaylistsByID(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting playlists")
	}

	if len(playlists) == len(ids) {
		return playlists, nil
	}

	var errs error

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

		playlistParams, err := s.spotifyService.GetPlaylist(ctx, id)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed getting playlist")) {
			continue
		}

		if playlistParams == nil {
			continue
		}

		playlist, err := s.databaseService.CreatePlaylist(ctx, playlistParams)
		if multierr.AppendInto(&errs, errors.Wrap(err, "failed creating playlist")) {
			continue
		}

		playlists = append(playlists, playlist)
	}

	return playlists, errors.Wrap(errs, "failed getting playlists")
}

func (s *Shared) GetPlayerState(ctx context.Context) (*models.PlayerState, error) {
	playerState, err := s.spotifyService.GetPlayerState(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting player state")
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
			channel <- errors.Wrap(err, "failed getting playerstate track")

			return
		}

		playerState.Track = tracks[0]
		channel <- nil
	}(ctx)

	go func() {
		playerStateContext, err := s.getContext(ctx, playerState.Context)
		if err != nil {
			channel <- errors.Wrap(err, "failed getting playerstate context")

			return
		}

		playerState.Context = playerStateContext
		channel <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-channel; err != nil {
			return nil, errors.Wrap(err, "failed getting playerstate")
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
			return nil, errors.Wrap(err, "failed getting context artist")
		}

		contextModel.Name = artists[0].Name
		contextModel.ImageURL = artists[0].ImageUrl
	case "album":
		albums, err := s.GetAlbumsByID(ctx, []string{contextModel.ID})
		if err != nil {
			return nil, errors.Wrap(err, "failed getting context album")
		}

		contextModel.Name = albums[0].Name
		contextModel.ImageURL = albums[0].ImageUrl
	case "playlist":
		playlists, err := s.GetPlaylistByID(ctx, []string{contextModel.ID})
		if err != nil {
			return nil, errors.Wrap(err, "failed getting context playlist")
		}

		if len(playlists) == 0 {
			contextModel.Name = "Probably Radio"
		} else {
			contextModel.Name = playlists[0].Name
			contextModel.ImageURL = playlists[0].ImageUrl
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

func (s *Shared) GetAlbumTracks(ctx context.Context, id string) ([]*models.Track, error) {
	return s.databaseService.GetAlbumTracks(ctx, id)
}

func (s *Shared) GetArtists(ctx context.Context) ([]*models.Artist, error) {
	return s.databaseService.GetArtists(ctx)
}

func (s *Shared) GetArtistAlbums(ctx context.Context, id string) ([]*models.Album, error) {
	return s.databaseService.GetArtistAlbums(ctx, id)
}

func (s *Shared) GetArtistTracks(ctx context.Context, id string) ([]*models.Track, error) {
	return s.databaseService.GetArtistTracks(ctx, id)
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
