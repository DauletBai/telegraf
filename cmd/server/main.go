package main

import (
	"log"
	"net/http"
	"telegraf/internal/auth"
	"telegraf/internal/config"
	//"telegraf/pkg/webrtc"
	"telegraf/pkg/ws"
)

func main() {
    cfg := config.LoadConfig()
    
    log.Println("Порт сервера:", cfg.ServerPort)
    log.Println("URL базы данных:", cfg.DatabaseURL)
    log.Println("STUN сервера:", cfg.STUNServers)

	http.HandleFunc("/register", auth.RegisterUser)
	http.HandleFunc("/login", auth.LoginUser)
	http.HandleFunc("/ws", ws.HandleSocket)

	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// package main

// import (
//     "log"
//     "telegraf/internal/config"
// )

// func main() {
//     cfg := config.LoadConfig()
    
//     log.Println("Порт сервера:", cfg.ServerPort)
//     log.Println("URL базы данных:", cfg.DatabaseURL)
//     log.Println("STUN сервера:", cfg.STUNServers)
// }