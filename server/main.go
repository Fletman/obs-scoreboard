package main

import (
	"fmt"
	"log"
	"net/http"
	"scoreboard/api/rest"
	"scoreboard/api/socket"
)

func main() {
	var port int = 8080

	http.HandleFunc("/live", socket.HandleConnection)

	http.HandleFunc("/scores", rest.HandleRequest)

	log.Println(fmt.Sprintf("Starting server on port %d", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
