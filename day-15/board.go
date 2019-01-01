package main

import (
	"fmt"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

type Board [][]Entity

func (b Board) String() string {
	s := ""
	for _, row := range b {
		for i := range row {
			s += fmt.Sprintf("%v", row[i])
		}
		s += "\n"
	}
	return s
}

func NewBoard(input string) Board {
	newBoard := make(Board, 0)
	for y, line := range util.FileLines(input) {
		row := make([]Entity, len(line))
		for x, ch := range line {
			switch ch {
			case '#':
				row[x] = Wall(x, y)
			case '.':
				row[x] = Space(x, y)
			case 'E':
				row[x] = Elf(x, y)
			case 'G':
				row[x] = Goblin(x, y)
			}
		}
		newBoard = append(newBoard, row)
	}
	return newBoard
}
