package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Convenience struct for star #1
type counts struct {
	two   int
	three int
}

func countTwosAndThrees(charMaps map[string](map[byte]int)) counts {
	// Iterates over the character count maps we built from the BoxIDs.
	// Each BoxID can have a pair of two identical characters, and/or a triplet of three.
	// Multiple pairs/triplets are disregarded.
	//
	// abcdef contains no letters that appear exactly two or three times.
	// bababc contains two a and three b, so it counts for both.
	// abbcde contains two b, but no letter appears exactly three times.
	// abcccd contains three c, but no letter appears exactly two times.
	// aabcdd contains two a and two d, but it only counts once.
	// abcdee contains two e.
	// ababab contains three a and three b, but it only counts once.

	totalCounts := counts{}
	for _, charMap := range charMaps {
		hasTwo := false
		hasThree := false
		for _, charCount := range charMap {
			if charCount == 2 {
				hasTwo = true
			}
			if charCount == 3 {
				hasThree = true
			}
		}
		if hasTwo {
			totalCounts.two++
		}
		if hasThree {
			totalCounts.three++
		}
	}
	return totalCounts
}

func findCommonChars(s1, s2 string) []byte {
	// Returns a list of characters that are positionally identical between the two strings.
	// It is known, and thus assumed, that s1 and s2 are the same length.

	commonChars := make([]byte, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			commonChars = append(commonChars, s1[i])
		}
	}
	return commonChars
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	// For the second star solution it's important that we iterate through the box IDs in a fixed order.
	// Iterating through the keys in a map does not guarantee preservation of order.
	// So we'll just store the box IDs in an array.

	boxIDs := make([]string, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		boxIDs = append(boxIDs, scanner.Text())
	}

	// First star
	// Generate maps of character counts for each BoxID, then find how many of those boxIDs contain pairs and triplets.
	// Solution is (#pairs) * (#triplets). (26 * 249 = 6474)

	allCounts := make(map[string](map[byte]int))
	for _, boxID := range boxIDs {
		boxIDCharCounts := make(map[byte]int)
		for i := 0; i < len(boxID); i++ {
			boxIDCharCounts[boxID[i]]++
		}
		allCounts[boxID] = boxIDCharCounts
	}
	twosAndThrees := countTwosAndThrees(allCounts)
	fmt.Println("Box checksum is", twosAndThrees.three, "*", twosAndThrees.two, "=", twosAndThrees.three*twosAndThrees.two)

	// Second star
	// Iterate through the BoxIDs to find a pair of BoxIDs where only one character differs between them.
	// Solution is the sequence of characters that are positionally identical. (mxhwoglxgeauywfkztndcvjqr)

	for i := 0; i < len(boxIDs)-1; i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			commonChars := findCommonChars(boxIDs[i], boxIDs[j])
			if len(commonChars) == len(boxIDs[i])-1 {
				for _, c := range commonChars {
					fmt.Printf("%c", c)
				}
				fmt.Println()
			}
		}
	}
}
