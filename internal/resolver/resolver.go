package resolver

import (
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SharedService               *services.Shared
	PlayerStateWebsocketHandler *models.PlayerStateWebsocketHandler
}
