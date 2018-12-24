package main

import (
	"errors"
	"fmt"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

var rules = Rules{}

// Translate is a common utility function to translate '.' into false and '#' into true
func Translate(b byte) (bool, error) {
	if b == '#' {
		return true, nil
	}
	if b == '.' {
		return false, nil
	}
	return false, errors.New("bad character")
}

func printPlants(plants Plants, min, max int) {
	for i := min; i < max; i++ {
		if plants[i] {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}

func main() {
	input := util.FileLines("input.txt")
	initialState := input[0][15:]
	for _, line := range input[2:] {
		if len(line) == 10 {
			rule, result := ParseRule(line)
			rules[rule] = result
		}
	}

	// first star
	plants := NewPlants(initialState)
	for i := 0; i < 1000; i++ {
		plants = plants.GenerateNext()
		sum := 0
		for key := range plants {
			if plants[key] {
				sum += key
			}
		}
		fmt.Printf("%6d", sum)
		printPlants(plants, -10, 175)
	}

	// second star
	// after 1000 generations, it becomes obvious that the sum always increases by 55 each generation
	// after 1000 generations, sum = 54911
	gens := 50000000000 - 1000
	fmt.Println((55 * gens) + 54911)
}
