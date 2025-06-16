package main

import (
	"log"
	"net/http"
	"todo-list-backend/internal/config"
	"todo-list-backend/internal/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := routes.SetupRouter(cfg)

	log.Printf("Server Starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Server failed to start %v", err)
	}
}
