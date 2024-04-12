package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/albe2669/spotify-viewer/generated"
	"github.com/albe2669/spotify-viewer/resolver"
	"github.com/albe2669/spotify-viewer/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const (
	QueryPath      = "/query"
	PlaygroundPath = "/playground"
	state          = "state" // TODO: unique state string to identify the session, should be random
)

var (
	ch = make(chan *spotify.Client)
)

func configureLogger(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			utils.Logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
}

func graphqlServer() *handler.Server {
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{}},
		),
	)

	return server
}

func completeAuth(auth *spotifyauth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		tok, err := auth.Token(ctx, state, c.Request())
		if err != nil {
			utils.Logger.Error("failed getting token", zap.Error(err))
			return echo.NewHTTPError(http.StatusForbidden, "failed getting token")
		}
		if st := c.FormValue("state"); st != state {
			utils.Logger.Fatal("state mismatch detected", zap.String("state", st), zap.String("expected", state))
			return echo.NewHTTPError(http.StatusForbidden, "state mismatch")
		}

		// use the token to get an authenticated client
		client := spotify.New(auth.Client(ctx, tok))
		fmt.Printf("authenticated")
		ch <- client

		user, err := client.CurrentUser(context.Background())
		if err != nil {
			utils.Logger.Fatal("failed getting current user", zap.Error(err))
		}
		return c.String(http.StatusOK, "You are logged in as: "+user.User.DisplayName)
	}
}

func httpServer(graphqlServer *handler.Server) *echo.Echo {
	e := echo.New()

	configureLogger(e)

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST(QueryPath, echo.WrapHandler(graphqlServer))
	e.GET(PlaygroundPath, func(c echo.Context) error {
		playground.Handler("GraphQL", QueryPath).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return e
}

func authSpotify(e *echo.Echo, config *utils.Config) {
	redirectURL := fmt.Sprintf("http://%s/callback", config.GetURL())

	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURL),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate),
		spotifyauth.WithClientID(config.Spotify.ClientID),
		spotifyauth.WithClientSecret(config.Spotify.ClientSecret),
	)

	e.GET("/callback", completeAuth(auth))

	// get the user to this URL - how you do that is up to you
	// you should specify a unique state string to identify the session
	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
}

func waitForClient() {
	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		utils.Logger.Fatal("failed getting current user", zap.Error(err))
	}
	fmt.Println("You are logged in as:", user.ID)
}

func main() {
	// Logger
	defer func() {
		_ = utils.Logger.Sync()
	}()

	// Config
	config, err := utils.ReadConfig()
	if err != nil {
		utils.Logger.Fatal("failed reading config", zap.Error(err))
	}

	graphqlServer := graphqlServer()
	e := httpServer(graphqlServer)
	authSpotify(e, config)

	go waitForClient()

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}
