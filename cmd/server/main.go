package main

import (
	"log"
	"telegraf/internal/auth"
	"telegraf/internal/config"
	"telegraf/internal/db"
	"telegraf/pkg/api"
	"telegraf/pkg/ws"
	"net/http"
)

func main() {
	// Loading the configuration
    cfg := config.LoadConfig()

	// Connecting to the database
	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer database.Close()

	// Application of migration
	err = db.ApplyMigrations(database, "internal/db/migrations")
	if err != nil {
		log.Fatal("Failed to apply migrations:", err)
	}

	// Setting up routers
	http.HandleFunc("/register", api.RegisterUser)
	http.HandleFunc("/login", api.LoginUser)

	// Protecting middleware routes
	http.Handle("/messages", auth.AuthMiddleware(http.HandlerFunc(api.HandleMessages)))
	http.Handle("/ws", auth.AuthMiddleware(http.HandlerFunc(ws.HandleWebSocket)))

	log.Println("Server running on port :", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}