package filereader

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filePath string) []string {
	var lines []string

	input, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
