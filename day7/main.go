package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// given a list of positive integers, find the value
// that all integers will converge on, that results
// in the smallest net value change across all integers (fuel cost)
func Day7(input []int) int {
	// Brute force
	// 1. find min and max: O(n) to find "m"
	// 2. check fuel cost for each position O(m * n)
	// 3. return the lowest cost

	var min = 999999999
	var max = 0
	for _, val := range input {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	fmt.Println("minmax", min, max)

	var lowestCost = 999999999

	for i := min; i <= max; i++ {
		cost := 0
		for _, j := range input {
			diff := j - i
			if diff < 0 {
				diff *= -1
			}

			// part 1
			// cost += diff

			// part 2
			cost += sumN(diff)
		}

		if cost < lowestCost {
			lowestCost = cost
		}

	}
	return lowestCost
}

//go:embed inputs.txt
var input string

func main() {
	// initialize var for inputs actual type
	var inputs []int

	// convert input plaintext to actual type
	stringSlice := strings.Split(input, ",")
	for _, char := range stringSlice {
		intval, err := strconv.Atoi(char)
		if err != nil {
			panic("failed to parse input")
		}
		inputs = append(inputs, intval)
	}

	// call the program
	res := Day7(inputs)

	// print the result
	fmt.Println(res)
}

// part 2
func sumN(n int) int {
	return n * (n + 1) / 2
}
