package chat

import (
    "github.com/gorilla/websocket"
    "sync"
)

type Client struct {
    Conn *websocket.Conn
    Send chan []byte
}

type Hub struct {
    Clients    map[*Client]bool
    Register   chan *Client
    Unregister chan *Client
    Broadcast  chan []byte
    Mutex      sync.Mutex
}

func NewHub() *Hub {
    return &Hub{
        Clients:    make(map[*Client]bool),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Broadcast:  make(chan []byte),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.Mutex.Lock()
            h.Clients[client] = true
            h.Mutex.Unlock()
        case client := <-h.Unregister:
            h.Mutex.Lock()
            if _, ok := h.Clients[client]; ok {
                delete(h.Clients, client)
                close(client.Send)
            }
            h.Mutex.Unlock()
        case message := <-h.Broadcast:
            h.Mutex.Lock()
            for client := range h.Clients {
                select {
                case client.Send <- message:
                default:
                    close(client.Send)
                    delete(h.Clients, client)
                }
            }
            h.Mutex.Unlock()
        }
    }
}
