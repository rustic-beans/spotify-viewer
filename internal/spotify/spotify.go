package spotify

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	spotifyLib "github.com/zmb3/spotify/v2"
)

type Spotify struct {
	Client *spotifyLib.Client
	auth   *Auth
}

func NewSpotify(config *utils.Config) *Spotify {
	s := &Spotify{
		auth: newAuth(config),
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
func callSpotify[R *Q, Q any](spot *Spotify, f func() (R, error)) (R, error) {
	var err error

	if spot.Client == nil {
		// TODO: return a custom error type
		return nil, NotAuthenticatedError{}
	}

	response, err := f()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *Spotify) GetPlayerState(ctx context.Context) (*spotifyLib.PlayerState, error) {
	playerState, err := callSpotify(s, func() (*spotifyLib.PlayerState, error) {
		return s.Client.PlayerState(ctx)
	})

	if err != nil {
		return nil, err
	}

	playerState.Timestamp = time.Now().UnixMilli()

	return playerState, nil
}

func (s *Spotify) GetArtist(ctx context.Context, id string) (*spotifyLib.FullArtist, error) {
	artist, err := callSpotify(s, func() (*spotifyLib.FullArtist, error) {
		return s.Client.GetArtist(ctx, spotifyLib.ID(id))
	})

	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (s *Spotify) GetAlbum(ctx context.Context, id string) (*spotifyLib.FullAlbum, error) {
	album, err := callSpotify(s, func() (*spotifyLib.FullAlbum, error) {
		return s.Client.GetAlbum(ctx, spotifyLib.ID(id))
	})

	if err != nil {
		return nil, err
	}

	return album, nil
}

func (s *Spotify) GetTrack(ctx context.Context, id string) (*spotifyLib.FullTrack, error) {
	track, err := callSpotify(s, func() (*spotifyLib.FullTrack, error) {
		return s.Client.GetTrack(ctx, spotifyLib.ID(id))
	})

	if err != nil {
		return nil, err
	}

	return track, nil
}
