package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sumValues(numbers []int) int {
	// First star solution
	total := 0
	for _, i := range numbers {
		total += i
	}
	return total
}

func findFirstReoccurence(numbers []int) int {
	// Second star solution
	// The problem implies that we may need to loop through the values multiple times.
	// In my case it took the 141st pass thorugh the array (the 146545th "index")
	frequencies := make(map[int]int)
	index := 0
	total := 0
	count := len(numbers)
	for {
		total += numbers[index%count]
		frequencies[total]++
		if frequencies[total] == 2 {
			break
		}
		index++
	}
	return total
}

func main() {
	fp, err := os.Open("input.txt")
	values := make([]int, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}
	fmt.Println(sumValues(values))
	fmt.Println(findFirstReoccurence(values))
}
