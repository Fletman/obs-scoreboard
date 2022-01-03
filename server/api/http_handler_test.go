package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"scoreboard/data"
	"testing"
)

func mockRequest(method string, url string, body interface{}) (res *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	HandleRequest(w, request)
	res = w.Result()
	return
}

func TestCreateScoreboard(t *testing.T) {
	// test valid body
	data.InitScores()

	score_id := "match-1"
	completed := false
	score_a := float32(10)
	score_b := float32(5)
	body := data.Scoreboard{
		Completed: &completed,
		Teams: []data.Score{
			{Name: "Team A", Score: &score_a},
			{Name: "Team B", Score: &score_b},
		},
	}

	response, err := mockRequest("PUT", fmt.Sprintf("/scores/%s", score_id), body)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 {
		t.Fatalf(fmt.Sprintf("Expected status code: %d, actual status code: %d", 200, response.StatusCode))
	}

	defer response.Body.Close()
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	var result data.Scoreboard
	err = json.Unmarshal(payload, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.Id != score_id {
		t.Error(fmt.Sprintf("Field 'score-id' should have value '%s', was '%s'", score_id, result.Id))
	}
	if *result.Completed != false {
		t.Error(fmt.Sprintf("Field 'completed' should have value 'false', was '%t'", *result.Completed))
	}
	if result.Featured != false {
		t.Error(fmt.Sprintf("Field 'featured' should have value 'false', was '%t'", result.Featured))
	}
	for i, team := range body.Teams {
		expected_bytes, _ := json.Marshal(team)
		actual_bytes, _ := json.Marshal(result.Teams[i])

		if expected, actual := string(expected_bytes), string(actual_bytes); expected != actual {
			t.Error(fmt.Sprintf("Struct mismatch for 'team' field.\nExpected: %s\nActual: %s", expected, actual))
		}
	}
}