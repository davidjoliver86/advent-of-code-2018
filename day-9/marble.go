package main

import "fmt"

// Marble is a circlar doubly-linked list
type Marble struct {
	Value int
	Next  *Marble
	Prev  *Marble
}

func (m *Marble) String() string {
	return fmt.Sprintf("Marble[%d]", m.Value)
}

// Insert adds itself to the existing marble chain. Special rules apply if the marble value
// is a multiple of 23. In either case, it returns the "current marble".
func (m *Marble) Insert(value int, inserter *Player) *Marble {
	// if newValue % 23 == 0:
	// 1) do not insert this marble
	// 2) remove the marble 7 previous marbles away
	// 3) add newValue + the removed marble to the player's score
	// 4) return the 'next' marble from that removed marble in step 2
	if value%23 == 0 {
		removed := m.Prev.Prev.Prev.Prev.Prev.Prev.Prev
		p := removed.Prev
		n := removed.Next
		p.Next = n
		n.Prev = p
		inserter.Score += (value + removed.Value)
		return removed.Next
	}

	// otherwise add it in between one and two spots ahead
	one := m.Next
	two := m.Next.Next
	new := &Marble{Value: value}
	one.Next = new
	new.Next = two
	two.Prev = new
	new.Prev = one
	return new
}

// NewGenesisMarble returns a new marble with value 0 and references to itself
func NewGenesisMarble() *Marble {
	m := &Marble{Value: 0}
	m.Next = m
	m.Prev = m
	return m
}
