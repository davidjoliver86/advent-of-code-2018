package main

// RuleSize is the number of characters that comprise a rule
const RuleSize = 5

// Rules is the mapping of a sequence of RuleSize booleans to determine whether the middle space should or
// should not contain a plant
type Rules map[[RuleSize]bool]bool

// ParseRule takes rules from input in the form of ".#.#. => #" and translates them to booleans
func ParseRule(input string) ([RuleSize]bool, bool) {
	rule := [RuleSize]bool{}
	for i := 0; i < RuleSize; i++ {
		val, _ := Translate(input[i])
		rule[i] = val
	}
	result, _ := Translate(input[RuleSize+4])
	return rule, result
}
