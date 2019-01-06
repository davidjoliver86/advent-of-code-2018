package main

import (
	"fmt"
	"sort"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

type Game struct {
	board     [][]rune
	creatures Creatures
}

func NewGame(input string) *Game {
	board := make([][]rune, 0)
	creatures := make(Creatures, 0)
	for y, line := range util.FileLines(input) {
		row := make([]rune, len(line))
		for x, char := range line {
			row[x] = char
			if char == Elf || char == Goblin {
				creatures = append(creatures, Creature{kind: char, hp: 200, attack: 3, x: x, y: y})
			}
		}
		board = append(board, row)
	}
	return &Game{
		board:     board,
		creatures: creatures,
	}
}

func (g *Game) Board() [][]rune {
	board := make([][]rune, len(g.board))
	for y, row := range g.board {
		newRow := make([]rune, len(row))
		for x, ch := range row {
			newRow[x] = ch
		}
		board[y] = newRow
	}
	return board
}

func (g *Game) Turn() bool {
	// are there only goblins or elves left
	elves := 0
	goblins := 0
	for _, c := range g.creatures {
		if c.kind == Elf {
			elves++
		}
		if c.kind == Goblin {
			goblins++
		}
	}
	if elves == 0 || goblins == 0 {
		return false // game is over, no more turns
	}

	// proceed through the units in reading order
	sort.Sort(g.creatures)

	for _, active := range g.creatures {
		fmt.Println(active)
		// whats in range
		inRange := make([]Node, 0)
		for _, creature := range g.creatures {
			if active.kind != creature.kind {
				if g.board[creature.y][creature.x-1] == Space {
					inRange = append(inRange, Node{creature.x - 1, creature.y})
				}
				if g.board[creature.y][creature.x+1] == Space {
					inRange = append(inRange, Node{creature.x + 1, creature.y})
				}
				if g.board[creature.y-1][creature.x] == Space {
					inRange = append(inRange, Node{creature.x, creature.y - 1})
				}
				if g.board[creature.y+1][creature.x] == Space {
					inRange = append(inRange, Node{creature.x, creature.y + 1})
				}
			}
		}
		fmt.Println(active, inRange)
		for _, node := range inRange {
			fmt.Println(active.Reachable(node, g.Board()))
			fmt.Println(active.DistanceFrom(node, g.Board()))
		}
	}
	return true
}
