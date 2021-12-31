package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"scoreboard/data"
	"strings"
)

type putRequest struct {
	Match    data.Match
	Status   string
	Featured bool
}

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
			if param_val := strings.TrimSpace(path_parts[i]); len(param_val) > 0 && param_regex.MatchString(part) {
				path_params[param_name] = param_val
			}
		}
	}

	return
}

// Handler for HTTP REST requests for matches
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s request from %s to %s", r.Method, r.RemoteAddr, r.URL.Path))

	path_params, err := getPathParams(r.URL.Path, "/matches/{match-id}")
	if err != nil {
		InternalServerError(w)
		return
	}

	switch r.Method {
	case "GET":
		if match_id, ok := path_params["match-id"]; ok {
			// get specific match scoreboard
			if scoreboard, exists := data.GetScoreBoard(match_id); exists {
				Ok(w, scoreboard)
			} else {
				NotFound(w, fmt.Sprintf("No scoreboard found for ID: %s", match_id))
			}

		} else {
			// list all match scoreboards
			scoreboards := data.GetScoreList()
			Ok(w, scoreboards)
		}
	case "PUT":
		if match_id, ok := path_params["match-id"]; ok {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				BadRequest(w, "Invalid request body")
				return
			}
			body := putRequest{}
			err = json.Unmarshal(bytes, &body)
			if err != nil {
				BadRequest(w, err.Error())
			} else if t := body.Match.Teams; t == nil {
				BadRequest(w, "Field 'teams' is missing or incorrectly formatted")
			} else {
				data.SetScoreBoard(match_id, body.Match, body.Featured)
				Ok(w, map[string]interface{}{"match-id": match_id, "match": body.Match})
			}
		} else {
			BadRequest(w, "No match-id provided in path")
		}
	case "DELETE":
		if match_id, ok := path_params["match-id"]; ok {
			// delete specified match scoreboard
			if _, exists := data.GetScoreBoard(match_id); exists {
				data.DeleteScoreBoard(match_id)
				Ok(w, map[string]string{"message": fmt.Sprintf("Successfully deleted scoreboard ID: %s", match_id)})
			} else {
				NotFound(w, fmt.Sprintf("No scoreboard found for ID: %s", match_id))
			}
		}
	default:
		MethodNotAllowed(w)
	}
}
