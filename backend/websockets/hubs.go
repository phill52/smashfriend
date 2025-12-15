package websockets

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms          map[string]*Room
	MessageChannel chan Message
	Connections    map[*websocket.Conn]bool
	mu             sync.RWMutex
}

func (h *Hub) createRoom(roomID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.Rooms[roomID] = &Room{
		ID:        roomID,
		clients:   make(map[*Client]bool),
		broadcast: make(chan Message, 100),
	}
}