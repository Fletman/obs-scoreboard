package data

type Score struct {
	TeamName string
	Score    float32
}

type Scoreboard struct {
	Teams map[string]Score
}

type ScoreList struct {
	Featured    *Scoreboard
	Scoreboards map[string]Scoreboard
}

var scores *ScoreList

func scoreboardToMap(s Scoreboard) map[string]interface{} {
	m := make(map[string]interface{})
	for _, t := range s.Teams {
		team := make(map[string]interface{})
		team["team-name"] = t.TeamName
		team["score"] = t.Score
		m[t.TeamName] = team
	}
	return m
}

// Initialize data
func InitScores() {
	scores = new(ScoreList)
}

// Return list of all current scoreboards and featured scoreboard
func GetScoreList() ScoreList {
	return *scores
}

// Return a scoreboard given its ID
func GetScoreBoard(score_id string) (map[string]interface{}, bool) {
	if scoreboard, ok := scores.Scoreboards[score_id]; ok {
		return scoreboardToMap(scoreboard), true
	} else {
		var sb map[string]interface{}
		return sb, false
	}
}

// Update a scoreboard given its ID
func SetScoreBoard(score_id string, new_board Scoreboard) {
	scores.Scoreboards[score_id] = new_board
}
