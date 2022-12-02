package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read input file and add an empty line to the bottom, so the last elf doesnt get ignored
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	snacks := strings.Split(string(input), "\n")
	snacks = append(snacks, "")

	//init 3 elves to count totals and
	strongestElves := [3]int{0, 0, 0}
	currentElfCalories := 0

	for _, calories := range snacks {

		snack, err := strconv.Atoi(strings.TrimSpace(calories))
		currentElfCalories += snack
		//If error is different from nil then I found an new elf
		//& I position the old one into the leaderboard
		if err != nil {
			if currentElfCalories > strongestElves[0] {
				strongestElves[0], strongestElves[1], strongestElves[2] = currentElfCalories, strongestElves[0], strongestElves[1]
			} else if currentElfCalories > strongestElves[1] {
				strongestElves[1], strongestElves[2] = currentElfCalories, strongestElves[1]
			} else if currentElfCalories > strongestElves[2] {
				strongestElves[2] = currentElfCalories
			}
			currentElfCalories = 0

		}
	}

	fmt.Println(`The strongest elf carries: `, strongestElves[0])

	fmt.Println(`The strongest three elves carries: `, strongestElves[0]+strongestElves[1]+strongestElves[2])
}
