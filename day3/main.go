package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thiskevinwang/advent-2021-go/utils"
)

// https://adventofcode.com/2021/day/3
func main() {
	// var binaryStrings []string = utils.GetInputs()
	inputs := utils.FetchInputs("https://adventofcode.com/2021/day/3/input")
	binaryStrings := strings.Split(inputs, "\n")

	// generate an array like [-10, 1, -5, 16]
	// - positive vals signify "1" was the most common digit at an index
	// - negative vals signify "0" was the most common digit at an index
	// - 0 val signifies that there were equal occurences of each.
	bitOccurences := getBitOccurences(binaryStrings)

	g, e := getRates(bitOccurences)

	fmt.Println("Part1: Gamma:", g, "Epsilon:", e, "Result:", g*e)

	o2Arr := binaryStrings
	co2Arr := binaryStrings

	// For each index in the width of the input,
	// continually filter down entries
	for i := range binaryStrings[0] {
		if len(o2Arr) > 1 {
			o2BitOccurences := getBitOccurences(o2Arr)
			newO2Arr := []string{}

			for _, bits := range o2Arr {
				if o2BitOccurences[i] >= 0 && bits[i] == '1' {
					newO2Arr = append(newO2Arr, bits)
				} else if o2BitOccurences[i] < 0 && bits[i] == '0' {
					newO2Arr = append(newO2Arr, bits)
				}
			}

			o2Arr = newO2Arr
		}

		if len(co2Arr) > 1 {
			co2BitOccurences := getBitOccurences(co2Arr)
			newCo2Arr := []string{}

			for _, bits := range co2Arr {
				if co2BitOccurences[i] >= 0 && bits[i] == '0' {
					newCo2Arr = append(newCo2Arr, bits)
				} else if co2BitOccurences[i] < 0 && bits[i] == '1' {
					newCo2Arr = append(newCo2Arr, bits)
				}
			}

			co2Arr = newCo2Arr
		}
	}

	o2, _ := strconv.ParseInt(o2Arr[0], 2, 64)
	co2, _ := strconv.ParseInt(co2Arr[0], 2, 64)

	fmt.Println("Part2: o2:", o2, "Co2:", co2, "Lift support:", o2*co2)
}

func getBitOccurences(binaryStrings []string) []int {
	size := len(binaryStrings[0])

	// initialize array of 0's, with len = the first line
	bitOccurences := make([]int, size)

	// O(n * m)
	for _, entry := range binaryStrings {
		for i, bit := range entry {
			switch bit {
			case '1':
				bitOccurences[i] += 1
			case '0':
				bitOccurences[i] -= 1
			}
		}
	}
	return bitOccurences
}

// build binary strings & return the decimal values
func getRates(bitOccurences []int) (int, int) {
	gammaString := ""
	epsilonString := ""

	for _, i := range bitOccurences {
		if i >= 0 {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)

	// https://go.dev/ref/spec#Conversions
	return int(gamma), int(epsilon)
}
