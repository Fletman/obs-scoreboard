package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"scoreboard/data"
	"strings"

	"github.com/go-playground/validator/v10"
)

var param_regex *regexp.Regexp

var v *validator.Validate = validator.New()

func validatePut(b io.ReadCloser) (body data.Scoreboard, err_msg string) {
	bytes, err := ioutil.ReadAll(b)
	if err != nil {
		err_msg = "Invalid JSON format"
		return
	}

	body = data.Scoreboard{}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		err_msg = err.Error()
	} else if err = v.Struct(body); err != nil {
		err_msg = "Invalid request body:"
		for _, e := range err.(validator.ValidationErrors) {
			err_msg = fmt.Sprintf("%s %s.", err_msg, e)
		}
	} else {
		err_msg = ""
		for _, s := range body.Teams {
			if err = v.Struct(s); err != nil {
				err_msg = fmt.Sprintf("%s %s", err_msg, err)
			}
		}
		if len(err_msg) > 0 {
			err_msg = fmt.Sprintf("Invalid request body: %s", err_msg)
		}
	}

	return
}

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

// Handler for HTTP REST requests for scores
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s request from %s to %s", r.Method, r.RemoteAddr, r.URL.Path))

	path_params, err := getPathParams(r.URL.Path, "/scores/{score-id}")
	if err != nil {
		InternalServerError(w)
		return
	}

	switch r.Method {
	case "GET":
		if score_id, ok := path_params["score-id"]; ok {
			// get specific scoreboard
			if sb, exists := data.GetScoreBoard(score_id); exists {
				Ok(w, sb)
			} else {
				NotFound(w, fmt.Sprintf("No scoreboard found for ID: %s", score_id))
			}
		} else {
			// list all scoreboards
			scoreboards := data.GetScoreList()
			Ok(w, scoreboards)
		}
	case "PUT":
		if score_id, ok := path_params["score-id"]; ok {
			body, err_msg := validatePut(r.Body)
			if len(err_msg) > 0 {
				BadRequest(w, err_msg)
			} else {
				sb := data.SetScoreBoard(score_id, body)
				result := map[string]interface{}{"score-id": score_id, "scoreboard": sb}
				Ok(w, result)
				Broadcast(result)
			}
		} else {
			BadRequest(w, "No score-id provided in path")
		}
	case "DELETE":
		if score_id, ok := path_params["score-id"]; ok {
			// delete specified scoreboard
			if del := data.DeleteScoreBoard(score_id); del {
				Ok(w, map[string]string{"message": fmt.Sprintf("Successfully deleted scoreboard ID: %s", score_id)})
			} else {
				NotFound(w, fmt.Sprintf("No scoreboard found for ID: %s", score_id))
			}
		} else {
			BadRequest(w, "No score-id provided in path")
		}
	default:
		MethodNotAllowed(w)
	}
}
