package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The signal starts a position: ", firstUniqueSequence(string(input), 4))
	fmt.Println("The message starts a position: ", firstUniqueSequence(string(input), 14))
}

func firstUniqueSequence(signal string, length int) int {
	for i := range signal {
		if unique(signal[i : i+length]) {
			return i + length
		}
	}
	return len(signal)
}

func unique(input string) bool {
	set := make(map[rune]bool)
	for _, char := range input {
		_, isInSet := set[char]
		if isInSet {
			return false
		}
		set[char] = true
	}
	return true
}
