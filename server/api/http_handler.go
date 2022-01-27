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

func parseBody(ptr interface{}, b io.ReadCloser) (err_msg string) {
	bytes, err := ioutil.ReadAll(b)
	if err != nil {
		err_msg = "Invalid JSON format"
		return
	}
	e := json.Unmarshal(bytes, ptr)
	if e != nil {
		err_msg = e.Error()
	}
	return
}

func validatePutScore(b io.ReadCloser) (body data.Scoreboard, err_msg string) {
	body = data.Scoreboard{}
	err_msg = parseBody(&body, b)
	if len(err_msg) > 0 {
		// error found
	} else if err := v.Struct(body); err != nil {
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

func validatePostPutBracket(b io.ReadCloser, method string) (body data.BracketDef, err_msg string) {
	body = data.BracketDef{}
	err_msg = parseBody(&body, b)
	if len(err_msg) > 0 {
		// error found
	} else if err := v.Struct(body); err != nil {
		err_msg = "Invalid request body:"
		for _, e := range err.(validator.ValidationErrors) {
			err_msg = fmt.Sprintf("%s %s.", err_msg, e)
		}
	} else if method == "POST" && len(body.Id) == 0 {
		err_msg = "Missing required field match-id"
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
func HandleScoreboardRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s request from %s to %s", r.Method, r.RemoteAddr, r.URL.Path))

	path_params, err := getPathParams(r.URL.Path, "/scores/{score-id}")
	if err != nil {
		InternalServerError(w)
		return
	}
	query_params := r.URL.Query()

	switch r.Method {
	case "GET":
		if score_id, ok := path_params["score-id"]; ok {
			// get specific scoreboard
			if sb, exists := data.GetScoreBoard(score_id); exists {
				Ok(w, sb)
			} else {
				NotFound(w, fmt.Sprintf("No scoreboard found for ID: %s", score_id))
			}
		} else if len(query_params["featured"]) > 0 && query_params["featured"][0] == "true" {
			fsb, exists := data.GetFeaturedScoreboard()
			if exists {
				Ok(w, fsb)
			} else {
				NotFound(w, "No featured scoreboard currently available")
			}
		} else if len(query_params["score-id"]) > 0 {
			Ok(w, map[string][]data.Scoreboard{"scoreboards": data.GetFilteredScoreList(query_params["score-id"])})
		} else {
			// list all scoreboards
			Ok(w, map[string][]data.Scoreboard{"scoreboards": data.GetScoreList()})
		}
	case "PUT":
		if score_id, ok := path_params["score-id"]; ok {
			body, err_msg := validatePutScore(r.Body)
			if len(err_msg) > 0 {
				BadRequest(w, err_msg)
			} else {
				sb := data.SetScoreBoard(score_id, body)
				Ok(w, sb)
				Broadcast(sb)
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
	case "OPTIONS":
		Ok(w, nil)
	default:
		MethodNotAllowed(w)
	}
}

func HandleBracketRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s request from %s to %s", r.Method, r.RemoteAddr, r.URL.Path))

	path_params, err := getPathParams(r.URL.Path, "/brackets/{bracket-id}")
	if err != nil {
		InternalServerError(w)
		return
	}

	switch r.Method {
	case "GET":
		if bracket_id, ok := path_params["bracket-id"]; ok {
			// get specific bracket
			if b, exists := data.GetBracket(bracket_id); exists {
				Ok(w, b)
			} else {
				NotFound(w, fmt.Sprintf("No scoreboard found for ID: %s", bracket_id))
			}
		} else {
			// list all brackets
			Ok(w, map[string][]string{"brackets": data.ListBrackets()})
		}
	case "POST": // create bracket if it doesn't exist, throw if it does
		if _, ok := path_params["bracket-id"]; ok {
			// bracket ID should be in body, not path, for POST
			MethodNotAllowed(w)
			return
		}
		body, err_msg := validatePostPutBracket(r.Body, r.Method)
		if len(err_msg) > 0 {
			BadRequest(w, err_msg)
		} else if _, exists := data.GetBracket(body.Id); exists {
			Conflict(w, fmt.Sprintf("Bracket %s already exists", body.Id))
		} else {
			b := data.GenerateBracket(body)
			Ok(w, b)
		}
	case "PUT": // create a bracket if it doesn't exist, or reset bracket if it does exist
		if bracket_id, ok := path_params["bracket-id"]; ok {
			body, err_msg := validatePostPutBracket(r.Body, r.Method)
			if len(err_msg) > 0 {
				BadRequest(w, err_msg)
			} else if _, ok := data.GetBracket(bracket_id); !ok {
				NotFound(w, fmt.Sprintf("No bracket found for ID: %s", bracket_id))
			} else {
				body.Id = bracket_id
				Ok(w, data.GenerateBracket(body))
			}
		} else {
			BadRequest(w, "No bracket-id provided in path")
		}
	case "DELETE":
		if bracket_id, ok := path_params["bracket-id"]; ok {
			if del := data.DeleteBracket(bracket_id); del {
				Ok(w, map[string]string{"message": fmt.Sprintf("Successfully deleted bracket ID: %s", bracket_id)})
			} else {
				NotFound(w, fmt.Sprintf("No scoreboard found for ID: %s", bracket_id))
			}
		} else {
			BadRequest(w, "No bracket-id provided in path")
		}
	case "OPTIONS":
		Ok(w, nil)
	default:
		MethodNotAllowed(w)
	}
}
