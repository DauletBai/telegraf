package main

import (
	"log"
	"telegraf/pkg/api"
	"telegraf/pkg/webrtc"
	"telegraf/pkg/ws"
	"net/http"
)

func main() {
	http.HandleFunc("/register", api.RegisterUser)
	http.HandleFunc("/login", api.LoginUser)
	http.HandleFunc("/ws", ws.HandleSocket)

	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}