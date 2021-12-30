package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"scoreboard/api/rest"
	"scoreboard/api/socket"
)

func main() {
	var port = flag.Int("p", 8080, "Port to run server on")
	flag.Parse()

	http.HandleFunc("/live", socket.HandleConnection)

	http.HandleFunc("/scores/", rest.HandleRequest)

	log.Println(fmt.Sprintf("Starting server on port %d", *port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
