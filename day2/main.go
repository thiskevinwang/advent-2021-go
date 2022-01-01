package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var byteArray []byte

func getList(byteArray []byte) []string {
	return strings.Split(string(byteArray), "\n")
}

type Pos struct {
	x int
	y int
}

type Pos2 struct {
	x   int
	y   int
	aim int
}

// https://adventofcode.com/2021/day/2
func main() {
	commands := getList(byteArray)
	pos := Pos{x: 0, y: 0}

	for _, command := range commands {
		parts := strings.Split(command, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			pos.x += amount
		case "down":
			pos.y += amount
		case "up":
			pos.y -= amount
		}
	}

	fmt.Println("Part1: Position:", pos.x*pos.y)

	pos2 := Pos2{x: 0, y: 0, aim: 0}
	for _, command := range commands {
		parts := strings.Split(command, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			pos2.x += amount
			pos2.y += pos2.aim * amount
		case "down":
			pos2.aim += amount
		case "up":
			pos2.aim -= amount
		}
	}

	fmt.Println("Part2: Position:", pos2.x*pos2.y)
}
