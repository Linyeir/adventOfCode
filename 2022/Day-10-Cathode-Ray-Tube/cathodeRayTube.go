package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Clock struct {
	tick, register, signalSum int
}

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	c := Clock{0, 1, 0}

	for _, line := range strings.Split(strings.TrimSuffix(string(input), "\n"), "\n") {
		c.draw()
		c.cycle()
		if line == "noop" {
			continue
		}
		c.draw()
		parts := strings.Fields(line)
		strength, _ := strconv.Atoi(parts[1])
		c.register = c.register + strength
		c.cycle()
	}

	fmt.Println("The sum of the six signal strengths is:", c.signalSum)

}

func (c *Clock) draw() {
	position := c.tick - (((c.tick) / 40) * 40)
	if position >= c.register-1 && position <= c.register+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if (c.tick+1)%40 == 0 {
		fmt.Println()
	}
}

func (c *Clock) cycle() {
	c.tick++
	if (c.tick-20)%40 == 0 {
		c.signalSum = c.signalSum + c.register*c.tick
	}
}
