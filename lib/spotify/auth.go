package spotify

import (
	"context"
	"fmt"
	"net/http"

	"github.com/albe2669/spotify-viewer/utils"
	"github.com/labstack/echo/v4"
	spotifyLib "github.com/zmb3/spotify/v2"
	spotifyAuth "github.com/zmb3/spotify/v2/auth"
	"go.uber.org/zap"
)

type SpotifyAuth struct {
	auth  *spotifyAuth.Authenticator
	state string
	ch    chan *spotifyLib.Client
}

func newSpotifyAuth(config *utils.Config) *SpotifyAuth {
	redirectURL := fmt.Sprintf("http://%s/callback", config.GetURL())

	auth := spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURL),
		spotifyAuth.WithScopes(spotifyAuth.ScopeUserReadPrivate),
		spotifyAuth.WithClientID(config.Spotify.ClientID),
		spotifyAuth.WithClientSecret(config.Spotify.ClientSecret),
	)

	return &SpotifyAuth{
		state: "state", // TODO: unique state string to identify the session, should be random
		auth:  auth,
		ch:    make(chan *spotifyLib.Client),
	}
}

// TODO: Refactor this method to be more readable
func (sa *SpotifyAuth) finalizeAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		tok, err := sa.auth.Token(ctx, sa.state, c.Request())
		if err != nil {
			utils.Logger.Error("failed getting token", zap.Error(err))
			return echo.NewHTTPError(http.StatusForbidden, "failed getting token")
		}
		if st := c.FormValue("state"); st != sa.state {
			utils.Logger.Fatal("state mismatch detected", zap.String("state", st), zap.String("expected", sa.state))
			return echo.NewHTTPError(http.StatusForbidden, "state mismatch")
		}

		// use the token to get an authenticated client
		client := spotifyLib.New(sa.auth.Client(ctx, tok))
		sa.ch <- client

		user, err := client.CurrentUser(context.Background())
		if err != nil {
			utils.Logger.Fatal("failed getting current user", zap.Error(err))
		}
		return c.String(http.StatusOK, "You are logged in as: "+user.User.DisplayName)
	}
}

func (sa *SpotifyAuth) waitForClient() *spotifyLib.Client {
	// wait for auth to complete
	client := <-sa.ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		utils.Logger.Fatal("failed getting current user", zap.Error(err))
	}
	fmt.Println("You are logged in as:", user.ID)

	return client
}

func (sa *SpotifyAuth) setupAuthRoutes(e *echo.Echo) {
	e.GET("/callback", sa.finalizeAuth())
}

func (sa *SpotifyAuth) authenticate() {
	url := sa.auth.AuthURL(sa.state)

	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
}
