package models

import "github.com/rustic-beans/spotify-viewer/internal/infrastructure/http"

type PlayerStateContext struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
	Href     string `json:"href"`
}

type PlayerState struct {
	Timestamp  int64               `json:"timestamp"`
	ProgressMs int64               `json:"progressMs"`
	IsPlaying  bool                `json:"isPlaying"`
	TrackID    string              `json:"trackId"`
	Track      *Track              `json:"track"`
	Context    *PlayerStateContext `json:"context"`
}

type PlayerStateWebsocketHandler = http.WebsocketHandler[*PlayerState]
