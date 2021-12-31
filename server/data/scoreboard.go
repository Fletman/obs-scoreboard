package data

// Match status enums
var Status_Enum_Set map[string]bool = map[string]bool{"Pending": true, "In Progress": true, "Completed": true}

// Struct representing score of an individual team
type Score struct {
	Name  string  `json:"name"`
	Score float32 `json:"score"`
}

// Struct containing all competing teams, their scores, and status of the game
type Match struct {
	Teams  []Score `json:"teams"`
	Status string  `json:"status"`
}

// Struct tracking all matches
type ScoreList struct {
	Featured    *Match            `json:"featured"`
	Scoreboards map[string]*Match `json:"matches"`
}

var scores *ScoreList

// Initialize data
func InitScores() {
	scores = new(ScoreList)
	scores.Scoreboards = make(map[string]*Match)
}

// Return list of all current scoreboards and featured scoreboard
func GetScoreList() ScoreList {
	return *scores
}

// Return a scoreboard given its ID
func GetScoreBoard(match_id string) (sb Match, ok bool) {
	scb, ok := scores.Scoreboards[match_id]
	if ok {
		sb = *scb
	}
	return
}

// Create/Update a scoreboard given its ID
func SetScoreBoard(match_id string, new_board Match, featured bool) {
	scores.Scoreboards[match_id] = &new_board
	if featured {
		scores.Featured = scores.Scoreboards[match_id]
	}
}

// Remove a scoreboard given its ID
func DeleteScoreBoard(match_id string) {
	delete(scores.Scoreboards, match_id)
	scores.Featured = nil
}
