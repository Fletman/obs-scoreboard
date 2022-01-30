package data

import (
	"fmt"
	"math"
	"scoreboard/util/locks"
	"sort"
	"strings"
)

type constraints struct {
	PoolSize    int      // Total number of players
	RoundCount  int      // Number of rounds in bracket
	BracketSize int      // Max possible number of players for bracket (match size^round count)
	Byes        []string // List of players who have first-round byes, ordered by seeding
	ByeCount    int      // Number of players with first-round byes
}

type Round struct {
	MatchIds []string `json:"match-ids"`
}

type Bracket struct {
	Id     string  `json:"bracket-id"`
	Rounds []Round `json:"rounds"`
}

type BracketDef struct {
	Id        string   `json:"bracket-id"`                     // Bracket ID
	PoolSize  int      `json:"pool-size"`                      // total number of participants in bracket (TODO: currently not in use)
	MatchSize int      `json:"match-size" validate:"required"` // number of players per match
	Teams     []string `json:"teams"`                          // bracket participants ordered by seeding
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
	c.RoundCount = int(math.Ceil(math.Log(float64(c.PoolSize)) / math.Log(float64(match_size))))
	c.BracketSize = int(math.Pow(float64(match_size), float64(c.RoundCount)))
	bye_count := c.BracketSize - c.PoolSize
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

// Allocate space for each round
func allocateRound(bdef BracketDef, c constraints, round_index int) Round {
	var match_count int
	if round_index == 0 {
		// check for first round byes before allocation
		match_count = int((c.PoolSize - c.ByeCount) / bdef.MatchSize)
	} else {
		match_count = int(c.BracketSize / int(math.Pow(float64(bdef.MatchSize), float64(round_index+1))))
	}
	return Round{MatchIds: make([]string, match_count)}
}

// Sort seeds in the order they should play each other
// TODO: always assumes match size of 2 currently
// TODO: will maybe come back to this at a later date? Not used for now
/*
func seedOrdering(bdef BracketDef, c constraints) []int {
	seeds := []int{1, 2}
	for i := 0; i < c.RoundCount; i++ {
		var match_seeds []int
		l := len(seeds) + 1
		for _, seed := range seeds {
			match_seeds = append(match_seeds, seed)
			match_seeds = append(match_seeds, l-seed)
		}
		seeds = match_seeds
	}
	return seeds[:c.BracketSize]
}
*/

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

	// Generate match-ids for every match, per round
	for i := 0; i < c.RoundCount; i++ {
		bracket.Rounds[i] = allocateRound(bdef, c, i)
		for j := 0; j < len(bracket.Rounds[i].MatchIds); j++ {
			match_id := getMatchID(bdef.Id, i, j, c)
			completed := false
			sb := Scoreboard{
				Id:        match_id,
				Teams:     []Score{},
				Completed: &completed,
				Featured:  false,
			}
			SetScoreBoard(match_id, sb)
			bracket.Rounds[i].MatchIds[j] = match_id
		}
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
