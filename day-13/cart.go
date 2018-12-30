package main

import (
	"fmt"
)

// DIRECTIONS

type Coordinates struct {
	X, Y int
}

type Cart struct {
	id            int
	Position      Coordinates
	Direction     Coordinates
	intersections int
}

var (
	up    = Coordinates{0, -1}
	down  = Coordinates{0, 1}
	left  = Coordinates{-1, 0}
	right = Coordinates{1, 0}
)

// COORDINATES

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

// CART

const (
	turnLeft = iota
	straight
	turnRight
)

var InitialDirection = map[rune]Coordinates{
	'<': left,
	'>': right,
	'^': up,
	'v': down,
}

func (c *Cart) String() string {
	return fmt.Sprintf("Cart(%d)[(%d, %d), %s]", c.id, c.Position.X, c.Position.Y, c.Direction)
}

func (c *Cart) TurnLeft() {
	// left -> down -> right -> up -> left...
	var newDir Coordinates
	switch c.Direction {
	case left:
		newDir = down
	case down:
		newDir = right
	case right:
		newDir = up
	case up:
		newDir = left
	}
	c.Direction = newDir
}

func (c *Cart) TurnRight() {
	// right -> down -> left -> up -> right...
	var newDir Coordinates
	switch c.Direction {
	case right:
		newDir = down
	case down:
		newDir = left
	case left:
		newDir = up
	case up:
		newDir = left
	}
	c.Direction = newDir
}

func (c *Cart) TurnSlash() {
	// Encountered a /
	// right/left -> turn left
	// up/down -> turn right
	switch {
	case c.Direction == left || c.Direction == right:
		c.TurnLeft()
	case c.Direction == up || c.Direction == down:
		c.TurnRight()
	}
}

func (c *Cart) TurnBackslash() {
	// Encountered a \
	// right/left -> turn right
	// up/down -> turn left
	switch {
	case c.Direction == left || c.Direction == right:
		c.TurnRight()
	case c.Direction == up || c.Direction == down:
		c.TurnLeft()
	}
}

func (c *Cart) HandleIntersection() {
	action := c.intersections % 3
	switch action {
	case turnLeft:
		c.TurnLeft()
	case turnRight:
		c.TurnRight()
	}
	c.intersections++
}

func (c *Cart) Move(track *Track) {
	// move according to the cart's direction
	c.Position.X += c.Direction.X
	c.Position.Y += c.Direction.Y

	// set new direction based on the track cell
	cell := track.Type(c.Position)
	switch {
	case cell == TurnSlash:
		c.TurnSlash()
	case cell == TurnBackslash:
		c.TurnBackslash()
	case cell == Intersection:
		c.HandleIntersection()
	}
}
