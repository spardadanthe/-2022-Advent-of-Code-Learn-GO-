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

	charsBeforeMarker := findCharsBeforeMarker(input, 4)
	charsBeforeWord := findCharsBeforeMarker(input, 14)
	fmt.Println(charsBeforeMarker)
	fmt.Println(charsBeforeWord)
}

func findCharsBeforeMarker(input string, uniqueChars int) int {
	markerHolder := ""

	for i, current := range input {
		res := strings.Index(markerHolder, string(current))

		if res != -1 {
			markerHolder = markerHolder[res+1:]
		}
		markerHolder = markerHolder + string(current)

		if len(markerHolder) == 14 {
			return i + 1
		}
	}

	return 0
}
