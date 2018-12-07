package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

var data = regexp.MustCompile(":([0-9]{2})] (.*)")
var guardBeginsShift = regexp.MustCompile("#([0-9]+) begins shift")
var sleepSchedule = make(map[int]([]shift))

type shift struct {
	Schedule   [60]int
	sleepStart int
}

func (s *shift) Sleep(minuteSleep int) {
	s.sleepStart = minuteSleep
}

func (s *shift) Wakeup(minuteWokenUp int) {
	for i := s.sleepStart; i < minuteWokenUp; i++ {
		s.Schedule[i] = 1
	}
	s.sleepStart = 0
}

func minutesSlept(shifts []shift) int {
	total := 0
	for _, shift := range shifts {
		for minute := 0; minute < len(shift.Schedule); minute++ {
			total += shift.Schedule[minute]
		}
	}
	return total
}

func minuteMap(shifts []shift) [60]int {
	minuteMap := [60]int{}
	for _, shift := range shifts {
		for minute := 0; minute < len(minuteMap); minute++ {
			minuteMap[minute] += shift.Schedule[minute]
		}
	}
	return minuteMap
}

func mostFrequentMinute(shifts []shift) int {
	minuteMap := minuteMap(shifts)
	minuteIndex := 0
	minuteValue := minuteMap[minuteIndex]
	for minute := 0; minute < len(minuteMap); minute++ {
		if minuteMap[minute] > minuteValue {
			minuteIndex = minute
			minuteValue = minuteMap[minute]
		}
	}
	return minuteIndex
}

func lastShift(guardID int) *shift {
	return &sleepSchedule[guardID][len(sleepSchedule[guardID])-1]
}

func main() {
	lines := util.FileLines("input.txt")
	sort.StringSlice(lines).Sort()
	var guardID int
	for _, line := range lines {
		values := data.FindStringSubmatch(line)
		minute, _ := strconv.Atoi(values[1])
		action := values[2]
		if guardBeginsShift.MatchString(action) {
			guardID, _ = strconv.Atoi(guardBeginsShift.FindStringSubmatch(line)[1])
			sleepSchedule[guardID] = append(sleepSchedule[guardID], shift{})
		}
		if strings.Contains(line, "falls asleep") {
			lastShift(guardID).Sleep(minute)
		}
		if strings.Contains(line, "wakes up") {
			lastShift(guardID).Wakeup(minute)
		}
	}
	for guardID, shifts := range sleepSchedule {
		minutes := minuteMap(shifts)
		minutesSlept := minutesSlept(shifts)
		mostFrequentMinute := mostFrequentMinute(shifts)
		fmt.Println(guardID, "has slept", minutesSlept, "most frequently on minute", mostFrequentMinute, "=>", util.Max(minutes[:]), "times")
	}
}
