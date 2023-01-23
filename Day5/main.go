package main

import (
	"filereader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	input := filereader.ReadFile("input.txt")
	warehouse := parseWarehouseInput(input)

	warehouseCopy := make(map[int]string)
	for k, v := range warehouse {
		warehouseCopy[k] = v
	}

	fmt.Println(warehouse)
	warehouse = moveBoxes(input, 10, warehouse, false)
	warehouseWithNewCrane := moveBoxes(input, 10, warehouseCopy, true)
	printResult(warehouse)
	printResult(warehouseWithNewCrane)
}

func parseWarehouseInput(input []string) map[int]string {
	result := map[int]string{}
	emptyStringMarker := "temp"
	length := 0
	for emptyStringMarker != "" {
		emptyStringMarker = input[length]
		length++
	}
	startingIndex := length - 2

	trimmedFirstRow := strings.ReplaceAll(input[startingIndex], " ", "")
	fmt.Println(trimmedFirstRow)

	for _, letter := range trimmedFirstRow {
		runeAsString := string(letter)
		runeAsInt, _ := strconv.Atoi(runeAsString)

		result[runeAsInt] = ""
	}

	for i := startingIndex - 1; i >= 0; i-- {
		counter := 1
		for j := 0; j < len(input[i]); j = j + 4 {
			currentLetter := string(input[i][j+1])
			if currentLetter != " " {
				result[counter] += currentLetter
			}
			counter++
		}
	}

	return result
}

func moveBoxes(input []string, startingIndex int, warehouse map[int]string, isNewCrane bool) map[int]string {
	for i := startingIndex; i < len(input); i++ {
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		directionsAsStringSlice := re.FindAllString(input[i], -1)
		directions := stringSliceToIntSlice(directionsAsStringSlice)

		if isNewCrane {
			warehouse = moveBoxWithImprovedCrane(warehouse, directions)
		} else {
			warehouse = moveBox(warehouse, directions)
		}
	}

	return warehouse
}

func stringSliceToIntSlice(input []string) []int {
	result := make([]int, len(input))

	for i, s := range input {
		result[i], _ = strconv.Atoi(s)
	}

	return result
}

func moveBox(warehouse map[int]string, directions []int) map[int]string {
	for i := 0; i < directions[0]; i++ {
		warehouseStack := warehouse[directions[1]]
		warehouseStackLength := len(warehouseStack)
		lastElement := warehouseStack[warehouseStackLength-1:]
		warehouse[directions[1]] = warehouseStack[0 : warehouseStackLength-1]
		warehouse[directions[2]] = warehouse[directions[2]] + lastElement
	}

	return warehouse
}

func moveBoxWithImprovedCrane(warehouse map[int]string, directions []int) map[int]string {
	warehouseStack := warehouse[directions[1]]
	warehouseStackLength := len(warehouseStack)
	elementsToMove := warehouseStack[warehouseStackLength-directions[0]:]
	warehouse[directions[1]] = warehouseStack[0 : warehouseStackLength-directions[0]]
	warehouse[directions[2]] = warehouse[directions[2]] + elementsToMove

	return warehouse
}

func printResult(warehouse map[int]string) {
	fmt.Println("\n Result is -> ")
	for i := 1; i <= len(warehouse); i++ {
		element := warehouse[i]

		fmt.Print(element[len(element)-1:])
	}
}
