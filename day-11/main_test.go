package main

import "testing"

type powerLevelTestCase struct {
	x, y, serial, expected int
}

var powerLevelTestCases = []powerLevelTestCase{
	{122, 79, 57, -5},
	{217, 196, 39, 0},
	{101, 153, 71, 4},
}

func TestPowerLevel(t *testing.T) {
	for _, testcase := range powerLevelTestCases {
		actual := PowerLevel(testcase.x, testcase.y, testcase.serial)
		if actual != testcase.expected {
			t.Errorf("(%d, %d)@%d: expected %d, got %d", testcase.x, testcase.y, testcase.serial, testcase.expected, actual)
		}
	}
}
