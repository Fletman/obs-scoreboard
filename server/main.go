package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)
import (
	"scoreboard_server/util/json"
	"scoreboard_server/util/response"
	//"scoreboard_server/api/rest"
	//"scoreboard_server/api/socket"
)
import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	var port int = 8080
	var upgrader = websocket.Upgrader {
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}

	var counter int = 0
	clients := make(map[uuid.UUID]*websocket.Conn)
	var broadcast_mutex sync.Mutex

	http.HandleFunc("/live", func (w http.ResponseWriter, r *http.Request) {
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
			broadcast_mutex.Lock()
			for _, c := range clients {
				c.WriteJSON(string(msg))
			}
			broadcast_mutex.Unlock()
		}
	})

	http.HandleFunc("/scores", func (w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("Received %s request from %s", r.Method, r.RemoteAddr))
		switch r.Method {
		case "GET":
			response.Ok(w, strconv.Itoa(counter))
		case "POST":
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				response.BadRequest(w, "Invalid request body")
			}
			body, err := json.JsonToMap(bytes)
			log.Println(body)
			response.Ok(w, "Message received")
		default:
			response.NotAllowed(w)
		}
	})

	log.Println(fmt.Sprintf("Starting server on port %d", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}