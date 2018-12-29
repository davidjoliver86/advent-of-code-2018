package main

import "testing"

func TestCrashes(t *testing.T) {
	var crashes []Coordinates
	expectedCrashX, expectedCrashY := 7, 3
	track := NewTrack("sample_track.txt")
	fleet := NewFleet("sample_track.txt")
	numCrashes := 0
	for numCrashes == 0 {
		fleet.MoveAll(&track)
		crashes = fleet.Crashes()
		numCrashes = len(crashes)
	}
	actualCrashX, actualCrashY := crashes[0].X, crashes[0].Y
	if expectedCrashX != actualCrashX || expectedCrashY != actualCrashY {
		t.Errorf("First crash expected to occur at (%d, %d) but occured at (%d, %d)", expectedCrashX, expectedCrashY, actualCrashX, actualCrashY)
	}
}
