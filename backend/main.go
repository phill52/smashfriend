package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	Action string          `json:"action"` // "join", "leave", "message", "system", "get users"
	User   string          `json:"user"`
	Body   json.RawMessage `json:"body"`
	Room   string          `json:"room"`
}

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

func (r *Room) removeClient(client *Client) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.clients[client]
	delete(r.clients, client)
	return exists
}

func (r *Room) getClients() []*Client {
	r.mu.RLock()
	defer r.mu.RUnlock()

	clients := make([]*Client, 0, len(r.clients))
	for client := range r.clients {
		clients = append(clients, client)
	}
	return clients
}


func (r *Room) handleMessages() {
	for {
		message := <-r.broadcast
		clients := r.getClients()

		for _, client := range clients {
			if err := client.sendMessage(message); err != nil {
				log.Printf("error broadcasting message: %v", err)
				r.removeClient(client)
			}
		}
	}
}

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

var hub = Hub{
	Rooms:          make(map[string]*Room),
	MessageChannel: make(chan Message, 100),
	Connections:    make(map[*websocket.Conn]bool),
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // temp allow all for dev purposes
	},
}

func main() {
	for i := range 3 {
		hub.createRoom(strconv.Itoa(i))
		go hub.Rooms[strconv.Itoa(i)].handleMessages()
	}

	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	r.HandleFunc("/", index).Methods(http.MethodGet)
	r.HandleFunc("/ws", handleConnections).Methods(http.MethodGet)

	port := ":8080"
	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	log.Printf("New connection request from user: %s", username)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer conn.Close()

	client := &Client{
		ID:       conn.RemoteAddr().String(),
		Conn:     conn,
		Username: username,
	}

	hub.mu.Lock()
	hub.Connections[conn] = true
	hub.mu.Unlock()

	cleanup := func() {
		hub.mu.Lock()
		delete(hub.Connections, conn)
		hub.mu.Unlock()

		hub.mu.RLock()
		for _, room := range hub.Rooms {
			if room.removeClient(client) {
				room.broadcast <- Message{
					Action: "leave",
					User:   client.Username,
					Body:   json.RawMessage(fmt.Sprintf(`"%s left the chat"`, client.Username)),
					Room:   room.ID,
				}
			}
		}
		hub.mu.RUnlock()
		conn.Close()
	}
	defer cleanup()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}

		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		hub.mu.RLock()
		room, exists := hub.Rooms[msg.Room]
		hub.mu.RUnlock()

		if !exists {
			client.sendMessage(Message{
				Action: "error",
				Body:   json.RawMessage(fmt.Sprintf(`"Room %s does not exist"`, msg.Room)),
			})
			continue
		}

		switch msg.Action {
		case "join":
			if !room.addClient(client) {
				client.sendMessage(Message{
					Action: "error",
					Body:   json.RawMessage(`"User already in room"`),
				})
				continue
			}

			room.broadcast <- Message{
				Action: "join",
				User:   "system",
				Body:   json.RawMessage(fmt.Sprintf(`"%s joined the chat"`, client.Username)),
				Room:   msg.Room,
			}

		case "message":
			if _, exists := room.clients[client]; !exists {
				client.sendMessage(Message{
					Action: "error",
					Body:   json.RawMessage(`"Not in room"`),
				})
				continue
			}

			room.broadcast <- Message{
				Action: "message",
				User:   client.Username,
				Body:   msg.Body,
				Room:   msg.Room,
			}

		case "leave":
			room.removeClient(client)
			room.broadcast <- Message{
				Action: "leave",
				User:   "system",
				Body:   json.RawMessage(fmt.Sprintf(`"%s left the chat"`, client.Username)),
				Room:   msg.Room,
			}

		case "get users":
			clients := room.getClients()
			users := make([]string, len(clients))
			for i, c := range clients {
				users[i] = c.Username
			}

			userJson, err := json.Marshal(users)
			if err != nil {
				log.Printf("error setting user json: %v", err)
				continue
			}

			client.sendMessage(Message{
				Action: "users",
				User:   "system",
				Body:   userJson,
				Room:   msg.Room,
			})
		}
	}
}
