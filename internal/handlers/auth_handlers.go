package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/login", LoginHandler).Methods("POST")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Implement login logic and JWT generation
}
