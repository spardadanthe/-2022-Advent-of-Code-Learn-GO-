package main

import (
	"filereader"
	"fmt"
	"strings"
	// "regexp"
	// "strconv"
	// "strings"
)

func main() {
	inputArray := filereader.ReadFile("input.txt")
	input := inputArray[0]

	charsBeforeMarker := findCharsBeforeMarker(input)
	fmt.Println(charsBeforeMarker)
}

func findCharsBeforeMarker(input string) int {
	markerHolder := ""

	for i, current := range input {
		res := strings.Index(markerHolder, string(current))

		if res != -1 {
			markerHolder = markerHolder[res+1:]
		}
		markerHolder = markerHolder + string(current)

		if len(markerHolder) == 4 {
			return i + 1
		}
	}

	return 0
}
