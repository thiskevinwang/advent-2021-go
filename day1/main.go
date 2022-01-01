package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thiskevinwang/advent-2021-go/utils"
)

// https://adventofcode.com/2021/day/1
func main() {
	input := utils.FetchInputs("https://adventofcode.com/2021/day/1/input")
	inputList := mapStringToInt(strings.Split(input, "\n"))

	increased := 0
	windowIncreased := 0

	for j := 1; j < len(inputList); j++ {
		i := j - 1
		prev := inputList[i]
		next := inputList[j]

		if next > prev {
			increased++
		}

		// check out of bounds
		if j+2 < len(inputList) {
			prevWindowSum := sum(inputList[i : i+3])
			nextWindowSum := sum(inputList[j : j+3])

			if nextWindowSum > prevWindowSum {
				windowIncreased++
			}
		}
	}

	fmt.Println("Part 1: Increased count:", increased)
	fmt.Println("Part 2: Window Increased count:", windowIncreased)
}

// map array of text to int
func mapStringToInt(array []string) []int {
	res := []int{}
	for _, v := range array {
		n, _ := strconv.Atoi(v)
		res = append(res, n)
	}
	return res
}

// sum array of int
func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
