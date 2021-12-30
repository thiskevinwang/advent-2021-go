package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	byteArray, _ := os.ReadFile("inputs.txt")

	lanternFishInit := strings.Split(strings.Trim(string(byteArray), "\n"), ",")

	res := []int{}

	// initial lantern fish array
	for _, str := range lanternFishInit {
		i, _ := strconv.Atoi(str)
		res = append(res, i)
	}

	// iterate through 80 days
	for d := 0; d < 80; d++ {
		// update all fish counters
		newFishes := []int{}

		for i, fish := range res {
			// if 0, reset to 6 and append an 8
			if fish == 0 {
				res[i] = 6
				newFishes = append(newFishes, 8)
			} else {
				// else decrement
				res[i] -= 1
			}

		}

		res = append(res, newFishes...)
	}

	fmt.Println("Part 1:", len(res))
}
