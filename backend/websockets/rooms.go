package websockets

import (
	"encoding/json"
	"sync"
)

type Message struct {
	Action string          `json:"action"`
	User   string          `json:"user"`
	Body   json.RawMessage `json:"body"`
	Room   string          `json:"room"`
}

type Room struct {
	ID        string
	clients   map[*Client]bool
	broadcast chan Message
	mu        sync.RWMutex
}

func (r *Room) addClient(client *Client) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.clients[client]; exists {
		return false
	}
	r.clients[client] = true
	return true
}
