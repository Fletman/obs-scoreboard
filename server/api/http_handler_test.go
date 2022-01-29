package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"scoreboard/data"
	ujson "scoreboard/util/json"
	"testing"
)

func mockRequest(method string, url string, body interface{}, handler func(http.ResponseWriter, *http.Request)) (res *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	handler(w, request)
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

	response, err := mockRequest("PUT", fmt.Sprintf("/scores/%s", score_id), body, HandleScoreboardRequest)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 200, response.StatusCode)
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
		t.Errorf("Field 'score-id' should have value '%s', was '%s'", score_id, result.Id)
	}
	if *result.Completed != false {
		t.Errorf("Field 'completed' should have value 'false', was '%t'", *result.Completed)
	}
	if result.Featured != false {
		t.Errorf("Field 'featured' should have value 'false', was '%t'", result.Featured)
	}
	for i, team := range body.Teams {
		expected_bytes, _ := json.Marshal(team)
		actual_bytes, _ := json.Marshal(result.Teams[i])

		if expected, actual := string(expected_bytes), string(actual_bytes); expected != actual {
			t.Errorf("Struct mismatch for 'team' field.\nExpected: %s\nActual: %s", expected, actual)
		}
	}
}

func TestCreateBracket(t *testing.T) {
	data.InitScores()

	createPlayers := func(pool_size int) []string {
		players := make([]string, pool_size)
		for i := 0; i < pool_size; i++ {
			players[i] = fmt.Sprintf("team%d", i+1)
		}
		return players
	}

	body := data.BracketDef{
		Id:        "test-bracket",
		MatchSize: 2,
		Teams:     createPlayers(8),
	}

	response, err := mockRequest("POST", "/brackets", body, HandleBracketRequest)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 200, response.StatusCode)
	}

	defer response.Body.Close()
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	var result data.Bracket
	err = json.Unmarshal(payload, &result)
	if err != nil {
		t.Fatal(err)
	}

	// duplicate POST
	response, err = mockRequest("POST", "/brackets", body, HandleBracketRequest)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 409 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 409, response.StatusCode)
	}

	response, err = mockRequest("PUT", "/brackets", body, HandleBracketRequest)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 400 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 400, response.StatusCode)
	}

	response, err = mockRequest("PUT", "/brackets/test-bracket", body, HandleBracketRequest)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 200 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 200, response.StatusCode)
	}
}

func TestCreateMultipleBrackets(t *testing.T) {
	data.InitScores()

	createPlayers := func(pool_size int) []string {
		players := make([]string, pool_size)
		for i := 0; i < pool_size; i++ {
			players[i] = fmt.Sprintf("team%d", i+1)
		}
		return players
	}

	body_1 := data.BracketDef{
		Id:        "test-bracket-1",
		MatchSize: 2,
		Teams:     createPlayers(8),
	}
	body_2 := data.BracketDef{
		Id:        "test-bracket-2",
		MatchSize: 2,
		Teams:     createPlayers(16),
	}

	response_1, err := mockRequest("POST", "/brackets", body_1, HandleBracketRequest)
	if err != nil {
		t.Fatal(err)
	}
	if response_1.StatusCode != 200 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 200, response_1.StatusCode)
	}
	response_2, err := mockRequest("POST", "/brackets", body_2, HandleBracketRequest)
	if err != nil {
		t.Fatal(err)
	}
	if response_2.StatusCode != 200 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 200, response_2.StatusCode)
	}

	list_response, err := mockRequest("GET", "/brackets", nil, HandleBracketRequest)
	if err != nil {
		t.Fatal(err)
	}
	if list_response.StatusCode != 200 {
		t.Fatalf("Expected status code: %d, actual status code: %d", 200, response_1.StatusCode)
	}

	payload, err := ioutil.ReadAll(list_response.Body)
	if err != nil {
		t.Fatal(err)
	}
	res_map, err := ujson.JsonToMap(payload)
	if err != nil {
		t.Fatal(err)
	}
	bracket_ids := res_map["brackets"].([]interface{})
	for i, b_id := range bracket_ids {
		if b_id.(string) == "" {
			t.Errorf("Empty string found at index %d in response", i)
		}
	}
}
