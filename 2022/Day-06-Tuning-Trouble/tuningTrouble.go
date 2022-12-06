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

	startOfSignal(string(input))
	startOfMessage(string(input))

}

func startOfSignal(signal string) {
	for i := range signal {
		if unique(signal[i : i+4]) {
			fmt.Println("The signal starts a position: ", i+4)
			return
		}
	}
}

func startOfMessage(signal string) {
	for i := range signal {
		if unique(signal[i : i+14]) {
			fmt.Println("The message starts a position: ", i+14)
			return
		}
	}
}

func unique(input string) bool {
	var set map[rune]bool
	for _, char := range input {
		_, isInSet := set[char]
		if isInSet {
			return false
		}
		set[char] = true
	}
	return true
}
