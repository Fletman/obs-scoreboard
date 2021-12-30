package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"scoreboard/util/json"
	"strconv"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("Received %s request from %s", r.Method, r.RemoteAddr))
	switch r.Method {
	case "GET":
		Ok(w, strconv.Itoa(5))
	case "POST":
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			BadRequest(w, "Invalid request body")
		}
		body, _ := json.JsonToMap(bytes)
		log.Println(body)
		Ok(w, "Message received")
	default:
		NotAllowed(w)
	}
}
