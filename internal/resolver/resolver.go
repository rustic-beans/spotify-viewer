package resolver

import (
	"github.com/rustic-beans/spotify-viewer/ent"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EntClient                   *ent.Client
	SpotifyClient               *spotify.Spotify
	PlayerStateWebsocketHandler *spotify.PlayerStateWebsocketHandler
}
