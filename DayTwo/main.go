package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	printTotal()

	//  A -> Rock
	//  B -> Paper
	//  C -> Scissors

	// X -> Rock
	// Y -> Paper
	// Z -> Scissors

	// Score for round -> Score for shape + otcome score

	// Shape scores:
	// X Rock -> 1
	// Y Paper -> 2
	// Z Scissors -> 3

	// Outcome
	// Lose -> 0
	// Draw -> 3
	// Win -> 6

	// Read
	// Give Points for Sign
	//
}

func printTotal() {
	input, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var pointsCounter int = 0
	var pointsCounterForB int = 0

	for scanner.Scan() {
		var text string = scanner.Text()
		pointsCounter += calcPointsPerLine(text)
		pointsCounterForB += calcPointsPerLineForB(text)
	}

	fmt.Printf("%v", pointsCounter)
	fmt.Println()
	fmt.Printf("%v", pointsCounterForB)
}

func calcPointsPerLine(line string) int {
	points := 0

	var splittedArray []string = strings.Split(line, " ")

	if splittedArray[0] == "A" && splittedArray[1] == "X" {
		points += 4
	} else if splittedArray[0] == "A" && splittedArray[1] == "Y" {
		points += 8
	} else if splittedArray[0] == "A" && splittedArray[1] == "Z" {
		points += 3
	} else if splittedArray[0] == "B" && splittedArray[1] == "X" {
		points += 1
	} else if splittedArray[0] == "B" && splittedArray[1] == "Y" {
		points += 5
	} else if splittedArray[0] == "B" && splittedArray[1] == "Z" {
		points += 9
	} else if splittedArray[0] == "C" && splittedArray[1] == "X" {
		points += 7
	} else if splittedArray[0] == "C" && splittedArray[1] == "Y" {
		points += 2
	} else if splittedArray[0] == "C" && splittedArray[1] == "Z" {
		points += 6
	}

	return points
}

func calcPointsPerLineForB(line string) int {
	points := 0

	// X = lose
	// Y = draw
	// Z = win

	var splittedArray []string = strings.Split(line, " ")

	if splittedArray[0] == "A" && splittedArray[1] == "X" {
		points += 3
	} else if splittedArray[0] == "A" && splittedArray[1] == "Y" {
		points += 4
	} else if splittedArray[0] == "A" && splittedArray[1] == "Z" {
		points += 8
	} else if splittedArray[0] == "B" && splittedArray[1] == "X" {
		points += 1
	} else if splittedArray[0] == "B" && splittedArray[1] == "Y" {
		points += 5
	} else if splittedArray[0] == "B" && splittedArray[1] == "Z" {
		points += 9
	} else if splittedArray[0] == "C" && splittedArray[1] == "X" {
		points += 2
	} else if splittedArray[0] == "C" && splittedArray[1] == "Y" {
		points += 6
	} else if splittedArray[0] == "C" && splittedArray[1] == "Z" {
		points += 7
	}

	return points
}
