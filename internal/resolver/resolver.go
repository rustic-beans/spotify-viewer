package resolver

import (
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SpotifyClient               *spotify.Spotify
	PlayerStateWebsocketHandler *spotify.PlayerStateWebsocketHandler
}
