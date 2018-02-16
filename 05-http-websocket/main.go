package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

func broadcast(mu *sync.Mutex, c map[*websocket.Conn]bool, typ int, data []byte) {
	mu.Lock()
	defer mu.Unlock()

	for client := range c {
		go client.WriteMessage(typ, data)
	}
}

func main() {
	// START OMIT
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool { return true },
	}
	var connectionsMu sync.Mutex
	connections := make(map[*websocket.Conn]bool)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return // Unable to upgrade connection
		}
		connections[conn] = true
		defer func() {
			conn.Close()
			delete(connections, conn)
		}()

		for {
			typ, data, err := conn.ReadMessage()
			if err != nil {
				return
			}
			broadcast(&connectionsMu, connections, typ, data)
		}
	})
	// END OMIT
	// Assumes we are running from the repo root
	fs := http.Dir("05-http-websocket/static")
	http.Handle("/", http.FileServer(fs))
	log.Println("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
