package main

import (
	"fmt"
	"os"
	"strings"
)

type Valve struct {
	rate         int
	destinations map[string]int
}

type Network map[string]Valve

type Situation struct {
	node  string
	time  int
	state int
}

// caching recursion results and trading space for time
var cache map[Situation]int

func main() {
	//Read input file
	input, err := os.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}
	start, n := parseInput(strings.TrimSuffix(string(input), "\n"))

	// doing this makes the calculation more expensive, cause one has to factor in weight of edges,
	// but cuts massivley on the numbver of calculations
	n = collapse(n, start)
	n.findMostRelieve(Situation{start, 30, 0})
}

func (n Network) findMostRelieve(s Situation) (maxRelieve int) {
	val, valOK := cache[s]
	if valOK {
		return val
	}

	return
}

func parseInput(input string) (string, Network) {
	n := make(Network)
	for _, line := range strings.Split(input, "\n") {
		v := Valve{}
		parts := strings.Fields(line)
		fmt.Sscanf(parts[4], "rate=%d;", &v.rate)
		v.destinations = make(map[string]int)
		for _, d := range parts[9:] {
			v.destinations[d[:2]] = 1
		}
		n[parts[1]] = v
	}
	return input[6:8], n
}

// remove any vertex with value 0
func collapse(nIn Network, start string) (nOut Network) {
	nOut = make(Network)
	for k, v := range nIn {
		nOut[k] = v
	}
	for k, v := range nOut {
		if k != start && v.rate == 0 {
			nOut = remove(nOut, k)
		}
	}
	return nOut
}

func remove(nIn Network, toRemove string) (nOut Network) {
	nOut = make(Network)
	for name, value := range nIn {
		if name != toRemove {
			_, isToRemoveInDestination := value.destinations[toRemove]
			if isToRemoveInDestination {
				for d, l := range nIn[toRemove].destinations {
					if d != name {
						distance, isDistance := value.destinations[d]
						if isDistance {
							value.destinations[d] = min(l+1, distance)
						} else {
							value.destinations[d] = l + 1
						}
					}
				}
				delete(value.destinations, toRemove)
			}
			nOut[name] = value
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
