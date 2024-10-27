package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rustic-beans/spotify-viewer/utils"
	"github.com/labstack/echo/v4"
	spotifyLib "github.com/zmb3/spotify/v2"
	spotifyAuth "github.com/zmb3/spotify/v2/auth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Auth struct {
	auth      *spotifyAuth.Authenticator
	state     string
	ch        chan *spotifyLib.Client
	tokenFile string
}

func newAuth(config *utils.Config) *Auth {
	redirectURL := fmt.Sprintf("http://%s/callback", config.GetURL())

	auth := spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURL),
		spotifyAuth.WithScopes(spotifyAuth.ScopeUserReadPrivate, spotifyAuth.ScopeUserReadPlaybackState),
		spotifyAuth.WithClientID(config.Spotify.ClientID),
		spotifyAuth.WithClientSecret(config.Spotify.ClientSecret),
	)

	return &Auth{
		state:     "state", // TODO: unique state string to identify the session, should be random
		auth:      auth,
		ch:        make(chan *spotifyLib.Client),
		tokenFile: config.Spotify.TokenFile,
	}
}

func (sa *Auth) createClient(ctx context.Context, token *oauth2.Token) *spotifyLib.Client {
	client := spotifyLib.New(sa.auth.Client(ctx, token))

	token, err := client.Token()
	if err != nil {
		utils.Logger.Error("failed getting token", zap.Error(err))
	}

	jsonData, err := json.Marshal(token)
	if err != nil {
		utils.Logger.Error("failed marshalling token", zap.Error(err))
	}

	err = os.WriteFile(sa.tokenFile, jsonData, 0o600)
	if err != nil {
		utils.Logger.Error("failed writing token to file", zap.Error(err))
	}

	return client
}

// TODO: Refactor this method to be more readable
func (sa *Auth) finalizeAuth() echo.HandlerFunc {
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
		client := sa.createClient(ctx, tok)
		sa.ch <- client

		user, err := client.CurrentUser(context.Background())
		if err != nil {
			utils.Logger.Fatal("failed getting current user", zap.Error(err))
		}

		return c.String(http.StatusOK, "You are logged in as: "+user.User.DisplayName)
	}
}

func (sa *Auth) waitForClient() *spotifyLib.Client {
	// wait for auth to complete
	client := <-sa.ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		utils.Logger.Fatal("failed getting current user", zap.Error(err))
		return nil
	}

	utils.Logger.Info("logged in as: " + user.User.DisplayName)

	return client
}

func (sa *Auth) setupAuthRoutes(e *echo.Echo) {
	e.GET("/callback", sa.finalizeAuth())
}

func (sa *Auth) attemptToReadTokenFromFile() *oauth2.Token {
	data, err := os.ReadFile(sa.tokenFile)
	if err != nil {
		utils.Logger.Error("failed reading token file", zap.Error(err))
		return nil
	}

	var token oauth2.Token

	err = json.Unmarshal(data, &token)
	if err != nil {
		utils.Logger.Error("failed unmarshalling token", zap.Error(err))
		return nil
	}

	return &token
}

func (sa *Auth) authenticate() {
	token := sa.attemptToReadTokenFromFile()
	if token != nil {
		utils.Logger.Info("attempting to use token from file")

		client := sa.createClient(context.Background(), token)

		sa.ch <- client

		utils.Logger.Info("new client created from token")

		return
	}

	url := sa.auth.AuthURL(sa.state)

	utils.Logger.Error("needs spotify login", zap.String("url", url))
}
