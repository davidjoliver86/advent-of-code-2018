package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const boardHeight = 1000
const boardWidth = 1000
const claimRegex = "#([[:digit:]]+) @ ([[:digit:]]+),([[:digit:]]+): ([[:digit:]]+)x([[:digit:]]+)"

var re = regexp.MustCompile(claimRegex)
var board = [boardHeight][boardWidth]int{}

type claim struct {
	id, x, y, width, height int
}

func claimArea(board *[boardHeight][boardWidth]int, claim claim) {
	// rows
	for row := claim.y; row < (claim.y + claim.height); row++ {
		for column := claim.x; column < (claim.x + claim.width); column++ {
			board[row][column]++
		}
	}
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	// populate 1000x1000 board with claims
	// we need to track claims in an array for the second star after having populated the board
	claims := []claim{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		values := re.FindStringSubmatch(scanner.Text())
		id, _ := strconv.Atoi(values[1])
		x, _ := strconv.Atoi(values[2])
		y, _ := strconv.Atoi(values[3])
		width, _ := strconv.Atoi(values[4])
		height, _ := strconv.Atoi(values[5])
		claim := claim{id, x, y, width, height}
		claims = append(claims, claim)
		claimArea(&board, claim)
	}

	// First star - how many squares are occupied by two or more claims? (116920)
	multipleClaims := 0
	for row := 0; row < boardHeight; row++ {
		for column := 0; column < boardWidth; column++ {
			if board[row][column] > 1 {
				multipleClaims++
			}
		}
	}
	fmt.Println("Squares with multiple claims:", multipleClaims)

	// Second star - which claim ID has no overlaps?
	// Just go through the claims and find the one that's all '1's
	for _, claim := range claims {
		found := true
		for row := claim.y; row < (claim.y + claim.height); row++ {
			for column := claim.x; column < (claim.x + claim.width); column++ {
				if found && board[row][column] > 1 {
					found = false
				}
			}
		}
		if found {
			fmt.Println("Claim ID", claim.id, "is untouched")
		}
	}
}
