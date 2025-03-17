package spotify

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	spotifyLib "github.com/zmb3/spotify/v2"
	spotifyAuth "github.com/zmb3/spotify/v2/auth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Auth struct {
	auth          *spotifyAuth.Authenticator
	state         string
	ch            chan *spotifyLib.Client
	token         *oauth2.Token
	tokenSaveFunc func(*oauth2.Token) error
}

func newAuth(config *utils.Config, token *oauth2.Token, tokenSaveFunc func(*oauth2.Token) error) *Auth {
	redirectURL := fmt.Sprintf("http://%s/callback", config.GetURL())

	auth := spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURL),
		spotifyAuth.WithScopes(spotifyAuth.ScopeUserReadPrivate, spotifyAuth.ScopeUserReadPlaybackState),
		spotifyAuth.WithClientID(config.Spotify.ClientID),
		spotifyAuth.WithClientSecret(config.Spotify.ClientSecret),
	)

	return &Auth{
		state:         "state", // TODO: unique state string to identify the session, should be random
		auth:          auth,
		ch:            make(chan *spotifyLib.Client),
		token:         token,
		tokenSaveFunc: tokenSaveFunc,
	}
}

func (sa *Auth) createClient(ctx context.Context, token *oauth2.Token) (*spotifyLib.Client, error) {
	client := spotifyLib.New(sa.auth.Client(ctx, token))

	token, err := client.Token()
	if err != nil {
		return nil, errors.Wrap(err, "failed getting token")
	}

	// save the token for future use
	if err = sa.tokenSaveFunc(token); err != nil {
		return nil, errors.Wrap(err, "failed saving token")
	}

	return client, nil
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
		}

		// use the token to get an authenticated client
		client, err := sa.createClient(ctx, tok)
		if err != nil {
			utils.Logger.Fatal("failed creating client", zap.Error(err))
		}
		sa.ch <- client

		user, err := client.CurrentUser(context.Background())
		if err != nil {
			utils.Logger.Fatal("failed getting current user", zap.Error(errors.WithContextTags(err, ctx)))
		}

		return c.String(http.StatusOK, "You are logged in as: "+user.User.DisplayName)
	}
}

func (sa *Auth) waitForClient() (*spotifyLib.Client, error) {
	// wait for auth to complete
	client := <-sa.ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		utils.Logger.Error("failed getting current user", zap.Error(err))
		return nil, errors.Wrap(err, "failed getting current user")
	}

	utils.Logger.Info("logged in as: " + user.User.DisplayName)

	return client, nil
}

func (sa *Auth) setupAuthRoutes(e *echo.Echo) {
	e.GET("/callback", sa.finalizeAuth())
}

func (sa *Auth) authenticate() error {
	if sa.token != nil {
		utils.Logger.Info("attempting to use saved token")

		client, err := sa.createClient(context.Background(), sa.token)
		if err != nil {
			utils.Logger.Error("failed creating client", zap.Error(err))
			return errors.Wrap(err, "failed creating client")
		}

		sa.ch <- client

		utils.Logger.Info("new client created from token")

		return nil
	}

	url := sa.auth.AuthURL(sa.state)

	utils.Logger.Error("needs spotify login", zap.String("url", url))

	return errors.New("needs spotify login")
}
