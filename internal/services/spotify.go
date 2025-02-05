package services

import (
	"context"
	"strings"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"

	spotifyLib "github.com/zmb3/spotify/v2"
)

type Spotify struct {
	client *spotify.Spotify
}

func NewSpotify(client *spotify.Spotify) *Spotify {
	return &Spotify{client: client}
}

func (s *Spotify) GetPlayerState(ctx context.Context) (*models.PlayerState, error) {
	playerState, err := s.client.GetPlayerState(ctx)
	if err != nil {
		return nil, err
	}

	context := s.getContext(string(playerState.PlaybackContext.URI))
	context.Href = playerState.PlaybackContext.ExternalURLs["spotify"]

	model := &models.PlayerState{
		Context: context,

		Timestamp:  playerState.Timestamp,
		ProgressMs: int64(playerState.Progress),
		IsPlaying:  playerState.Playing,
	}

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

	context := strings.Split(contextURI, ":")
	//nolint:mnd // Magic number is fine here
	if len(context) != 3 {
		return &models.PlayerStateContext{}
	}

	return &models.PlayerStateContext{
		Type: context[1],
		ID:   context[2],
	}
}

func (s *Spotify) GetArtist(ctx context.Context, id string) (*models.CreateArtistParams, []*models.CreateImageParams, error) {
	artist, err := s.client.GetArtist(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	return spotify.FullArtistToParams(artist), spotify.ImageSliceToModelParams(artist.Images), nil
}

// TODO: This is kind bad. Use a DTO or CreateInput from gqlgen instead
func (s *Spotify) GetAlbum(ctx context.Context, id string) (*models.CreateAlbumParams, []*models.CreateImageParams, []string, error) {
	album, err := s.client.GetAlbum(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	artistIDs := make([]string, 0, len(album.Artists))
	for _, artist := range album.Artists {
		artistIDs = append(artistIDs, string(artist.ID))
	}

	return spotify.FullAlbumToParams(album), spotify.ImageSliceToModelParams(album.Images), artistIDs, nil
}

func (s *Spotify) GetTrack(ctx context.Context, id string) (*models.CreateTrackParams, []string, error) {
	track, err := s.client.GetTrack(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	artistIDs := make([]string, 0, len(track.Artists))
	for _, artist := range track.Artists {
		artistIDs = append(artistIDs, string(artist.ID))
	}

	return spotify.FullTrackToParams(track), artistIDs, nil
}

func (s *Spotify) GetPlaylist(ctx context.Context, id string) (*models.CreatePlaylistParams, []*models.CreateImageParams, error) {
	playlist, err := s.client.GetPlaylist(ctx, id)

	if err != nil {
		errCast, castSucceed := err.(spotifyLib.Error)
		if castSucceed && errCast.Status == 404 {
			return nil, nil, nil
		}

		return nil, nil, err
	}

	return spotify.FullPlaylistToParams(playlist), spotify.ImageSliceToModelParams(playlist.Images), nil
}
