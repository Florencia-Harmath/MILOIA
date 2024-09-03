package router

import (
	"github.com/gorilla/mux"

	"milo-ia/internal/handlers"
)

func UsersRoutes(r *mux.Router) {
	usersRouter := r.PathPrefix("/users").Subrouter()
	
	usersRouter.HandleFunc("/update/{userID}", handlers.UpdateHandler).Methods("PUT")
	usersRouter.HandleFunc("/profile/{userID}", handlers.GetProfileHandler).Methods("GET")
}
