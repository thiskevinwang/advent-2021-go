package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var byteArray string

func main() {
	lanternFishInit := strings.Split(strings.Trim(string(byteArray), "\n"), ",")

	{ // part 1
		// This is a brute force approach
		// - `res` is a variable-size SLICE that will grow as new lantern fish are spawned, and appended to it.
		// - Each element in `res` represents the countdown timer for a single fish.
		//
		// Step 1:
		// 	- do an 80-day for-loop
		// Step 1a:
		// 	- within each for loop, do a nested for-loop on `res`
		// 	- countdown timers
		//	- reset any 0's to 6's
		//      - accumulate new fishes (8's)
		//      - append new fishes to `res`
		//
		// This quickly becomes too slow and will take too long for a 256-day for-loop
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

	{ // part 2
		// This is an optimized algorithm;
		// (source: https://github.com/alexchao26/advent-of-code-go/blob/main/2021/day06/main.go#L36-L66)
		// - `counters` is a fixed-size ARRAY
		// - each element represents the total number of lantern fish with the same counter
		//   - array index corresponds with countdown timer (0..8)
		//
		// Step 1:
		// 	- do a 256-day for-loop
		//	- take note [0] item
		//	  - this is the number of lantern fish that will
		//	    - reset their timers to 6-days
		//	    - "spawn" new lantern fish with 8-day-timers
		//	- shift all elements to the "left"; wrap [0] to [8]
		//	- [9, 5, 3, 0, 0, 0, 0, 0, 0]
		//	  becomes
		//	  [5, 3, 0, 0, 0, 0, 0, 0, 9]
		//      - increment [6] by [0]
		counters := [9]int{0}

		// initial lantern fish array
		for _, str := range lanternFishInit {
			i, _ := strconv.Atoi(str)
			counters[i] += 1
		}

		// iterate through 256 days
		for d := 0; d < 256; d++ {
			save := counters[0]
			for i := 0; i < len(counters)-1; i++ {
				counters[i] = counters[i+1]
			}
			counters[8] = save
			counters[6] += save
		}

		sum := 0
		for _, n := range counters {
			sum += n
		}
		fmt.Println("Part 2:", sum)
	}
}
