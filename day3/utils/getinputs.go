package utils

import (
	"os"
	"strings"
)

// use local file
func GetInputs() []string {
	byteArray, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("./input.txt doesn't exist")
	}
	return strings.Split(string(byteArray), "\n")
}
