package main

import (
	"strings"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

type Track [][]rune

const (
	StraightH     = '-'
	StraightV     = '|'
	TurnSlash     = '/'
	TurnBackslash = '\\'
	Intersection  = '+'
)

func NewTrack(input string) Track {
	track := make(Track, 0)
	lines := util.FileLines(input)
	for _, line := range lines {
		// strip out carts
		trackLine := strings.Map(func(r rune) rune {
			switch {
			case r == '<' || r == '>':
				return StraightH
			case r == 'v' || r == '^':
				return StraightV
			}
			return r
		}, line)
		row := []rune(trackLine)
		track = append(track, row)
	}
	return track
}

func (t *Track) Type(c Coordinates) rune {
	return (*t)[c.Y][c.X]
}
