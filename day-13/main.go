package main

import "fmt"

func main() {
	// First star
	track := NewTrack("input.txt")
	fleet := NewFleet("input.txt")

	// Move until theres a crash
	var crashes []Coordinates
	numCrashes := 0
	for numCrashes == 0 {
		fleet.MoveAll(&track)
		crashes = fleet.Crashes()
		numCrashes = len(crashes)
		fmt.Println(fleet)
	}
	// actualCrashX, actualCrashY := crashes[0].X, crashes[0].Y
	fmt.Println(crashes)
}
