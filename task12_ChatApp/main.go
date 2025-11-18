package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// --- WebSocket setup ---
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Connected clients
var clients = make(map[*websocket.Conn]struct{})

// Message format
type Message struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

// Broadcast message to all clients
func broadcast(msg Message) {
	for conn := range clients {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error writing to client: %v", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}

// Handle each WebSocket client
func handleClient(conn *websocket.Conn) {
	defer func() {
		delete(clients, conn)
		log.Println("Closing WebSocket")
		conn.Close()
	}()

	clients[conn] = struct{}{}

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading JSON: %v", err)
			break
		}
		broadcast(msg)
	}
}

// --- HTTP Handlers ---
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	go handleClient(conn)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./stattic/index.html")
}

// --- Main ---
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
