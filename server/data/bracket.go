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
	match_count := int(c.BracketSize / int(math.Pow(float64(bdef.MatchSize), float64(round_index+1))))
	return Round{MatchIds: make([]string, match_count)}
}

// Check if a given seed has a bye round
// If yes, return -1; if no, return player seed
func checkForBye(seed int, pool_size int) int {
	if seed <= pool_size {
		return seed
	} else {
		return -1
	}
}

// Translate a seed to a team name
func getSeedNames(seed_names []string, seeds []int) []map[string]interface{} {
	names := make([]map[string]interface{}, len(seeds))
	for i, s := range seeds {
		if s > 0 {
			names[i] = map[string]interface{}{
				"name": seed_names[s-1],
				"seed": s,
			}
		} else {
			names[i] = map[string]interface{}{
				"name": "Bye",
				"seed": 0,
			}
		}
	}
	return names
}

// Generate first-round seeded matchups, including first-round byes
// TODO: currently only functional for match sizes of 2
func getStartingMatches(bdef BracketDef, c constraints) [][]map[string]interface{} {
	if c.PoolSize < bdef.MatchSize {
		matchup := make([]string, c.PoolSize)
		for i, t := range bdef.Teams {
			matchup[i] = t
		}
		return [][]map[string]interface{}{}
	}

	seed_matchups := [][]int{{1, 2}} // start from final matchup, featuring 1 vs 2 seed
	for r := 1; r < c.RoundCount; r++ {
		round_matchups := [][]int{}
		max_seed := int(math.Pow(float64(bdef.MatchSize), float64(r+1)) + 1)
		for _, sm := range seed_matchups {
			//upper half draw
			seed_1 := checkForBye(sm[0], c.PoolSize)
			seed_2 := checkForBye(max_seed-sm[0], c.PoolSize)
			if seed_2 == -1 {
				round_matchups = append(round_matchups, []int{seed_1})
			} else {
				round_matchups = append(round_matchups, []int{seed_1, seed_2})
			}

			// bottom half draw
			seed_1 = checkForBye(max_seed-sm[1], c.PoolSize)
			seed_2 = checkForBye(sm[1], c.PoolSize)
			if seed_1 == -1 {
				round_matchups = append(round_matchups, []int{seed_2})
			} else {
				round_matchups = append(round_matchups, []int{seed_1, seed_2})
			}

		}
		seed_matchups = round_matchups
	}

	team_matchups := make([][]map[string]interface{}, len(seed_matchups))
	for i, sm := range seed_matchups {
		team_matchups[i] = getSeedNames(bdef.Teams, sm)
	}
	return team_matchups
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

	starting_matches := getStartingMatches(bdef, c)

	// Generate match-ids for every match, per round
	for i := 0; i < c.RoundCount; i++ {
		bracket.Rounds[i] = allocateRound(bdef, c, i)
		for j := 0; j < len(bracket.Rounds[i].MatchIds); j++ {
			match_id := getMatchID(bdef.Id, i, j, c)
			sb := Scoreboard{Id: match_id, Featured: false}
			if i == 0 {
				// initialize first round matches
				sb.Teams = make([]Score, len(starting_matches[j]))
				for k, team := range starting_matches[j] {
					var s float32 = 0
					sb.Teams[k] = Score{
						Name:  team["name"].(string),
						Seed:  team["seed"].(int),
						Score: &s,
					}
				}
				is_bye := len(sb.Teams) < bdef.MatchSize
				sb.Completed = &is_bye
			} else {
				sb.Teams = make([]Score, bdef.MatchSize)
				for t := range sb.Teams {
					var s float32 = 0
					sb.Teams[t] = Score{Name: "TBD", Score: &s}
				}
				completed := false
				sb.Completed = &completed
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
