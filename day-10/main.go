package main

import (
	"fmt"
	"regexp"
	"strconv"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

var source = regexp.MustCompile("position=<\\s*(-?\\d+),\\s+(-?\\d+)> velocity=<\\s*(-?\\d+),\\s+(-?\\d+)>")

// Point is a pair of X,Y coordinates
type Point struct {
	X, Y int
}

// LightPoint is a point of light with an initial position and a set velocity for each "tick" of time
type LightPoint struct {
	Position Point
	Velocity Point
}

// LightPoints is an array of pointers to LightPoints
type LightPoints []*LightPoint

// Move shifts the position of the light point by its velocity
func (l *LightPoint) Move() {
	l.Position.X += l.Velocity.X
	l.Position.Y += l.Velocity.Y
}

func (l *LightPoint) String() string {
	return fmt.Sprintf("position=<%d, %d> velocity=<%d, %d>", l.Position.X, l.Position.Y, l.Velocity.X, l.Velocity.Y)
}

// Boundaries returns the absolute value of the most distant point coordinate.
// We use this later to determine when the points are sufficiently close together that we can begin observing the message.
func (l LightPoints) Boundaries() int {
	max := util.Abs(l[0].Position.X)
	for _, lp := range l {
		bX := util.Abs(lp.Position.X)
		bY := util.Abs(lp.Position.Y)
		if bX > max {
			max = bX
		}
		if bY > max {
			max = bY
		}
	}
	return max
}

// MoveAll is a convenience function to call .Move() on each LightPoint
func (l LightPoints) MoveAll() {
	for _, lp := range l {
		lp.Move()
	}
}

// ParseLightPoint parses the input data and returns a *LightPoint
func ParseLightPoint(s string) *LightPoint {
	matches := source.FindStringSubmatch(s)
	pX, _ := strconv.Atoi(matches[1])
	pY, _ := strconv.Atoi(matches[2])
	vX, _ := strconv.Atoi(matches[3])
	vY, _ := strconv.Atoi(matches[4])
	return &LightPoint{Position: Point{pX, pY}, Velocity: Point{vX, vY}}
}

func view(lps *LightPoints, fov int) {
	bounds := lps.Boundaries()
	view := make([][]bool, fov*2)
	iteration := 0
	for i := range view {
		view[i] = make([]bool, fov*2)
	}
	for {
		lps.MoveAll()
		iteration++
		newbounds := lps.Boundaries()

		if newbounds < fov {
			// clear view
			for i := range view {
				for j := range view {
					view[i][j] = false
				}
			}
			// project coordinates, adding newbounds to ensure no negative indices
			for _, lp := range *lps {
				row := lp.Position.Y + newbounds
				col := lp.Position.X + newbounds
				view[row][col] = true
			}
			// now print
			for row := range view {
				for col := range view {
					if view[row][col] {
						fmt.Printf("X")
					} else {
						fmt.Printf(" ")
					}
				}
				fmt.Printf("\n")
			}
			fmt.Printf("Iteration %d", iteration)
			fmt.Scanf("placeholder to hit enter")
		}
		// once points start to decay outward, there's really no point continuing
		if newbounds > bounds {
			fmt.Println("Points are dissipating!")
			return
		}
		bounds = newbounds
	}
}

func main() {
	lightPointsSource := util.FileLines("input.txt")
	lightPoints := make(LightPoints, len(lightPointsSource))
	for i := range lightPoints {
		lightPoints[i] = ParseLightPoint(lightPointsSource[i])
	}
	// The viewer incorporates both stars' solutions
	view(&lightPoints, 200)
}
