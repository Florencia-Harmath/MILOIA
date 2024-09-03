package main

import (
	"log"
	"net/http"

	"milo-ia/internal/chat"
	"milo-ia/internal/config"
	"milo-ia/internal/database"
	"milo-ia/internal/router"
	"milo-ia/pkg/auth"

	"github.com/rs/cors"
)


func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    if err := database.ConnectDatabase(cfg); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    database.SetupExtension()

	if err := database.Migrate(database.DB); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}	

    auth.InitJWT(cfg)

    hub := chat.NewHub()
    go hub.Run()

    r := router.SetupRouter(hub)
    c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Origin"},
		AllowCredentials: true,
	})

    handler := c.Handler(r)

    log.Println("Server started on :3000")
    if err := http.ListenAndServe(":3000", handler); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
