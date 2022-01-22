package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = flag.Int("p", 8081, "Port to run server on")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./dist")))

	log.Println(fmt.Sprintf("Starting server on port %d", *port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
