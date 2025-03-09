package http

import (
	"sync"

	"github.com/google/uuid"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"
)

type WebsocketHandler[M any] struct {
	mu               sync.RWMutex
	messageQueueSize int
	connections      map[string]chan M
}

func NewWebsocketHandler[M any](messageQueueSize int) *WebsocketHandler[M] {
	return &WebsocketHandler[M]{
		connections:      make(map[string]chan M),
		messageQueueSize: messageQueueSize,
	}
}

func (w *WebsocketHandler[M]) AddConnection() (id string, ch chan M) {
	utils.Logger.Info("Adding connection")

	id = uuid.New().String()
	ch = make(chan M, w.messageQueueSize)

	w.mu.Lock()
	w.connections[id] = ch
	w.mu.Unlock()

	utils.Logger.Info("Connection added")

	return id, ch
}

func (w *WebsocketHandler[M]) RemoveConnection(id string) {
	w.mu.Lock()
	delete(w.connections, id)
	w.mu.Unlock()
}

func (w *WebsocketHandler[M]) Broadcast(m M) {
	utils.Logger.Info("Broadcasting message")

	w.mu.RLock()
	defer w.mu.RUnlock()

	utils.Logger.Info("Lock acquired", zap.Int("numOfConn", len(w.connections)))

	for id, c := range w.connections {
		select {
		case c <- m:
		default:
			utils.Logger.Info("Message dropped", zap.String("id", id))
		}
	}

	utils.Logger.Info("Message broadcasted")
}
