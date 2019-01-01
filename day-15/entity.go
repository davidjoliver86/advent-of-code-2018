package main

type Entity interface{}

type Position struct {
	X, Y int
}

type wall struct {
	Position
}

type space struct {
	Position
}

func (w wall) String() string {
	return "#"
}

func (s space) String() string {
	return "."
}

func Wall(x, y int) Entity {
	return wall{Position{x, y}}
}

func Space(x, y int) Entity {
	return space{Position{x, y}}
}
