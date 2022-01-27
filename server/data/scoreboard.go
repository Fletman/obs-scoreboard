package data

import (
	"scoreboard/util/locks"
	"strings"
)

// Struct representing score of an individual team
type Score struct {
	Name  string   `json:"name" validate:"required"`
	Score *float32 `json:"score" validate:"required"`
}

// Struct containing all competing teams, their scores, and status of the game
type Scoreboard struct {
	Id        string  `json:"score-id"`
	Teams     []Score `json:"teams" validate:"required"`
	Completed *bool   `json:"completed" validate:"required"`
	Featured  bool    `json:"featured"`
}

// Struct tracking all scoreboards
type ScoreMap struct {
	Scoreboards   map[string]*Scoreboard
	FeaturedScore *Scoreboard
}

var scores ScoreMap = ScoreMap{}

// Initialize data
func InitScores() {
	scores.Scoreboards = make(map[string]*Scoreboard)
	scores.FeaturedScore = nil
}

// Return list of all current scoreboards
func GetScoreList() []Scoreboard {
	locks.Score_Mutex.RLock()
	defer locks.Score_Mutex.RUnlock()
	list := make([]Scoreboard, len(scores.Scoreboards))
	i := 0
	for _, s := range scores.Scoreboards {
		s.Featured = (s == scores.FeaturedScore)
		list[i] = *s
		i++
	}
	return list
}

// Return list of scoreboards containing only supplied score-ids
func GetFilteredScoreList(score_ids []string) []Scoreboard {
	list := []Scoreboard{}
	locks.Score_Mutex.RLock()
	defer locks.Score_Mutex.RUnlock()
	for _, id := range score_ids {
		if s, ok := scores.Scoreboards[id]; ok {
			s.Featured = (s == scores.FeaturedScore)
			list = append(list, *s)
		}
	}
	return list
}

// Return a scoreboard given its ID
func GetScoreBoard(score_id string) (sb Scoreboard, ok bool) {
	id := strings.ToLower(score_id)
	locks.Score_Mutex.RLock()
	defer locks.Score_Mutex.RUnlock()
	scb, ok := scores.Scoreboards[id]
	if ok {
		scb.Featured = (scb == scores.FeaturedScore)
		sb = *scb
	}
	return
}

func GetFeaturedScoreboard() (sb Scoreboard, ok bool) {
	locks.Score_Mutex.RLock()
	defer locks.Score_Mutex.RUnlock()
	if ok = scores.FeaturedScore != nil; ok {
		sb = *scores.FeaturedScore
	}
	return
}

// Create/Update a scoreboard given its ID
func SetScoreBoard(score_id string, new_board Scoreboard) Scoreboard {
	id := strings.ToLower(score_id)
	new_board.Id = score_id
	locks.Score_Mutex.Lock()
	defer locks.Score_Mutex.Unlock()
	scores.Scoreboards[id] = &Scoreboard{}
	*scores.Scoreboards[id] = new_board
	if scores.Scoreboards[id].Featured {
		scores.FeaturedScore = scores.Scoreboards[id]
	}
	return *scores.Scoreboards[id]
}

// Remove a scoreboard given its ID
func DeleteScoreBoard(score_id string) (ok bool) {
	id := strings.ToLower(score_id)
	locks.Score_Mutex.Lock()
	defer locks.Score_Mutex.Unlock()
	s, ok := scores.Scoreboards[id]
	if ok {
		if s == scores.FeaturedScore {
			scores.FeaturedScore = nil
		}
		delete(scores.Scoreboards, id)
	}
	return
}
