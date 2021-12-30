package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type res_log struct {
	Status    int
	Body      map[string]interface{}
	Timestamp string
}

func res(w http.ResponseWriter, status int, body map[string]interface{}) error {
	log_msg, err := json.MarshalIndent(
		res_log{
			Status:    status,
			Body:      body,
			Timestamp: time.Now().UTC().Format("2006-01-02 15:04:05 UTC"),
		},
		"",
		"  ",
	)
	if err != nil {
		return err
	}
	log.Println(string(log_msg))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	output, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = w.Write(output)
	return err
}

func errRes(w http.ResponseWriter, status int, msg string) (err error) {
	var body = make(map[string]interface{})
	body["message"] = msg
	err = res(w, status, body)
	return err
}

func Ok(w http.ResponseWriter, body map[string]interface{}) (err error) {
	err = res(w, 200, body)
	return
}

func Accepted(w http.ResponseWriter, body map[string]interface{}) (err error) {
	err = res(w, 202, body)
	return
}

func BadRequest(w http.ResponseWriter, msg string) (err error) {
	err = errRes(w, 400, msg)
	return
}

func NotAllowed(w http.ResponseWriter) (err error) {
	err = errRes(w, 405, "Method Not Allowed")
	return
}

func InternalServerError(w http.ResponseWriter) (err error) {
	err = errRes(w, 500, "An Internal Server Error has occurred")
	return
}
