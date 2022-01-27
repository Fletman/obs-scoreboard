package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"scoreboard/api"
	"scoreboard/data"
)

func main() {
	var port = flag.Int("p", 8080, "Port to run server on")
	flag.Parse()

	data.InitScores()

	http.HandleFunc("/live", api.HandleConnection)
	http.HandleFunc("/scores", api.HandleScoreboardRequest)
	http.HandleFunc("/scores/", api.HandleScoreboardRequest)
	http.HandleFunc("/brackets", api.HandleBracketRequest)
	http.HandleFunc("/brackets/", api.HandleBracketRequest)

	log.Println(fmt.Sprintf("Starting server on port %d", *port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
