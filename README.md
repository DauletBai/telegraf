![Header](static/assets/brand/logo.png)

# Telegraf

Simple messenger similar to WhatsApp Telegram, only written in Go.

## Project directory structure

/telegraf-app
│
├── /cmd
│   └── /server
│       └── main.go               # Entry point to start the server
│
├── /pkg
│   ├── /api
│   │   ├── auth.go                # Registration and authorization logic
│   │   └── message.go             # Logic forbsending and receiving messeges
│   │
│   ├── /webrtc
│   │   ├── peerconnection.go      # Logic for creating peerConnection and processing webRTC signals
│   │   └── signaling.go           # Logic of SDP and ICE exchange via WebSocket
│   │
│   ├── /ws
│   │   └── websocket.go           # Handing WebSocket connections
│   │
│   └── /db
│       ├── db.go                  # Database connection logic
│       └── schema.sql             # SQL schemas for users and messages table
│
├── /static
│   ├── /css
│   └── /js
│       └── app.js                 # Client side logic for working with WebRTC and WebSocket
│
├── /internal
│   └── config
│       └── config.go              # Application configuration (ports, БД, STUN-servers Etc.)
│
├── /test
│   └── webrtc_test.go             # Testing the work of WebRTC connections and API
│
└── go.mod                         # Go module

