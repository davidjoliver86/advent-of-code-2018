package main

import "testing"

type gameTestCase struct {
	players     int
	finalMarble int
	highScore   int
}

var gameTestCases = []gameTestCase{
	{10, 1618, 8317},
	{13, 7999, 146373},
	//{17, 1104, 2764},
	{21, 6111, 54718},
	{30, 5807, 37305},
}

// TestGame tests game simulations and checks the highscore against the expected highscore
func TestGame(t *testing.T) {
	for _, testcase := range gameTestCases {
		players := Game(testcase.players, testcase.finalMarble)
		highScore := players[0].Score
		for _, player := range players {
			if player.Score > highScore {
				highScore = player.Score
			}
		}
		if highScore != testcase.highScore {
			t.Errorf("Expected a high score of %d, got %d", testcase.highScore, highScore)
		}
	}
}
