package main

import "fmt"

const gridSize int = 300
const gridSerial int = 2866

type power struct {
	x, y, size, power int
}

func PowerLevel(x, y, serial int) int {
	// Find the fuel cell's rack ID, which is its X coordinate plus 10.
	// Begin with a power level of the rack ID times the Y coordinate.
	// Increase the power level by the value of the grid serial number (your puzzle input).
	// Set the power level to itself multiplied by the rack ID.
	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	// Subtract 5 from the power level.
	rackID := x + 10

	powerLevel := rackID * y
	powerLevel += serial
	powerLevel *= rackID

	hundreds := (powerLevel / 100) % 10
	return hundreds - 5
}

func NewGrid() [gridSize][gridSize]int {
	grid := [gridSize][gridSize]int{}
	for y := range grid {
		for x := range grid {
			grid[y][x] = PowerLevel(x+1, y+1, gridSerial)
		}
	}
	return grid
}

func SquarePower(x, y, size int, g *[gridSize][gridSize]int) int {
	pow := 0
	for row := y; row < (y + size); row++ {
		for col := x; col < (x + size); col++ {
			pow += g[row][col]
		}
	}
	return pow
}

// MaxPower.. he's the man who's name you'd love to touch.. but you mustn't touch...
func MaxPower(size int, g *[gridSize][gridSize]int) power {
	maxPow := 0
	maxPowX := 0
	maxPowY := 0
	for y := 0; y < (gridSize - (size - 1)); y++ {
		for x := 0; x < (gridSize - (size - 1)); x++ {
			// sum power level of the 3x3 block where x, y is the top-left coordinate
			pow := SquarePower(x, y, size, g)
			if pow > maxPow {
				maxPow = pow
				maxPowX = x + 1
				maxPowY = y + 1
			}
		}
	}
	return power{maxPowX, maxPowY, size, maxPow}
}

func main() {
	grid := NewGrid()
	p3 := MaxPower(3, &grid)
	fmt.Println("Highest power level of", p3.power, "found at", p3.x, p3.y)

	highestPower := 0
	winner := power{0, 0, 0, 0}
	for size := 1; size <= gridSize; size++ {
		sizePower := MaxPower(size, &grid)
		if sizePower.power > highestPower {
			fmt.Println("Current high score:", sizePower.power)
			highestPower = sizePower.power
			winner = sizePower
		}
	}
	fmt.Println("Highest power level found in a", winner.size, "x", winner.size, "grid at", winner.x, ",", winner.y)
}
