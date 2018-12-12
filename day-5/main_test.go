package main

import "testing"

type testcaseStripPolarity struct {
	input, expected string
}

type testcaseStripUnitType struct {
	input    string
	unit     rune
	expected string
}

var testcasesStripPolarity = []testcaseStripPolarity{
	{"aA", ""},
	{"abBA", ""},
	{"abAB", "abAB"},
	{"aabAAB", "aabAAB"},
	{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
}

var testcasesStripUnitType = []testcaseStripUnitType{
	{"dabAcCaCBAcCcaDA", 'a', "dbCBcD"},
	{"dabAcCaCBAcCcaDA", 'b', "daCAcaDA"},
	{"dabAcCaCBAcCcaDA", 'c', "daDA"},
	{"dabAcCaCBAcCcaDA", 'd', "abCBAc"},
}

func TestStripPolarity(t *testing.T) {
	for _, testcase := range testcasesStripPolarity {
		actual := StripPolarity(testcase.input)
		t.Logf("StripPolarity(\"%s\") => \"%s\"", testcase.input, actual)
		if actual != testcase.expected {
			t.Errorf("Expected \"%s\" but got \"%s\"!", testcase.expected, actual)
		}
	}
}

func TestStripUnitType(t *testing.T) {
	for _, testcase := range testcasesStripUnitType {
		actual := StripUnitType(testcase.input, testcase.unit)
		t.Logf("StripUnitType(\"%s\", \"%c\") => \"%s\"", testcase.input, testcase.unit, actual)
		if actual != testcase.expected {
			t.Errorf("Expected \"%s\" but got \"%s\"!", testcase.expected, actual)
		}
	}
}
