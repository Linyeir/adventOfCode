package main

import (
	"errors"
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)

type heightMap map[image.Point]field
type field struct {
	height   rune
	distance int
}

const maxInt = 2147483647

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(string(input))
	part2(string(input))
}

func part1(input string) {
	//parse Input and look for shortest path from S to E
	mountains, start, end := parseInput(input)
	shortestPath, _ := mountains.shortestRoute(start, end)
	fmt.Println("The shortest Path from A is", shortestPath)
}

func part2(input string) {
	mountains, _, end := parseInput(input)
	shortestPath := maxInt
	//look for the shortest path for every starting field a
	for k, v := range mountains {
		cpMountains := copyMap(mountains)
		if v.height == 'a' {
			tmpShortestPath, err := cpMountains.shortestRoute(k, end)
			if err == nil && tmpShortestPath < shortestPath {
				shortestPath = tmpShortestPath
			}
		}
	}
	fmt.Println("The shortest Path from A is", shortestPath)
}

func (h heightMap) shortestRoute(start image.Point, end image.Point) (int, error) {
	//define possible moves
	directions := []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	h[start] = field{h[start].height, 0}
	//BFS: take a stepin every direction until the End point is reached
	for i := 0; h[end].distance == maxInt; i++ {
		exploredAnything := false
		for currentPoint, currentField := range h {
			if currentField.distance == i {
				exploredAnything = true
				for _, dir := range directions {
					nextPoint := currentPoint.Add(dir)
					nextField, exists := h[nextPoint]
					if !exists || nextField.distance <= i {
						continue
					}
					if nextField.height <= currentField.height+1 {
						nextField.distance = i + 1
					}
					h[nextPoint] = nextField
				}
			}
		}
		//if there wasnt any way to go, throw an error
		if !exploredAnything {
			return maxInt, errors.New("no possible solution")
		}
	}
	return h[end].distance, nil
}

func parseInput(input string) (m heightMap, start, end image.Point) {
	m = make(heightMap)
	for x, line := range strings.Split(input, "\n") {
		for y, letter := range line {

			m[image.Point{x, y}] = field{letter, maxInt}
			if letter == 'S' {
				m[image.Point{x, y}] = field{'a', maxInt}
				start = image.Point{x, y}
			}
			if letter == 'E' {
				m[image.Point{x, y}] = field{'z', maxInt}
				end = image.Point{x, y}
			}
		}
	}
	return
}

func copyMap(input heightMap) heightMap {
	cp := make(heightMap)
	for k, v := range input {
		cp[k] = v
	}
	return cp
}
