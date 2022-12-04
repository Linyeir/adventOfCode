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
	assignments := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	counterInclude := 0
	counterOverlap := 0

	//itarte through lines
	for _, elves := range assignments {
		//split lines into usable ints
		var firstBegin, firstEnd, secondBegin, secondEnd int
		fmt.Sscanf(elves, "%d-%d,%d-%d", &firstBegin, &firstEnd, &secondBegin, &secondEnd)
		//check boundaries and raise counters
		if ((firstBegin >= secondBegin) && (firstEnd <= secondEnd)) ||
			((firstBegin <= secondBegin) && (firstEnd >= secondEnd)) {
			counterInclude++
		}
		if (firstEnd >= secondBegin) && (secondEnd >= firstBegin) {
			counterOverlap++
		}
	}

	fmt.Println("The Assignments include each other in:", counterInclude)
	fmt.Println("The Assignments overlap  in:", counterOverlap)
}
