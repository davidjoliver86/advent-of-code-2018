package util

import (
	"bufio"
	"log"
	"os"
)

func FileLines(filepath string) []string {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Max(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max
}

func Min(arr []int) int {
	min := arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
