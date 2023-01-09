package main

func main() {
	var elfs []Elf = ReadCalories()
	PrintElfWithHighestCalories(elfs)
	PrintTopThreeElfs(elfs)
}
