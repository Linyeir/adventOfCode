package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputTmp := strings.TrimSuffix(string(input), "\n")
	rucksacks := strings.Split(string(inputTmp), "\n")

	partOne(rucksacks)
	partTwo(rucksacks)
}

func partOne(rucksacks []string) {
	prioritySum := 0

	for _, items := range rucksacks {
		front, back := items[:len(items)/2], items[len(items)/2:]
		for _, char := range front {
			if strings.ContainsRune(back, char) {
				prioritySum += priority(char)
			}
		}

	}

	fmt.Println("The sum of all priorities found in front and back pockets are: ", prioritySum)
}

func partTwo(rucksacks []string) {
	prioritySum := 0

	for i := 0; i < len(rucksacks); i += 3 {
		for _, char := range rucksacks[i] {
			if strings.ContainsRune(rucksacks[i+1], char) && strings.ContainsRune(rucksacks[i+2], char) {
				prioritySum += priority(char)
				break
			}
		}

	}

	fmt.Println("The sum of all priorities for each triple are: ", prioritySum)
}

func priority(item rune) int {
	if unicode.IsLower(item) {
		return int(item - 'a' + 1)
	} else {
		return int(item - 'A' + 27)
	}
}
