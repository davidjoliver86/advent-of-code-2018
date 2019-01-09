package main

import (
	"errors"
	"fmt"
)

const (
	Up = iota
	Down
	Left
	Right
)

type Creature struct {
	kind   rune
	hp     int
	attack int
	x      int
	y      int
}

type Creatures []Creature

type Node struct {
	x, y int
}

type DistNode struct {
	x, y, dist int
}

func (d DistNode) Adjacent(x, y int) bool {
	return (d.x-1 == x && d.y == y) ||
		(d.x+1 == x && d.y == y) ||
		(d.x == x && d.y-1 == y) ||
		(d.x == x && d.y+1 == y)
}

func (c Creature) String() string {
	var kind string
	if c.kind == Elf {
		kind = "Elf"
	}
	if c.kind == Goblin {
		kind = "Goblin"
	}
	return fmt.Sprintf("%s@%d,%d", kind, c.x, c.y)
}

// satisfy sort.interface
func (c Creatures) Len() int {
	return len(c)
}

func (c Creatures) Less(i, j int) bool {
	if c[i].y < c[j].y {
		return true
	}
	return c[i].x < c[j].x
}

func (c Creatures) Swap(i, j int) {
	c[i].kind, c[j].kind = c[j].kind, c[i].kind
	c[i].hp, c[j].hp = c[j].hp, c[i].hp
	c[i].attack, c[j].attack = c[j].attack, c[i].attack
	c[i].x, c[j].x = c[j].x, c[i].x
	c[i].y, c[j].y = c[j].y, c[i].y
}

func (c *Creature) Reachable(target Node, board [][]rune) bool {
	queue := make([]Node, 1)
	queue[0] = Node{c.x, c.y}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// west
		if board[node.y][node.x-1] == Space {
			board[node.y][node.x-1] = c.kind
			queue = append(queue, Node{node.x - 1, node.y})
		}
		if target.x == node.x-1 && target.y == node.y {
			return true
		}
		// east
		if board[node.y][node.x+1] == Space {
			board[node.y][node.x+1] = c.kind
			queue = append(queue, Node{node.x + 1, node.y})
		}
		if target.x == node.x+1 && target.y == node.y {
			return true
		}
		// north
		if board[node.y-1][node.x] == Space {
			board[node.y-1][node.x] = c.kind
			queue = append(queue, Node{node.x, node.y - 1})
		}
		if target.x == node.x && target.y == node.y-1 {
			return true
		}
		// south
		if board[node.y+1][node.x] == Space {
			board[node.y+1][node.x] = c.kind
			queue = append(queue, Node{node.x, node.y + 1})
		}
		if target.x == node.x && target.y == node.y+1 {
			return true
		}
	}
	return false
}

func (c *Creature) DistanceFrom(target Node, board [][]rune) (int, error) {
	distances := make([]DistNode, 1)
	distances[0] = DistNode{c.x, c.y, 0}
	currentDist := 0
	found := false
	for !found {
		for _, distNode := range distances {
			if distNode.dist == currentDist {
				// west
				if board[distNode.y][distNode.x-1] == Space {
					board[distNode.y][distNode.x-1] = c.kind
					distances = append(distances, DistNode{distNode.x - 1, distNode.y, currentDist + 1})
				}
				if target.x == distNode.x-1 && target.y == distNode.y {
					found = true
				}
				// east
				if board[distNode.y][distNode.x+1] == Space {
					board[distNode.y][distNode.x+1] = c.kind
					distances = append(distances, DistNode{distNode.x + 1, distNode.y, currentDist + 1})
				}
				if target.x == distNode.x+1 && target.y == distNode.y {
					found = true
				}
				// north
				if board[distNode.y-1][distNode.x] == Space {
					board[distNode.y-1][distNode.x] = c.kind
					distances = append(distances, DistNode{distNode.x, distNode.y - 1, currentDist + 1})
				}
				if target.x == distNode.x && target.y == distNode.y-1 {
					found = true
				}
				// south
				if board[distNode.y+1][distNode.x] == Space {
					board[distNode.y+1][distNode.x] = c.kind
					distances = append(distances, DistNode{distNode.x, distNode.y + 1, currentDist + 1})
				}
				if target.x == distNode.x && target.y == distNode.y+1 {
					found = true
				}
			}
		}
		if !found {
			currentDist++
		}
	}
	winner := DistNode{target.x, target.y, currentDist + 1}
	for currentDist > 0 {
		for _, distNode := range distances {
			if distNode.dist == currentDist {
				if distNode.Adjacent(winner.x, winner.y) {
					if distNode.y < winner.y {
						winner = distNode
						currentDist--
					} else if distNode.x < winner.x {
						winner = distNode
						currentDist--
					}
				}
			}
		}
	}
	if (winner.x+1) == c.x && winner.y == c.y {
		return Left, nil
	}
	if (winner.x-1) == c.x && winner.y == c.y {
		return Right, nil
	}
	if winner.x == c.x && (winner.y+1) == c.y {
		return Up, nil
	}
	if winner.x == c.x && (winner.y-1) == c.y {
		return Down, nil
	}
	return 0, errors.New("I'm confused")
}
