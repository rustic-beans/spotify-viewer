package spotify

import (
	"context"
	"fmt"

	"github.com/albe2669/spotify-viewer/utils"
	"github.com/labstack/echo/v4"
	spotifyLib "github.com/zmb3/spotify/v2"
)

type Spotify struct {
	Client *spotifyLib.Client
	auth   *SpotifyAuth
}

func NewSpotify(config *utils.Config) *Spotify {
	s := &Spotify{
		auth: newSpotifyAuth(config),
	}

	return s
}

func (s *Spotify) SetupRoutes(e *echo.Echo) {
	s.auth.setupAuthRoutes(e)
}

func (s *Spotify) Authenticate() {
	go func() {
		s.Client = s.auth.waitForClient()
	}()

	s.auth.authenticate()
}

// Don't question it: https://groups.google.com/g/golang-nuts/c/y9IvZgiNowk
func callSpotify[R *Q, Q any](ctx context.Context, spot *Spotify, f func(ctx context.Context, opts ...spotifyLib.RequestOption) (R, error), opts ...spotifyLib.RequestOption) (R, error) {
	var err error

	if spot.Client == nil {
		return nil, fmt.Errorf("client is not authenticated yet")
	}

	spot.Client, err = spot.auth.checkClient(ctx, spot.Client)
	if err != nil {
		return nil, err
	}

	response, err := f(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *Spotify) GetPlayerState(ctx context.Context) (*spotifyLib.PlayerState, error) {
	return callSpotify(ctx, s, s.Client.PlayerState)
}
