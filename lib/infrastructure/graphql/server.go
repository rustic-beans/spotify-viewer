package graphql

import (
	stdhttp "net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/rustic-beans/spotify-viewer/generated"
	"github.com/rustic-beans/spotify-viewer/lib/spotify"
	"github.com/rustic-beans/spotify-viewer/resolver"
)

func NewServer(
	spotifyClient *spotify.Spotify,
	playerStateWebsocketHandler *spotify.PlayerStateWebsocketHandler,
) *handler.Server {
	server := handler.New(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{
				SpotifyClient:               spotifyClient,
				PlayerStateWebsocketHandler: playerStateWebsocketHandler,
			}},
		),
	)

	server.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(_ *stdhttp.Request) bool {
				return true // TODO: Add origin check
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})

	server.SetQueryCache(lru.New(1000))
	server.Use(extension.Introspection{})
	server.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return server
}
