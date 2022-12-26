package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type Pack map[image.Point]image.Point

func main() {
	//Read input file
	input, err := os.ReadFile("input.txt")
	lineToCheck := 2000000
	maxRange := 4000000

	if err != nil {
		panic(err)
	}
	pack := parseInput(string(input))

	//part 1
	noPossibleBeacons := pack.getLine(lineToCheck)
	fmt.Println("in row ", lineToCheck, " can be no beacons in ", noPossibleBeacons, "spaces.")

	// part 2
	point := pack.checkRange(maxRange)
	fmt.Println(point)
	fmt.Println("The value of the empty place is", point.X*4000000+point.Y)

}

func (p Pack) checkRange(max int) image.Point {
	//create a distance map
	d := make(map[image.Point]int)
	for k, v := range p {
		d[k] = getDistance(k, v)
	}

	//loop through all sensors
	//if any point aounrd their perimeter is not in any elses sensor, thats your boi
	for sL, dL := range d {
		dL++
		for i := 0; i <= dL; i++ {
			corners := []image.Point{
				sL.Add(image.Pt(-dL+i, i)),
				sL.Add(image.Pt(i, dL-i)),
				sL.Add(image.Pt(dL-i, -i)),
				sL.Add(image.Pt(-i, -dL+i))}
			for _, c := range corners {
				isCornerImpossible := false
				if c.X >= 0 && c.X <= max && c.Y >= 0 && c.Y <= max {

					for sF, dF := range d {
						if getDistance(c, sF) <= dF {
							isCornerImpossible = true
							break
						}
					}
					if !isCornerImpossible {
						return c
					}
				}
			}
		}
	}

	return image.Pt(-1, -1)
}

func (p Pack) getLine(line int) (ctr int) {
	minX, _, maxX, _ := p.getDimensions()
	for iX := minX; iX <= maxX; iX++ {
		noSignalAt := false
		for k, v := range p {
			if iX == v.X && line == v.Y {
				break
			}
			if getDistance(k, image.Pt(iX, line)) <= getDistance(k, v) {
				noSignalAt = true
			}
		}
		if noSignalAt {
			ctr++
		}
	}
	return
}

func (p Pack) noSignalAt(a image.Point) rune {
	symbol := '.'

	return symbol
}

func parseInput(input string) Pack {
	p := make(Pack)
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var scanner, beacon image.Point
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &scanner.X, &scanner.Y, &beacon.X, &beacon.Y)
		p[scanner] = beacon
	}
	return p
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getDistance(a, b image.Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func (p Pack) getDimensions() (minX, minY, maxX, maxY int) {
	for s, b := range p {
		minX = min(minX, s.X-getDistance(s, b))
		minY = min(minY, s.Y-getDistance(s, b))
		maxX = max(maxX, s.X+getDistance(s, b))
		maxY = max(maxY, s.Y+getDistance(s, b))
	}
	return
}
