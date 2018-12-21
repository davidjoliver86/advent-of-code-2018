package main

import "fmt"

// Player plays the game and has a score.
type Player struct {
	Score int
}

// GetHighScore returns the highest score amongst the players
func GetHighScore(players []*Player) int {
	highScore := players[0].Score
	for i := 1; i < len(players); i++ {
		if players[i].Score > highScore {
			highScore = players[i].Score
		}
	}
	return highScore
}

// Game simulates a playthrough with a number of players and number of marbles.
// Returns a slice of the players' with their accumulated scores.
func Game(numPlayers, iterations int) []*Player {
	players := make([]*Player, numPlayers)
	for i := range players {
		players[i] = &Player{}
	}
	marble := NewGenesisMarble()
	for i := 0; i < iterations; i++ {
		marble = marble.Insert(i, players[i%numPlayers])
	}
	return players
}

func main() {
	// First gold star
	players1 := Game(419, 72164)
	fmt.Println("Game 1 high score was", GetHighScore(players1))

	// Second gold star
	players2 := Game(419, 7216400)
	fmt.Println("Game 2 high score was", GetHighScore(players2))
}
