package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"scoreboard/util/json"
	"strings"
)

var param_regex *regexp.Regexp

func getPathParams(url string, template string) (path_params map[string]string, err error) {
	path_parts := strings.Split(url, "/")
	part_count := len(path_parts)
	template_parts := strings.Split(template, "/")

	if param_regex == nil {
		param_regex, err = regexp.Compile(`\{(.*?)\}`)
		if err != nil {
			return
		}
	}

	path_params = make(map[string]string)
	for i, part := range template_parts {
		param_name := strings.TrimSuffix(strings.TrimPrefix(part, "{"), "}")
		if part_count > i {
			if param_val := strings.TrimSpace(path_parts[i]); len(param_val) > 0 {
				path_params[param_name] = param_val
			}
		}
	}

	return
}

// Handler for HTTP REST requests for scores
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s request from %s to %s", r.Method, r.RemoteAddr, r.URL.Path))

	path_params, err := getPathParams(r.URL.Path, "/scores/{score-id}")
	if err != nil {
		InternalServerError(w)
	}

	switch r.Method {
	case "GET":
		body := make(map[string]interface{})
		if score_id, ok := path_params["score-id"]; ok {
			// get specific scoreboard
			body[score_id] = score_id

		} else {
			// list all scoreboards
			body["scoreboards"] = []string{"score_one", "score2", "score3"}
		}
		Ok(w, body)
	case "POST":
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			BadRequest(w, "Invalid request body")
		}
		body, _ := json.JsonToMap(bytes)
		log.Println(body)
		Ok(w, body)
	default:
		MethodNotAllowed(w)
	}
}
