package chat

import (
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleConnection(hub *Hub, w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }

    client := &Client{Conn: conn, Send: make(chan []byte)}
    hub.Register <- client

    go client.ReadMessages(hub)
    go client.WriteMessages()
}

func (c *Client) ReadMessages(hub *Hub) {
    defer func() {
        hub.Unregister <- c
        c.Conn.Close()
    }()
    for {
        _, msg, err := c.Conn.ReadMessage()
        if err != nil {
            return
        }
        hub.Broadcast <- msg
    }
}

func (c *Client) WriteMessages() {
    for msg := range c.Send {
        if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
            return
        }
    }
}
