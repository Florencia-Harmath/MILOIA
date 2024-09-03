package handlers

import (
    "encoding/json"
    "errors"
    "net/http"
    "milo-ia/internal/database"
    "milo-ia/internal/models"
    "milo-ia/pkg/auth"
    "gorm.io/gorm"
    "github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/register", RegisterHandler).Methods("POST")
    r.HandleFunc("/login", LoginHandler).Methods("POST")
}
// RegisterHandler maneja la lógica para registrar un nuevo usuario
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Hash de la contraseña del usuario
    if err := user.HashPassword(user.Password); err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Guardar el usuario en la base de datos
    if err := database.DB.Create(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            http.Error(w, "Username or email already exists", http.StatusConflict)
        } else {
            http.Error(w, "Failed to create user", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginHandler maneja la lógica para autenticar un usuario y generar un JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    if !user.CheckPassword(req.Password) {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Generar el token JWT
    token, err := auth.GenerateToken(user.Username)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
