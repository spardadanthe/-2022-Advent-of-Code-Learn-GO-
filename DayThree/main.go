package main

import (
	"filereader"
	"fmt"
)

func main() {
	fmt.Println("test")
	lines := filereader.ReadFile("input.txt")
	counter := 0

	for _, line := range lines {
		counter += getRepeatedItemNumber(line)
	}

	fmt.Println("==========================")
	fmt.Print("Final Answer is: ")
	fmt.Println(counter)
}

func getRepeatedItemNumber(line string) int {
	length := len(line)

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := 0; i < length/2; i++ {
		for j := length / 2; j < length; j++ {
			if line[i] == line[j] {
				repeatedItemNumber := getNumberOfItem(line[i])
				fmt.Printf("%v -> %v \n", line[i], repeatedItemNumber)

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

// nhgPFmsnLPGLhPDJhG
// TcDjMfrMMjMZWfjfWj
