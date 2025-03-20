package services

import (
	"context"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"

	spotifyLib "github.com/zmb3/spotify/v2"
)

type Spotify struct {
	client           *spotify.Spotify
	playerStateCache *utils.SingleValueCache[*models.PlayerState]
}

func NewSpotify(client *spotify.Spotify) *Spotify {
	return &Spotify{
		client:           client,
		playerStateCache: utils.NewSingleValueCache[*models.PlayerState](),
	}
}

func (s *Spotify) GetPlayerState(ctx context.Context) (*models.PlayerState, error) {
	if playerState, ok := s.playerStateCache.Get(); ok {
		utils.Logger.Info("Using cached player state", zap.Duration("timeToExpiry", s.playerStateCache.TimeToExpiry()))
		return playerState, nil
	}

	playerState, err := s.client.GetPlayerState(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting player state")
	}

	playerStateContext := s.getContext(string(playerState.PlaybackContext.URI))
	playerStateContext.Href = playerState.PlaybackContext.ExternalURLs["spotify"]

	model := &models.PlayerState{
		Context: playerStateContext,

		Timestamp:  playerState.Timestamp,
		ProgressMs: int64(playerState.Progress),
		IsPlaying:  playerState.Playing,
	}

	//nolint:mnd // Magic number is fine here
	s.playerStateCache.SetWithExpiry(model, 9*time.Second)

	if playerState.Item == nil {
		return model, nil
	}

	model.TrackID = string(playerState.Item.ID)

	return model, nil
}

func (s *Spotify) getContext(contextURI string) *models.PlayerStateContext {
	if contextURI == "" {
		return &models.PlayerStateContext{}
	}

	splitContext := strings.Split(contextURI, ":")
	//nolint:mnd // Magic number is fine here
	if len(splitContext) != 3 {
		return &models.PlayerStateContext{}
	}

	return &models.PlayerStateContext{
		Type: splitContext[1],
		ID:   splitContext[2],
	}
}

func (s *Spotify) GetArtist(ctx context.Context, id string) (artistParams *models.CreateArtistParams, imageParams []*models.CreateImageParams, err error) {
	artist, err := s.client.GetArtist(ctx, id)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed getting artist with id %s", id)
	}

	return spotify.FullArtistToParams(artist), spotify.ImageSliceToModelParams(artist.Images), nil
}

// TODO: This is kind bad. Use a DTO or CreateInput from gqlgen instead
func (s *Spotify) GetAlbum(ctx context.Context, id string) (albumParams *models.CreateAlbumParams, imageParams []*models.CreateImageParams, artistIDs []string, err error) {
	album, err := s.client.GetAlbum(ctx, id)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "failed getting album with id %s", id)
	}

	artistIDs = make([]string, 0, len(album.Artists))
	for _, artist := range album.Artists {
		artistIDs = append(artistIDs, string(artist.ID))
	}

	return spotify.FullAlbumToParams(album), spotify.ImageSliceToModelParams(album.Images), artistIDs, nil
}

func (s *Spotify) GetTrack(ctx context.Context, id string) (trackParams *models.CreateTrackParams, artistIDs []string, err error) {
	track, err := s.client.GetTrack(ctx, id)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed getting track with id %s", id)
	}

	artistIDs = make([]string, 0, len(track.Artists))
	for _, artist := range track.Artists {
		artistIDs = append(artistIDs, string(artist.ID))
	}

	return spotify.FullTrackToParams(track), artistIDs, nil
}

func (s *Spotify) GetPlaylist(ctx context.Context, id string) (playlistParams *models.CreatePlaylistParams, imageParams []*models.CreateImageParams, err error) {
	playlist, err := s.client.GetPlaylist(ctx, id)
	if err != nil {
		errCast, castSucceed := errors.UnwrapAll(err).(spotifyLib.Error)
		if castSucceed && errCast.Status == 404 {
			return nil, nil, nil
		}

		return nil, nil, errors.Wrapf(err, "failed getting playlist with id %s", id)
	}

	return spotify.FullPlaylistToParams(playlist), spotify.ImageSliceToModelParams(playlist.Images), nil
}
