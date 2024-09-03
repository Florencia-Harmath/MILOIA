package middleware

import (
    "net/http"
    "strings"
    "milo-ia/pkg/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Excepciones para las rutas de registro y login
        if strings.HasPrefix(r.URL.Path, "/register") || strings.HasPrefix(r.URL.Path, "/login") {
            next.ServeHTTP(w, r)
            return
        }

        // Verificar el token de autenticación
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        _, err := auth.ValidateToken(token)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
