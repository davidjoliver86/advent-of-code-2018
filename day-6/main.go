package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

type Point struct {
	X, Y int
}

// distindes and distindicies keeps track of which point (index of a []Point) has a given
// manhattan distance from any arbitrary spot on a grid
type distindex struct {
	index, distance int
}

// because we need to implement Sort taking into account distance only
type distindicies []distindex

func (d distindicies) Len() int {
	return len(d)
}

func (d distindicies) Less(i, j int) bool {
	return d[i].distance < d[j].distance
}

func (d distindicies) Swap(i, j int) {
	d[i].distance, d[j].distance = d[j].distance, d[i].distance
	d[i].index, d[j].index = d[j].index, d[i].index
}

type grid [][]int

func manhattan(x1, y1, x2, y2 int) int {
	return util.Abs(x1-x2) + util.Abs(y1-y2)
}

func Grid(points []Point) grid {
	// find max coordinate
	max := 0
	for _, point := range points {
		if point.X > max {
			max = point.X
		}
		if point.Y > max {
			max = point.Y
		}
	}

	// allow one extra row/column of buffer
	max++

	// declare 2d array of max-coordinate length x width
	grid := make([][]int, max)
	for i := range grid {
		grid[i] = make([]int, max)
	}

	// fill in grid with shortest manhattan distances; there must be a clear winner
	for gridY := 0; gridY < max; gridY++ {
		for gridX := 0; gridX < max; gridX++ {
			candidates := make(distindicies, len(points))
			for i, point := range points {
				candidates[i] = distindex{distance: manhattan(gridX, gridY, point.X, point.Y), index: i}
			}
			sort.Sort(candidates)
			if candidates[0].distance == candidates[1].distance {
				grid[gridY][gridX] = -1
			} else {
				grid[gridY][gridX] = candidates[0].index
			}
		}
	}
	return grid
}

func Area(g grid) map[int]int {
	areas := make(map[int]int)

	// quick summation across the grid
	for _, row := range g {
		for _, col := range row {
			areas[col]++
		}
	}

	// then remove any index on any of the outer edges
	for _, val := range g[0] {
		delete(areas, val)
	}
	for _, val := range g[len(g)-1] {
		delete(areas, val)
	}
	for _, row := range g {
		last := len(row) - 1
		delete(areas, row[0])
		delete(areas, row[last])
	}
	return areas
}

func main() {
	input := util.FileLines("input.txt")
	points := make([]Point, len(input))
	for i, line := range input {
		values := strings.Split(line, ", ")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		points[i] = Point{x, y}
	}
	grid := Grid(points)
	areas := Area(grid)
	for x, y := range areas {
		fmt.Println(x, y)
	}
}
