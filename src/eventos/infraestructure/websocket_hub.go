package infraestructure

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketMessage estructura para mensajes en WebSocket (eventos)
// Type: "create" | "update" | "delete" | "subscribe" | "attendance"
type WebSocketMessage struct {
	Type     string      `json:"type"`
	EventoID int32       `json:"evento_id"`
	Data     interface{} `json:"data"`
}

// Hub gestiona todas las conexiones WebSocket del módulo eventos
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan WebSocketMessage
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

// Client representa una conexión de cliente
type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan WebSocketMessage
	eventoID int32
}

var hub *Hub

func init() {
	hub = &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan WebSocketMessage, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go hub.run()
}

func GetHub() *Hub {
	return hub
}

// run ejecuta el bucle principal del hub
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("[Eventos WS] Cliente registrado. Total: %d", len(h.clients))

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mu.Unlock()
			log.Printf("[Eventos WS] Cliente desregistrado. Total: %d", len(h.clients))

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					go func(c *Client) {
						h.unregister <- c
					}(client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// readPump lee mensajes del cliente
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadDeadline(time.Now().Add(1440 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(1440 * time.Second))
		return nil
	})

	for {
		var msg WebSocketMessage
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		switch msg.Type {
		case "subscribe":
			c.eventoID = msg.EventoID
			log.Printf("[Eventos WS] Cliente suscrito al evento: %d", msg.EventoID)

		case "create", "update", "delete", "attendance":
			c.hub.broadcast <- msg
		}
	}
}

// writePump escribe mensajes al cliente
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			data, _ := json.Marshal(message)
			w.Write(data)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// BroadcastEvent envía un evento a todos los clientes conectados
func (h *Hub) BroadcastEvent(msgType string, eventoID int32, data interface{}) {
	msg := WebSocketMessage{
		Type:     msgType,
		EventoID: eventoID,
		Data:     data,
	}
	h.broadcast <- msg
}
