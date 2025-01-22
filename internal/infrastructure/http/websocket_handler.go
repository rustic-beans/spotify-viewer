package http

import (
	"sync"

	"github.com/google/uuid"
)

type WebsocketHandler[M any] struct {
	mu          sync.Mutex
	connections map[string]chan<- M
}

func NewWebsocketHandler[M any]() *WebsocketHandler[M] {
	return &WebsocketHandler[M]{
		connections: make(map[string]chan<- M),
	}
}

func (w *WebsocketHandler[M]) AddConnection(c chan<- M) string {
	w.mu.Lock()
	defer w.mu.Unlock()

	id := uuid.New().String()
	w.connections[id] = c

	return id
}

func (w *WebsocketHandler[M]) RemoveConnection(id string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	delete(w.connections, id)
}

func (w *WebsocketHandler[M]) Broadcast(m M) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, c := range w.connections {
		c <- m
	}
}
