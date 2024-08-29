package main

import (
    "log"
    "net/http"
    "milo-ia/internal/config"
    "milo-ia/internal/database"
    "milo-ia/internal/chat"
    "milo-ia/internal/router"
    "milo-ia/pkg/auth"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    if err := database.ConnectDatabase(cfg); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

	if err := database.Migrate(database.DB); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}	

    auth.InitJWT(cfg)

    hub := chat.NewHub()
    go hub.Run()

    r := router.SetupRouter(hub)
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
