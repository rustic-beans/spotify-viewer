package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/rustic-beans/spotify-viewer/generated"
	"github.com/rustic-beans/spotify-viewer/lib/spotify"
	"github.com/rustic-beans/spotify-viewer/resolver"
)

func NewServer(
	spotifyClient *spotify.Spotify,
) *handler.Server {
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{
				SpotifyClient: spotifyClient,
			}},
		),
	)

	return server
}
