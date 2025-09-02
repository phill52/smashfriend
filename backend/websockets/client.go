package websockets

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	Conn     *websocket.Conn
	Username string
	mu       sync.Mutex
}

func (c *Client) sendMessage(msg Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Conn.WriteJSON(msg)
}