package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"time"
)

type Forest map[image.Point]int

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	forest := dataToTreeMap(strings.Split(string(input), "\n"))
	forest.treeHouseSearch()

	fmt.Println("Time taken:", time.Since(start))
}

func (f Forest) treeHouseSearch() {
	visibilitySum := 0
	maxScore := 0
	for point := range f {
		//if a tree is reached which is equal or higher, skip to next tree
		//if edge of map is reached add tree to the visibleSum
		isVisible, score := f.evaluateTree(point)
		if isVisible {
			visibilitySum++
		}
		if score > maxScore {
			maxScore = score
		}
	}

	fmt.Println("The number of visible trees is: ", visibilitySum)
	fmt.Println("The highest possible scenic score is: ", maxScore)
}

func (f Forest) evaluateTree(point image.Point) (bool, int) {
	visible, score := false, 1

	for _, direction := range []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		for factor := 1; ; factor++ {
			targetHeight, exists := f[point.Add(direction.Mul(factor))]
			if !exists {
				visible = true
				score = score * (factor - 1)
				break
			} else if targetHeight >= f[point] {
				score *= factor
				break
			}
		}
	}
	return visible, score
}

func dataToTreeMap(input []string) Forest {
	forest := make(Forest)
	for x, line := range input {
		for y, tree := range strings.TrimSpace(line) {
			forest[image.Point{x, y}] = int(tree - '0')
		}
	}
	return forest
}
