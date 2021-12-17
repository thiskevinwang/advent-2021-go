package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInputList(byteArray []byte) []string {
	return strings.Split(strings.Trim(string(byteArray), "\n"), "\n")
}

type LineSegment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (ls LineSegment) getOrientation() string {
	res := "diagonal"
	if ls.x1 == ls.x2 {
		res = "vertical"
	}
	if ls.y1 == ls.y2 {
		res = "horizontal"
	}
	return res

}

func (ls LineSegment) getXRange() []int {
	delta := ls.x2 - ls.x1
	dir := 1

	if delta < 0 {
		dir = -1
	}

	xs := []int{}
	for a := ls.x1; a != ls.x2+dir; a += dir {
		xs = append(xs, a)
	}
	return xs
}

func (ls LineSegment) getYRange() []int {
	delta := ls.y2 - ls.y1
	dir := 1

	if delta < 0 {
		dir = -1
	}

	ys := []int{}
	for a := ls.y1; a != ls.y2+dir; a += dir {
		ys = append(ys, a)
	}
	return ys
}

// convert strings to structs
func getlineSegments(inputList []string) []LineSegment {
	lineSegments := []LineSegment{}

	for _, inputLine := range inputList {

		numStrings := regexp.MustCompile(`,| -> `).Split(inputLine, 4)

		coords := [4]int{}
		for i, coord := range numStrings {
			num, _ := strconv.Atoi(coord)
			coords[i] = num
		}

		x1 := coords[0]
		y1 := coords[1]
		x2 := coords[2]
		y2 := coords[3]

		lineSegment := LineSegment{
			x1, y1, x2, y2,
		}

		lineSegments = append(lineSegments, lineSegment)

	}
	return lineSegments
}

func main() {
	byteArray, _ := os.ReadFile("inputs.txt")
	inputList := getInputList(byteArray)
	lineSegments := getlineSegments(inputList)

	// establish matrix bounds
	maxX := math.MinInt
	maxY := math.MinInt

	for _, ls := range lineSegments {
		maxX = int(math.Max(float64(maxX), float64(ls.x1)))
		maxX = int(math.Max(float64(maxX), float64(ls.x2)))
		maxY = int(math.Max(float64(maxY), float64(ls.y1)))
		maxY = int(math.Max(float64(maxY), float64(ls.y2)))
	}

	{
		// part 1
		matrix := make([][]int, maxY)
		for i := range matrix {
			matrix[i] = make([]int, maxX)
		}

		for _, ls := range lineSegments {
			xRange := ls.getXRange()
			yRange := ls.getYRange()
			// fmt.Println(xRange, yRange)

			switch ls.getOrientation() {
			case "horizontal":
				for _, x := range xRange {
					matrix[ls.y1][x] += 1
				}
			case "vertical":
				for _, y := range yRange {
					matrix[y][ls.x1] += 1
				}
			}
		}

		count := 0

		for _, row := range matrix {
			for _, item := range row {
				if item > 1 {
					count += 1
				}
			}
		}

		fmt.Println("Part1: Count:", count)
	}

	{
		// part 2
		matrix := make([][]int, len(inputList)*2)
		for i := range matrix {
			matrix[i] = make([]int, len(inputList)*2)
		}

		for _, ls := range lineSegments {
			xRange := ls.getXRange()
			yRange := ls.getYRange()

			switch ls.getOrientation() {
			case "horizontal":
				for _, x := range xRange {
					matrix[ls.y1][x] += 1
				}
			case "vertical":
				for _, y := range yRange {
					matrix[y][ls.x1] += 1
				}
			case "diagonal":
				for i, y := range yRange {
					x := xRange[i]
					matrix[y][x] += 1
				}
			}
		}

		count := 0

		for _, row := range matrix {
			for _, item := range row {
				if item > 1 {
					count += 1
				}
			}
		}

		fmt.Println("Part2: Count:", count)
	}

}
