package main

import (
	"fmt"
	"sort"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

type carts []*Cart

// Fleet represents all the carts on the track and their positions.
type Fleet struct {
	Carts     carts
	Positions map[*Cart]Coordinates
}

// CARTS
// Mainly to satisfy sort.Interface

func (c carts) Len() int {
	return len(c)
}

func (c carts) Less(i, j int) bool {
	this, that := c[i], c[j]
	if this.Position.Y < that.Position.Y {
		return true
	}
	if this.Position.Y > that.Position.Y {
		return false
	}
	if this.Position.X < that.Position.X {
		return true
	}
	return false
}

func (c carts) Swap(i, j int) {
	this, that := c[i], c[j]
	this.id, that.id = that.id, this.id
	this.Direction, that.Direction = that.Direction, this.Direction
	this.Position, that.Position = that.Position, this.Position
	this.intersections, that.intersections = that.intersections, this.intersections
}

// NewFleet reads the ASCII art of the track and carts and returns a struct of pointers to the carts and their positions.
func NewFleet(input string) *Fleet {
	lines := util.FileLines(input)
	carts := make([]*Cart, 0)
	fleet := &Fleet{}
	id := 1
	positions := make(map[*Cart]Coordinates)
	for row, line := range lines {
		for col, ch := range line {
			if ch == '<' || ch == '>' || ch == 'v' || ch == '^' {
				coordinates := Coordinates{col, row}
				cart := &Cart{id, coordinates, InitialDirection[ch], 0}
				positions[cart] = coordinates
				carts = append(carts, cart)
				id++
			}
		}
	}
	fleet.Carts = carts
	fleet.Positions = positions
	return fleet
}

// MoveAll moves all the carts in the fleet, but the carts must be moved in a certain order.
// All the carts must move in the top->bottom, left->right order they're currently in.
func (f *Fleet) MoveAll(track *Track) {
	sort.Sort(f.Carts)
	for _, cart := range f.Carts {
		fmt.Println(cart, "is moving")
		cart.Move(track)
		f.Positions[cart] = cart.Position
	}
}

// Crashes polls all the carts' coordinates and returns the coordinate pairs of crashes.
// A crash is defined as two or more carts sharing the same coordinates.
func (f *Fleet) Crashes() []Coordinates {
	positions := make(map[Coordinates]int)
	for _, cart := range f.Carts {
		positions[cart.Position]++
	}
	crashes := make([]Coordinates, 0)
	for pos, numCrashes := range positions {
		if numCrashes > 1 {
			crashes = append(crashes, pos)
		}
	}
	return crashes
}
