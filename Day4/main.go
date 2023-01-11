package main

import (
	"filereader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := filereader.ReadFile("input.txt")
	var fullyOverlapingCounter int = 0
	var partlyOverlapingCounter int = 0

	for _, line := range lines {
		fullyOverlapingCounter += GetFullyOverlapping(line)
		partlyOverlapingCounter += GetPartlyOverlapping(line)
	}

	fmt.Println(fullyOverlapingCounter)
	fmt.Println(partlyOverlapingCounter)
}

func GetFullyOverlapping(line string) int {
	ranges := strings.Split(line, ",")
	r1 := strings.Split(ranges[0], "-")
	r2 := strings.Split(ranges[1], "-")

	x, _ := strconv.Atoi(r1[0])
	y, _ := strconv.Atoi(r2[0])
	z, _ := strconv.Atoi(r1[1])
	k, _ := strconv.Atoi(r2[1])

	result := (x-y)*(z-k) <= 0

	if result {
		return 1
	} else {
		return 0
	}
}

func GetPartlyOverlapping(line string) int {
	ranges := strings.Split(line, ",")
	r1 := strings.Split(ranges[0], "-")
	r2 := strings.Split(ranges[1], "-")

	a, _ := strconv.Atoi(r1[0])
	b, _ := strconv.Atoi(r1[1])
	x, _ := strconv.Atoi(r2[0])
	y, _ := strconv.Atoi(r2[1])

	result := (b-x)*(y-a) >= 0

	if result {
		return 1
	} else {
		return 0
	}
}
