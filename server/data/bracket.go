package data

import (
	"fmt"
	"math"
	"scoreboard/util/locks"
	"sort"
	"strings"
)

type constraints struct {
	PoolSize          int
	RoundCount        int
	BracketConstraint int
	Byes              []string
	ByeCount          int
}

type Round struct {
	MatchIds []string `json:"match-ids"`
}

type Bracket struct {
	Id     string  `json:"bracket-id"`
	Rounds []Round `json:"rounds"`
}

type BracketDef struct {
	Id        string   `json:"bracket-id"`
	MatchSize int      `json:"match-size" validate:"required"`
	Teams     []string `json:"teams" validate:"required"`
}

var brackets map[string]Bracket = make(map[string]Bracket)

// Declare constraints for bracket:
//
// - number of rounds
//
// - max logarithmic integer number of rounds
//
// - byes
func getConstraints(participants []string, match_size int) constraints {
	c := constraints{PoolSize: len(participants)}

	c.RoundCount = int(math.Ceil(math.Log(float64(c.PoolSize) / math.Log(float64(match_size)))))
	c.BracketConstraint = int(math.Pow(float64(match_size), float64(c.RoundCount)))
	bye_count := c.BracketConstraint - c.PoolSize
	c.Byes = participants[:bye_count]
	c.ByeCount = len(c.Byes)
	return c
}

// Generate match ID
func getMatchID(bracket_id string, round_index int, match_index int, c constraints) (match_id string) {
	switch round_index {
	case (c.RoundCount - 1):
		match_id = fmt.Sprintf("%s: Finals", bracket_id)
	case (c.RoundCount - 2):
		match_id = fmt.Sprintf("%s: Semifinal-%d", bracket_id, match_index+1)
	case (c.RoundCount - 3):
		match_id = fmt.Sprintf("%s: Quarterfinal-%d", bracket_id, match_index+1)
	default:
		match_id = fmt.Sprintf("%s: Round %d-%d", bracket_id, round_index+1, match_index+1)
	}
	return
}

// Retrieve a list of bracket-ids
func ListBrackets() []string {
	locks.Bracket_Mutex.RLock()
	defer locks.Bracket_Mutex.RUnlock()
	bracket_ids := make([]string, len(brackets))
	i := 0
	for _, b := range brackets {
		bracket_ids[i] = b.Id
		i++
	}
	sort.Strings(bracket_ids)
	return bracket_ids
}

// Retrieve a bracket given its ID
func GetBracket(bracket_id string) (bracket Bracket, ok bool) {
	id := strings.ToLower(bracket_id)
	locks.Bracket_Mutex.RLock()
	defer locks.Bracket_Mutex.RUnlock()
	bracket, ok = brackets[id]
	return
}

func GenerateBracket(bdef BracketDef) Bracket {
	c := getConstraints(bdef.Teams, bdef.MatchSize)

	bracket := Bracket{Id: bdef.Id, Rounds: make([]Round, c.RoundCount)}

	// allocate space for Rounds
	r1_match_count := int((c.PoolSize - c.ByeCount) / bdef.MatchSize)
	bracket.Rounds[0] = Round{MatchIds: make([]string, r1_match_count)}
	for i := 1; i < c.RoundCount; i++ {
		match_count := int(c.BracketConstraint / int(math.Pow(float64(bdef.MatchSize), float64(i+1))))
		bracket.Rounds[i] = Round{MatchIds: make([]string, match_count)}
		for j := 0; j < match_count; j++ {
			match_id := getMatchID(bdef.Id, i, j, c)
			var completed bool = false
			scoreboard := Scoreboard{
				Id:        match_id,
				Completed: &completed,
				Featured:  false,
				Teams:     []Score{},
			}
			SetScoreBoard(match_id, scoreboard)
			bracket.Rounds[i].MatchIds[j] = match_id
		}
	}

	// initialize first-round matches
	for i := 0; i < r1_match_count; i++ {
		match_id := getMatchID(bdef.Id, 0, i, c)
		var completed bool = false
		var team_a string = bdef.Teams[i+c.ByeCount]
		var team_b string = bdef.Teams[c.PoolSize-1-i]
		var score_a float32 = 0
		var score_b float32 = 0

		scoreboard := Scoreboard{
			Id:        match_id,
			Completed: &completed,
			Featured:  false,
			Teams: []Score{
				{Name: team_a, Score: &score_a},
				{Name: team_b, Score: &score_b},
			},
		}
		SetScoreBoard(match_id, scoreboard)
		bracket.Rounds[0].MatchIds[i] = match_id
	}

	// initialize bye matches for 2nd rounds
	for index, bye := range c.Byes {
		match_id := getMatchID(bdef.Id, 1, index, c)
		var completed bool = false
		var score float32 = 0

		scoreboard := Scoreboard{
			Id:        match_id,
			Completed: &completed,
			Featured:  false,
			Teams: []Score{
				{Name: bye, Score: &score},
			},
		}
		SetScoreBoard(match_id, scoreboard)
	}

	id := strings.ToLower(bdef.Id)
	locks.Bracket_Mutex.Lock()
	defer locks.Bracket_Mutex.Unlock()
	brackets[id] = bracket
	return bracket
}

// Delete a bracket given its ID
func DeleteBracket(bracket_id string) (ok bool) {
	id := strings.ToLower(bracket_id)
	locks.Bracket_Mutex.Lock()
	defer locks.Bracket_Mutex.Unlock()
	if bracket, ok := brackets[id]; ok {
		for _, r := range bracket.Rounds {
			for _, m := range r.MatchIds {
				DeleteScoreBoard(m)
			}
		}
		delete(brackets, id)
	}
	return
}
