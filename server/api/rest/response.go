package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type res_log struct {
	Status    int
	Message   string
	Timestamp string
}

type res_msg struct {
	Message string
}

func res(w http.ResponseWriter, status int, msg string) error {
	res_data, err := json.MarshalIndent(
		res_log{
			Status:    status,
			Message:   msg,
			Timestamp: time.Now().UTC().Format("2006-01-02 15:04:05 UTC"),
		},
		"",
		"  ",
	)
	if err != nil {
		return err
	}
	log.Println(string(res_data))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	output, err := json.Marshal(res_msg{Message: msg})
	if err != nil {
		return err
	}
	_, err = w.Write(output)
	return err
}

func Ok(w http.ResponseWriter, msg string) (err error) {
	err = res(w, 200, msg)
	return
}

func BadRequest(w http.ResponseWriter, msg string) (err error) {
	err = res(w, 400, msg)
	return
}

func NotAllowed(w http.ResponseWriter) (err error) {
	err = res(w, 405, "Method Not Allowed")
	return
}

func InternalServerError(w http.ResponseWriter) (err error) {
	err = res(w, 500, "An Internal Server Error has occurred")
	return
}
