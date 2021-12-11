package main

import (
	"fmt"
	"main/utils"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var inputs string = utils.FetchInputs("https://adventofcode.com/2021/day/4/input")
	cleanedInput := strings.Split(strings.Trim(inputs, "\n"), "\n\n")

	numStrings := cleanedInput[0]
	boardStrings := cleanedInput[1:]

	boards := []Board{}

	for _, boardString := range boardStrings {
		boards = append(boards, createBoard(boardString))
	}

	scores := []int{}

	for _, num := range strings.Split(numStrings, ",") {
		n, _ := strconv.Atoi(num)

		// Structs are copied in range loops. You need to access by index.
		for bIdx := range boards {
			if boards[bIdx].Bingo() {
				continue
			}

			boards[bIdx].Check(n)

			if boards[bIdx].Bingo() {
				boardScore := boards[bIdx].GetScore()
				total := n * boardScore
				scores = append(scores, total)

			}
		}
	}

	fmt.Println("Part1: Winning score:", scores[0])
	fmt.Println("Part2: Winning score:", scores[len(scores)-1])
}

type Board struct {
	checks [5][5]bool
	nums   [5][5]int
}

func createBoard(s string) Board {
	b := Board{}

	rows := strings.Split(strings.Trim(s, "\n"), "\n")

	// whitespace regex matcher
	re := regexp.MustCompile(`\s+`)

	for i, row := range rows {
		// trim whitespace
		row := strings.Trim(row, " ")
		// split by whitespace of varying size
		nums := re.Split(row, -1)

		for j, val := range nums {
			num, _ := strconv.Atoi(val)
			b.nums[i][j] = num
		}
	}

	return b
}

// Note: https://go.dev/tour/methods/4
// (b Board) "Value receiver" —  operates on a copy of the original value
// (b *Board) "Pointer receiver" — Methods with pointer receivers can modify the value to which the receiver points
func (b *Board) Check(val int) {
out:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.nums[i][j] == val {
				b.checks[i][j] = true
				break out
			}
		}
	}
}

// return if the board has bingo
func (b Board) Bingo() bool {
	isBingo := false
	// check rows
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.checks[i][j] {
				break
			}
			if j == 4 {
				isBingo = true
			}
		}
	}

	if isBingo {
		return isBingo
	}

	// check cols
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.checks[j][i] {
				break
			}
			if j == 4 {
				isBingo = true
			}
		}
	}

	return isBingo
}

// get sum of unchecked cells
func (b Board) GetScore() int {
	sum := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.checks[i][j] {
				sum += b.nums[i][j]
			}
		}
	}

	return sum
}
