package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func PrintElfWithHighestCalories(elfs []Elf) {
	var largestElf Elf

	for index, elf := range elfs {
		if elf.calories > largestElf.calories {
			largestElf.number = index
			largestElf.calories = elf.calories
		}
	}

	fmt.Println("Elf with number -> ", largestElf.number, "and with calories -> ", largestElf.calories, " is the one.")
}

func ReadCalories() []Elf {
	input, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var caloriesCounter int = 0
	var elfCounter int = 1

	var elfs = make([]Elf, 0)

	for scanner.Scan() {
		var text string = scanner.Text()
		if text == "" {
			var tempElf Elf
			tempElf.number = elfCounter
			tempElf.calories = caloriesCounter
			elfs = append(elfs, tempElf)
			elfCounter++
			caloriesCounter = 0
			continue
		}

		intVar, _ := strconv.Atoi(scanner.Text())
		caloriesCounter += intVar
	}

	return elfs
}

func PrintTopThreeElfs(elfs []Elf) {
	topThree := make(map[int]Elf)

	topThree[1] = Elf{number: 0, calories: 0}
	topThree[2] = Elf{number: 0, calories: 0}
	topThree[3] = Elf{number: 0, calories: 0}

	for _, elf := range elfs {
		var tempFirst Elf = topThree[1]
		var tempSecond Elf = topThree[2]

		if elf.calories >= topThree[1].calories {
			topThree[1] = Elf{number: elf.number, calories: elf.calories}
			topThree[2] = tempFirst
			topThree[3] = tempSecond
			continue
		} else if elf.calories >= topThree[2].calories {
			topThree[2] = Elf{number: elf.number, calories: elf.calories}
			topThree[3] = tempSecond
			continue
		} else if elf.calories >= topThree[3].calories {
			topThree[3] = Elf{number: elf.number, calories: elf.calories}
			continue
		}
	}

	fmt.Println("[1] Elf-> ", topThree[1].number, "and with calories -> ", topThree[1].calories)
	fmt.Println("[1] Elf-> ", topThree[2].number, "and with calories -> ", topThree[2].calories)
	fmt.Println("[1] Elf-> ", topThree[3].number, "and with calories -> ", topThree[3].calories)
	fmt.Printf("Sum of their calories is Total -> %v", topThree[1].calories+topThree[2].calories+topThree[3].calories)
}
