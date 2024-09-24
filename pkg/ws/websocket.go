package ws

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "telegraf/pkg/webrtc"
)

var upgrader = websocket.Upgrader{}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Ошибка при апгрейде до WebSocket:", err)
        return
    }
    defer conn.Close()

    for {
        var signal webrtc.SignalMessage
        err := conn.ReadJSON(&signal)
        if err != nil {
            log.Println("Ошибка при чтении сигнала:", err)
            break
        }

        // Обработка SDP или ICE сигналов
        webrtc.HandleSignal(signal)
    }
}