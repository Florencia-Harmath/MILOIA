package router

import (
    "github.com/gorilla/mux"
    "milo-ia/internal/handlers"
    "milo-ia/internal/chat"
)

func SetupRouter(hub *chat.Hub) *mux.Router {
    r := mux.NewRouter()
    handlers.RegisterRoutes(r)
    handlers.RegisterChatRoutes(r, hub)
    return r
}
