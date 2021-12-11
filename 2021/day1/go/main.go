package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const Url = "https://adventofcode.com/2021/day/1/input"
const SessionTokenKey = "AOC_SESSION_TOKEN"

// https://adventofcode.com/2021/day/1
func main() {
	inputList := fetchInputs()

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

// fetch inputs from advent of code
func fetchInputs() []int {
	// session=...
	sessionToken, ok := os.LookupEnv(SessionTokenKey)
	if !ok {
		fmt.Println(SessionTokenKey, `is missing.
	You can grab one from the dev tools from the
	"session=..." cookie.
	- Then run export AOC_SESSION_TOKEN=<token>
	- Or :let $AOC_SESSION_TOKEN=<token>, from vim`)
	}

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionToken})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	stringList := strings.Split(string(bodyBytes), "\n")

	return mapStringToInt(stringList)
}
