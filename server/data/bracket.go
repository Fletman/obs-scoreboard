package data

import (
	"fmt"
	"math"
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
	Rounds []Round `json:"rounds"`
}

var brackets map[string]*Bracket = make(map[string]*Bracket)

func getConstraints(participants []string, match_size int) constraints {
	c := constraints{PoolSize: len(participants)}

	c.RoundCount = int(math.Ceil(math.Log(float64(c.PoolSize) / math.Log(float64(match_size)))))
	c.BracketConstraint = int(math.Pow(float64(match_size), float64(c.RoundCount)))
	fmt.Println(c)
	bye_count := c.BracketConstraint - c.PoolSize
	c.Byes = participants[:bye_count]
	c.ByeCount = len(c.Byes)
	return c
}

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

func GenerateBracket(bracket_id string, participants []string, match_size int) *Bracket {
	c := getConstraints(participants, match_size)

	bracket := &Bracket{Rounds: make([]Round, c.RoundCount)}

	// allocate space for Rounds
	r1_match_count := int((c.PoolSize - c.ByeCount) / match_size)
	bracket.Rounds[0] = Round{MatchIds: make([]string, r1_match_count)}
	for i := 1; i < c.RoundCount; i++ {
		match_count := int(c.BracketConstraint / int(math.Pow(float64(match_size), float64(i+1))))
		bracket.Rounds[i] = Round{MatchIds: make([]string, match_count)}
		for j := 0; j < match_count; j++ {
			match_id := getMatchID(bracket_id, i, j, c)
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
		match_id := getMatchID(bracket_id, 0, i, c)
		var completed bool = false
		var team_a string = participants[i+c.ByeCount]
		var team_b string = participants[c.PoolSize-1-i]
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
		match_id := getMatchID(bracket_id, 1, index, c)
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

	brackets[bracket_id] = bracket
	return bracket
}
