package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//Read input file and add an empty line to the bottom, so the last elf doesnt get ignored
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	results := strings.Split(string(input), "\n")

	firstStrat(results)
	secondStrat(results)
}

func firstStrat(results []string) {

	//map all possible scores
	//1 for Rock (X A), 2 for Paper(Y B), and 3 for Scissors(Z C)
	// + 0 if you lost, 3 for a draw, and 6 if you won
	points := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6}
	score := 0

	for _, result := range results {
		score += points[strings.TrimSpace(result)]
	}

	fmt.Println(`Your final result for the first strat would be: `, score)
}

func secondStrat(results []string) {

	//map all possible scores
	//1 for Rock, 2 for Paper, and 3 for Scissors
	// + 0 if you lost, 3 for a draw, and 6 if you won
	points := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7}
	score := 0

	for _, result := range results {
		score += points[strings.TrimSpace(result)]
	}

	fmt.Println(`Your final result for the second strat would be: `, score)
}
