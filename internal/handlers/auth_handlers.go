package handlers

import (
    "encoding/json"
    "errors"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "gorm.io/gorm"

    "milo-ia/internal/database"
    "milo-ia/internal/models"
    "milo-ia/pkg/auth"
)

// RegisterHandler maneja la lógica para registrar un nuevo usuario
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user models.RegisterUser
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := user.HashPassword(user.Password); err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

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
    var userLogin models.LoginUser

    if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    var user models.RegisterUser
    if err := database.DB.Where("email = ?", userLogin.Email).First(&user).Error; err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    if !user.CheckPassword(userLogin.Password) {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    token, err := auth.GenerateToken(user.Username)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// UpdateHandler maneja la lógica para actualizar los datos de un usuario
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["userID"]

    if id == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }

    userID, err := uuid.Parse(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.RegisterUser
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := database.DB.Model(&user).Where("id = ?", userID).Updates(user).Error; err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// GetProfileHandler maneja la lógica para obtener los datos de perfil de un usuario por ID
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["userID"]

    if id == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }

    userID, err := uuid.Parse(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.RegisterUser
    if err := database.DB.First(&user, userID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
