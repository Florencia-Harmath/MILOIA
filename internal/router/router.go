package router

import (
    "github.com/gorilla/mux"
    "milo-ia/internal/handlers"
    "milo-ia/internal/chat"
)

func SetupRouter(hub *chat.Hub) *mux.Router {
    r := mux.NewRouter()
    
    // Llama a RegisterRoutes para registrar las rutas de autenticación
    handlers.RegisterRoutes(r)
    
    // Registra las rutas del chat
    handlers.RegisterChatRoutes(r, hub)
    
    return r
}
