package http

import (
	"sync"

	"github.com/rustic-beans/spotify-viewer/internal/utils"
)

type WebsocketHandler[M any] struct {
	mu         sync.RWMutex
	numOfConn  int
	connection chan M
}

func NewWebsocketHandler[M any]() *WebsocketHandler[M] {
	return &WebsocketHandler[M]{
		connection: make(chan M),
	}
}

func (w *WebsocketHandler[M]) AddConnection() chan M {
	utils.Logger.Info("Adding connection")

	w.mu.Lock()
	w.numOfConn++
	w.mu.Unlock()

	utils.Logger.Info("Connection added")

	return w.connection
}

func (w *WebsocketHandler[M]) RemoveConnection() {
	w.mu.Lock()
	w.numOfConn--
	w.mu.Unlock()
}

func (w *WebsocketHandler[M]) Broadcast(m M) {
	utils.Logger.Info("Broadcasting message")

	w.mu.RLock()
	defer w.mu.RUnlock()

	for range w.numOfConn {
		utils.Logger.Debug("Broadcasting message to connection")
		w.connection <- m
	}

	utils.Logger.Info("Message broadcasted")
}
