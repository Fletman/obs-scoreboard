package socket

import (
	"fmt"
	"log"
	"net/http"
	"scoreboard/util/locks"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var clients = make(map[uuid.UUID]*websocket.Conn)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	log.Println("Establishing new websocket connection")
	conn, err := upgrader.Upgrade(w, r, nil) // upgrade http connection to websocket
	if err != nil {
		log.Println("Http upgrade failed: ", err)
		return
	}
	defer conn.Close()

	client_id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		conn.WriteJSON("Failed to establish connection. Please try again")
		return
	}
	clients[client_id] = conn
	log.Println(fmt.Sprintf("Websocket connection established with %s, ID: %s", conn.RemoteAddr(), client_id))

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, client_id)
			return
		}
		log.Println(fmt.Sprintf("Received message from %s: %s", client_id, string(msg)))
		locks.Broadcast_mutex.Lock()
		for _, c := range clients {
			c.WriteJSON(string(msg))
		}
		locks.Broadcast_mutex.Unlock()
	}
}
