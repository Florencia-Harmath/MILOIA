package router

import (
	"github.com/gorilla/mux"
	"milo-ia/internal/handlers"
	"milo-ia/internal/chat"

	"milo-ia/pkg/middleware"
)

func SetupRouter(hub *chat.Hub) *mux.Router {
	r := mux.NewRouter()

	// Rutas públicas
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Middleware de autenticación para todas las demás rutas
	r.Use(middleware.AuthMiddleware)

	// Rutas de usuario
	UsersRoutes(r)
	
	// Rutas del chat
	handlers.ChatRoutes(r, hub)

	return r
}
