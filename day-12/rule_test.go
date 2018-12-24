package main

import "testing"

var expected = Rules{
	{false, false, false, true, true}:  true,
	{false, false, true, false, false}: true,
	{false, true, false, false, false}: true,
	{false, true, false, true, false}:  true,
	{false, true, false, true, true}:   true,
	{false, true, true, false, false}:  true,
	{false, true, true, true, true}:    true,
	{true, false, true, false, true}:   true,
	{true, false, true, true, true}:    true,
	{true, true, false, true, false}:   true,
	{true, true, false, true, true}:    true,
	{true, true, true, false, false}:   true,
	{true, true, true, false, true}:    true,
	{true, true, true, true, false}:    true,
}

var ruleInput = []string{
	"...## => #",
	"..#.. => #",
	".#... => #",
	".#.#. => #",
	".#.## => #",
	".##.. => #",
	".#### => #",
	"#.#.# => #",
	"#.### => #",
	"##.#. => #",
	"##.## => #",
	"###.. => #",
	"###.# => #",
	"####. => #",
}

func TestParseRule(t *testing.T) {
	for _, input := range ruleInput {
		rule, result := ParseRule(input)
		if expected[rule] != result {
			t.Errorf("bad")
		}
	}
}
