package data

import (
	"fmt"
	"testing"
)

func createPlayers(pool_size int) []string {
	players := make([]string, pool_size)
	for i := 0; i < pool_size; i++ {
		players[i] = fmt.Sprintf("team%d", i+1)
	}
	return players
}

// create bracket with evenly divided teams
func TestCreateEvenBracket(t *testing.T) {
	InitScores()

	bdef := BracketDef{
		Id:        "8-person",
		Teams:     createPlayers(8),
		MatchSize: 2,
	}
	bracket := GenerateBracket(bdef)

	if len(bracket.Rounds) != 3 {
		t.Errorf("Invalid rount count. Expected: 3 | Actual: %d", len(bracket.Rounds))
	}

	match_counts := []int{4, 2, 1}
	for i, mc := range match_counts {
		if len(bracket.Rounds[i].MatchIds) != mc {
			t.Errorf("Invalid count for Round %d. Expected: %d | Actual: %d", i+1, mc, len(bracket.Rounds[i].MatchIds))
		}
	}
}

// create bracket with #1 seed having first round bye
func TestOneByeBracket(t *testing.T) {
	InitScores()

	bdef := BracketDef{
		Id:        "7-person",
		Teams:     createPlayers(7),
		MatchSize: 2,
	}
	bracket := GenerateBracket(bdef)

	if len(bracket.Rounds) != 3 {
		t.Errorf("Invalid rount count. Expected: 3 | Actual: %d", len(bracket.Rounds))
	}

	match_counts := []int{4, 2, 1}
	for i, mc := range match_counts {
		if len(bracket.Rounds[i].MatchIds) != mc {
			t.Errorf("Invalid count for Round %d. Expected: %d | Actual: %d", i+1, mc, len(bracket.Rounds[i].MatchIds))
		}
	}
	_, ok := GetScoreBoard(bracket.Rounds[1].MatchIds[0])
	if !ok {
		t.Errorf("Missing scoreboard-id %s", bracket.Rounds[1].MatchIds[0])
	}
}
