package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/albe2669/spotify-viewer/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	QueryPath      = "/query"
	PlaygroundPath = "/playground"
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

func NewServer(srv *handler.Server) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST(QueryPath, echo.WrapHandler(srv))
	e.GET(PlaygroundPath, func(c echo.Context) error {
		playground.Handler("GraphQL", QueryPath).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return e
}
