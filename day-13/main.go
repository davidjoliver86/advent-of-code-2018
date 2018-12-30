package main

import "fmt"

func main() {
	// First star
	track := NewTrack("input.txt")
	fleet := NewFleet("input.txt")

	// Move until theres a crash
	crashed := false
	for {
		crashed = fleet.MoveAll(&track)
		if crashed {
			break
		}
	}
	// actualCrashX, actualCrashY := crashes[0].X, crashes[0].Y
	fmt.Println(fleet.Crashes())
}
