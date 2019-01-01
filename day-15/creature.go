package main

import "fmt"

type Creature interface {
	Entity
	Alive() bool
}

func (c creature) Alive() bool {
	return c.hp > 0
}

func (c creature) String() string {
	return fmt.Sprintf("%c", c.char)
}

type creature struct {
	Position
	hp   int
	char rune
}

func Goblin(x, y int) Creature {
	return creature{Position{x, y}, 200, 'E'}
}

func Elf(x, y int) Creature {
	return creature{Position{x, y}, 10, 'G'}
}
