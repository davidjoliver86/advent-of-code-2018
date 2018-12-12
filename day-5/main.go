package main

import (
	"fmt"
	"strings"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

const caseDifference = 'a' - 'A'

// StripPolarity removes adjacent pairs of letters from a string where the letters are the same
// but the capitalization differs.
func StripPolarity(input string) string {
	for i := 0; i < len(input)-1; i++ {
		var diff byte
		this := input[i]
		next := input[i+1]
		if this > next {
			diff = this - next
		} else {
			diff = next - this
		}
		if diff == caseDifference {
			return StripPolarity(input[:i] + input[i+2:])
		}
	}
	return input
}

// StripUnitType removes a given character from the polymer, then runs StripPolarty on the result.
// unit must be lowercase
func StripUnitType(input string, unit rune) string {
	stripper := func(r rune) rune {
		if r == unit || (r+caseDifference) == unit {
			return -1
		}
		return r
	}
	result := StripPolarity(strings.Map(stripper, input))
	fmt.Printf("Stripping %c resulting length is %d\n", unit, len(result))
	return result
}

func main() {
	// First star: find the length of the input polymer (11118)
	originalPolymer := util.FileLines("input.txt")[0]
	stripped := StripPolarity(originalPolymer)
	fmt.Println("Length of the remaining polymer:", len(stripped))

	// Second star: find the shortest polymer given that we run StripPolymer on each letter
	results := [26]int{}
	for unit := 'a'; unit <= 'z'; unit++ {
		results[unit-'a'] = len(StripUnitType(originalPolymer, unit))
	}
	lowest := util.Min(results[:])
	fmt.Println("Shortest polymer from stripping units:", lowest)
}
