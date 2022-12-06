package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//split input at double new line (on windows and elsewhere)
	parts := strings.Split(strings.ReplaceAll(string(strings.TrimSuffix(string(input), "\n")), "\r\n", "\n"), "\n\n")

	crateMover9000(parseCrates(strings.Split(parts[0], "\n")), strings.Split(parts[1], "\n"))
	crateMover9001(parseCrates(strings.Split(parts[0], "\n")), strings.Split(parts[1], "\n"))

}

func crateMover9000(crates [][]rune, instructions []string) {
	for _, move := range instructions {
		//split lines into usable ints
		var amount, start, destination int
		var onCrane rune
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &start, &destination)
		//for each move push and pop crates
		for i := 0; i < amount; i++ {
			onCrane, crates[start-1] = crates[start-1][len(crates[start-1])-1], crates[start-1][:len(crates[start-1])-1] //pop
			crates[destination-1] = append(crates[destination-1], onCrane)                                               // push
		}
	}
	fmt.Print("The top crates on the Candy Crane 9000 will be: ")
	for _, stacks := range crates {
		fmt.Print(string(stacks[len(stacks)-1]))
	}
	fmt.Println()
}

func crateMover9001(crates [][]rune, instructions []string) {

	for _, move := range instructions {
		//split lines into usable ints
		var amount, start, destination int
		var onCrane []rune
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &start, &destination)
		//decrease the numbers to compSci numbers
		start--
		destination--

		onCrane, crates[start] = crates[start][len(crates[start])-amount:], crates[start][:len(crates[start])-amount] //split amount of the start slice
		crates[destination] = append(crates[destination], onCrane...)                                                 // push those onto the destination slice

	}
	fmt.Print("The top crates on the Candy Crane 9001 will be: ")
	for _, stacks := range crates {
		fmt.Print(string(stacks[len(stacks)-1]))
	}
}

func parseCrates(input []string) [][]rune {
	//make a 2 dimensional array, with the width of 1/4 of a line, aka. the number of stacks
	crates := make([][]rune, (len(input[0])+1)/4)

	//for each line of crates, beginning at the bottom
	for c := (len(input) - 2); c >= 0; c-- {
		//look at every 4th symbol (to ignore brackets and spaces)
		for i := 1; i <= len(input[c]); i += 4 {
			//if there is a crate
			if input[c][i] != ' ' {
				//append the letter to the appropriate row
				crates[((i+3)/4)-1] = append(crates[((i+3)/4)-1], rune(input[c][i]))
			}
		}
	}
	return crates
}
