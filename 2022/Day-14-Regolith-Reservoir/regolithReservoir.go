package main

import (
	"errors"
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

type Cave map[image.Point]rune

func main() {
	//Read input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	waterfallCave := parseInput(string(input))
	var i int
	root := image.Pt(500, 0)
	for i = 0; ; i++ {
		err := waterfallCave.addSand(root)
		if err != nil {
			break
		}
	}
	fmt.Println("The number of setteling kernels in the scan is: ", i)
	waterfallCave.addBottom(root)

	for i++; ; i++ {
		err := waterfallCave.addSand(root)
		if err != nil {
			break
		}
	}

	fmt.Println("The total number of setteling kernels is: ", i)

}

func (c Cave) addSand(kernel image.Point) error {
	directions := []image.Point{{0, 1}, {-1, 1}, {1, 1}}
	_, _, _, maxY := c.getDimensions()
	for kernel.Y <= maxY {
		hasFallen := false
		for _, dir := range directions {
			_, goalExists := c[kernel.Add(dir)]
			if !goalExists {
				kernel = kernel.Add(dir)
				hasFallen = true
				break
			}
		}
		if kernel == image.Pt(500, 0) {
			c[kernel] = 'o'
			return errors.New("at spawn")
		}

		if !hasFallen {
			c[kernel] = 'o'
			return nil
		}

	}
	return errors.New("out of bounds")
}

// add a bottom 2 under maxY
func (c Cave) addBottom(root image.Point) {
	_, _, _, maxY := c.getDimensions()

	for i := (root.X - maxY - 3); i < (root.X + maxY + 3); i++ {
		c[image.Pt(i, maxY+2)] = '#'
	}

}

// turn aoc input into s map slice
func parseInput(input string) Cave {
	c := make(Cave)
	for _, line := range strings.Split(input, "\n") {
		points := strings.Split(line, "->")
		for i := 0; i < len(points)-1; i++ {
			var start, end image.Point
			fmt.Sscanf(points[i], "%d,%d", &start.X, &start.Y)
			fmt.Sscanf(points[i+1], "%d,%d", &end.X, &end.Y)

			for direction := getDirection(start, end); start != end.Add(direction); start = start.Add(direction) {
				c[start] = '#'
			}
		}
	}

	return c
}

// assuming exactly one direction will be different
func getDirection(first, second image.Point) image.Point {
	if first.X == second.X {
		if first.Y > second.Y {
			return image.Point{0, -1}
		}
		//else second Y will be larger
		return image.Point{0, 1}
	}
	//else direction will be on the X axis
	if first.X > second.X {
		return image.Point{-1, 0}
	}
	//else second X will be larger
	return image.Point{1, 0}
}

func (c Cave) print() {
	minX, minY, maxX, maxY := c.getDimensions()

	time.Sleep(30 * time.Millisecond)
	fmt.Println()
	fmt.Println("Printing Slice")

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pixel, pOk := c[image.Pt(x, y)]
			if pOk {
				fmt.Print(string(pixel))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (c Cave) getDimensions() (minX, minY, maxX, maxY int) {
	minX, minY, maxX, maxY = 500, 0, 500, 0
	for point := range c {
		minX = min(minX, point.X)
		minY = min(minY, point.Y)
		maxX = max(maxX, point.X)
		maxY = max(maxY, point.Y)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
