package models

import "github.com/rustic-beans/spotify-viewer/internal/infrastructure/http"

type PlayerState struct {
	ContextType string `json:"contextType"`
	ContextURI  string `json:"contextUri"`

	Timestamp  int64  `json:"timestamp"`
	ProgressMs int64  `json:"progressMs"`
	IsPlaying  bool   `json:"isPlaying"`
	Track      *Track `json:"track"`
}

type PlayerStateWebsocketHandler = http.WebsocketHandler[*PlayerState]
