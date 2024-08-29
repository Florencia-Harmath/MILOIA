package handlers

import (
    "net/http"
    "milo-ia/internal/chat"
    "github.com/gorilla/mux"
)

func RegisterChatRoutes(r *mux.Router, hub *chat.Hub) {
    r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        chat.HandleConnection(hub, w, r)
    })
}
