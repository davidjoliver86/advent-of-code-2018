package main

import "sort"

// Plants is map containing which spaces contain a plant
type Plants map[int]bool

// Buffer is the amount of spaces we "pad" out each line to accomoate checking for rules
const Buffer = RuleSize - 1

// Middle refers to the cell in question - that is, after applying the rule, should there be a plant?
const Middle = Buffer / 2

func (p Plants) minMax() (int, int) {
	var keys []int
	for k := range p {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys[0], keys[len(keys)-1]
}

// GenerateNext scans the current generation of plants, applies the appropriate rules, and returns the next generation
func (p Plants) GenerateNext() Plants {
	min, max := p.minMax()
	new := make(Plants)
	for i := (min - Buffer); i <= max; i++ {
		scan := [RuleSize]bool{}
		for el := range scan {
			scan[el] = p[i+el]
		}
		if rules[scan] {
			new[i+Middle] = true
		}
	}
	return new
}

// NewPlants generates a map of Plants based on initial input of #/.
func NewPlants(input string) Plants {
	plants := make(Plants)
	for i, ch := range input {
		plant, _ := Translate(byte(ch))
		if plant {
			plants[i] = true
		}
	}
	return plants
}
