package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"sync"
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
	visibilitySum, maxScore := 0, 0

	cPoint, cVisibility, cScore := make(chan image.Point, len(f)), make(chan bool, len(f)), make(chan int, len(f))

	var wg sync.WaitGroup

	//start workers
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go f.wEvaluateTree(i, cPoint, cVisibility, cScore)
	}

	//send work
	for point := range f {
		cPoint <- point
	}

	close(cPoint)
	for i := 0; i < len(f); i++ {
		isVisible, score := <-cVisibility, <-cScore
		if isVisible {
			visibilitySum++
		}
		if score > maxScore {
			maxScore = score
		}
	}
	close(cVisibility)
	close(cScore)

	fmt.Println("The number of visible trees is: ", visibilitySum)
	fmt.Println("The highest possible scenic score is: ", maxScore)
}

func (f Forest) wEvaluateTree(id int, cPoint <-chan image.Point, cVisibility chan<- bool, cScore chan<- int) {
	var wg sync.WaitGroup
	for point := range cPoint {
		wg.Add(1)
		go func(point image.Point) {
			tmpVis, tmpScore := f.evaluateTree(point)
			cVisibility <- tmpVis
			cScore <- tmpScore
			wg.Done()
		}(point)
	}
	wg.Wait()
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
