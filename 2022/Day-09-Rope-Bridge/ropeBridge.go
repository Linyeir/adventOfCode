package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)

type Rope []image.Point

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//allow translation of letter to coordinate based move
	directionToCoordsMap := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}

	elvenRope := make(Rope, 10)
	coordsVisitedBy1, coordsVisitedBy9 := map[image.Point]bool{}, map[image.Point]bool{}
	//for every line
	for _, motion := range strings.Split(string(input), "\n") {
		var direction rune
		var amount int
		//split into direction and amount
		fmt.Sscanf(motion, "%c %d", &direction, &amount)

		//for every move
		for i := 0; i < amount; i++ {
			//move head
			elvenRope.move(directionToCoordsMap[direction])
			coordsVisitedBy1[elvenRope[1]] = true
			coordsVisitedBy9[elvenRope[9]] = true
		}
	}
	fmt.Println("The Tail visited by a one long Rope", len(coordsVisitedBy1), "fields")
	fmt.Println("The Tail visited by a ten long Rope", len(coordsVisitedBy9), "fields")
}
func (r Rope) move(direction image.Point) {
	r[0] = r[0].Add(direction)
	//follow tail
	//get distance from head to tail
	for i := 1; i < len(r); i++ {
		r[i] = follow(r[i-1], r[i])
	}
}
func follow(lead image.Point, back image.Point) image.Point {
	distance := lead.Sub(back)
	//if distance greater 1
	//no double exeedance, cause the lead can only move in one direction
	if abs(distance.X) > 1 || abs(distance.Y) > 1 {
		//move tail closer
		return back.Add(image.Point{sgn(distance.X), sgn(distance.Y)})
	}
	return back
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
