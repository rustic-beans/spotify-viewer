package spotify

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	spotifyLib "github.com/zmb3/spotify/v2"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Spotify struct {
	Client *spotifyLib.Client
	auth   *Auth
}

func NewSpotify(config *utils.Config, token *oauth2.Token, tokenSaveFunc func(*oauth2.Token) error) *Spotify {
	s := &Spotify{
		auth: newAuth(config, token, tokenSaveFunc),
	}

	return s
}

func (s *Spotify) SetupRoutes(e *echo.Echo) {
	s.auth.setupAuthRoutes(e)
}

func (s *Spotify) Authenticate() error {
	go func() {
		client, err := s.auth.waitForClient()
		if err != nil {
			utils.Logger.Error("Failed to authenticate", zap.Error(err))
		}

		s.Client = client
	}()

	return s.auth.authenticate()
}

// Don't question it: https://groups.google.com/g/golang-nuts/c/y9IvZgiNowk
func callSpotify[R *Q, Q any](spot *Spotify, f func() (R, error)) (R, error) {
	if spot.Client == nil {
		return nil, NotAuthenticatedError{}
	}

	response, err := f()

	return response, errors.Wrap(err, "failed to call spotify")
}

func (s *Spotify) GetPlayerState(ctx context.Context) (*spotifyLib.PlayerState, error) {
	playerState, err := callSpotify(s, func() (*spotifyLib.PlayerState, error) {
		return s.Client.PlayerState(ctx)
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed getting player state from spotify")
	}

	playerState.Timestamp = time.Now().UnixMilli()

	return playerState, nil
}

func (s *Spotify) GetArtist(ctx context.Context, id string) (*spotifyLib.FullArtist, error) {
	artist, err := callSpotify(s, func() (*spotifyLib.FullArtist, error) {
		return s.Client.GetArtist(ctx, spotifyLib.ID(id))
	})

	return artist, errors.Wrap(err, "failed getting artist from spotify")
}

func (s *Spotify) GetAlbum(ctx context.Context, id string) (*spotifyLib.FullAlbum, error) {
	album, err := callSpotify(s, func() (*spotifyLib.FullAlbum, error) {
		return s.Client.GetAlbum(ctx, spotifyLib.ID(id))
	})

	return album, errors.Wrap(err, "failed getting album from spotify")
}

func (s *Spotify) GetTrack(ctx context.Context, id string) (*spotifyLib.FullTrack, error) {
	track, err := callSpotify(s, func() (*spotifyLib.FullTrack, error) {
		return s.Client.GetTrack(ctx, spotifyLib.ID(id))
	})

	return track, errors.Wrap(err, "failed getting track from spotify")
}

func (s *Spotify) GetPlaylist(ctx context.Context, id string) (*spotifyLib.FullPlaylist, error) {
	playlist, err := callSpotify(s, func() (*spotifyLib.FullPlaylist, error) {
		return s.Client.GetPlaylist(ctx, spotifyLib.ID(id))
	})

	return playlist, errors.Wrap(err, "failed getting playlist from spotify")
}
