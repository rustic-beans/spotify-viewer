package spotify

import (
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
	s.auth.authenticate()

	go func() {
		s.Client = s.auth.waitForClient()
	}()
}
