package main

import (
	"filereader"
	"fmt"
)

func main() {
	input := filereader.ReadFile("input.txt")

	counter := findTotalPriority(input)
	badgesPriority := findBadgesTotalPriority(input)

	fmt.Println(counter)
	fmt.Println(badgesPriority)
}

func findTotalPriority(lines []string) int {

	counter := 0
	for _, line := range lines {
		counter += getPriorityOfItemThatRepeats(line)
	}

	return counter
}

func findBadgesTotalPriority(lines []string) int {
	counter := 0
	counterggg := 0
	increaseWith := 3
	for i := 0; i < len(lines); i += increaseWith {
		value := getGroupBadge(lines, i)
		counterggg++
		// fmt.Println("counter -> ", counterggg, " and with common value -> ", value, "and i = ", i)
		counter += value
	}

	return counter
}

func getGroupBadge(lines []string, sentenceNumber int) int {
	for j, letter := range lines[sentenceNumber] {
		for _, letter2 := range lines[sentenceNumber+1] {
			for _, letter3 := range lines[sentenceNumber+2] {
				if letter == letter2 && letter == letter3 {
					return getNumberOfItem(lines[sentenceNumber][j])
				}
			}
		}
	}

	return 0
}

func getPriorityOfItemThatRepeats(line string) int {
	length := len(line)

	for i := 0; i < length/2; i++ {
		for j := length / 2; j < length; j++ {
			if line[i] == line[j] {
				repeatedItemNumber := getNumberOfItem(line[i])
				//fmt.Printf("%v -> %v \n", line[i], repeatedItemNumber)

				return repeatedItemNumber
			}

		}
	}
	return 0
}

func getNumberOfItem(letter byte) int {
	codeInAscii := int(letter)

	if letter >= 65 && letter <= 90 {
		return codeInAscii - 38
	} else {
		return codeInAscii - 96
	}
}
