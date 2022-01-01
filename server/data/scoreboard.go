package data

import (
	"scoreboard/util/locks"
	"strings"
)

// Struct representing score of an individual team
type Score struct {
	Name  string  `json:"name" validate:"required"`
	Score float32 `json:"score" validate:"required"`
}

// Struct containing all competing teams, their scores, and status of the game
type Scoreboard struct {
	Teams     []Score `json:"teams" validate:"required"`
	Completed bool    `json:"completed" validate:"required"`
}

// Struct tracking all scoreboards
type ScoreList struct {
	Scoreboards map[string]*Scoreboard `json:"scoreboards"`
}

var scores *ScoreList

// Initialize data
func InitScores() {
	scores = new(ScoreList)
	scores.Scoreboards = make(map[string]*Scoreboard)
}

// Return list of all current scoreboards
func GetScoreList() ScoreList {
	var s ScoreList
	locks.Data_Mutex.Lock()
	defer locks.Data_Mutex.Unlock()
	s = *scores
	return s
}

// Return a scoreboard given its ID
func GetScoreBoard(score_id string) (sb Scoreboard, ok bool) {
	id := strings.ToLower(score_id)
	locks.Data_Mutex.Lock()
	defer locks.Data_Mutex.Unlock()
	scb, ok := scores.Scoreboards[id]
	if ok {
		sb = *scb
	}
	return
}

// Create/Update a scoreboard given its ID
func SetScoreBoard(score_id string, new_board Scoreboard) Scoreboard {
	id := strings.ToLower(score_id)
	locks.Data_Mutex.Lock()
	defer locks.Data_Mutex.Unlock()
	scores.Scoreboards[id] = &new_board
	return *scores.Scoreboards[id]
}

// Remove a scoreboard given its ID
func DeleteScoreBoard(score_id string) (ok bool) {
	id := strings.ToLower(score_id)
	locks.Data_Mutex.Lock()
	defer locks.Data_Mutex.Unlock()
	_, ok = scores.Scoreboards[id]
	if ok {
		delete(scores.Scoreboards, id)
	}
	return
}
