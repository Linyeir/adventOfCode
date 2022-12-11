package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	//worry level of all items
	items []int
	//operation, term
	operation []string
	//divisor, truth-monkey, false-monkey
	test []int
	//total inspected items
	business int
}

type Troop []Monkey

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	troop := parseInput(string(input))

	// calculate least common multiple so the itemvalues dont escalate in part 2
	lcm := 1
	for _, m := range troop {
		lcm = lcm * m.test[0]
	}

	//pass 20 rounds and a division of my stress level by 3
	checkedShenanigans := troop.keepAway(20, func(w int) int { return w / 3 })
	fmt.Println("The two busiest monkeys have a shenanigan level of: ", checkedShenanigans)
	//pass 10000 rounds and a division of the item by the lcm
	uncheckedShenanigans := troop.keepAway(10000, func(w int) int { return w % lcm })
	fmt.Println("The two busiest unchecked monkeys have a shenanigan level of: ", uncheckedShenanigans)
}

func parseInput(input string) Troop {
	t := Troop{}
	for _, p := range strings.Split(input, "\n\n") {
		m := Monkey{}
		l := strings.Split(p, "\n")
		for _, item := range strings.Split(l[1][18:], ", ") {
			intItem, _ := strconv.Atoi(item)
			m.items = append(m.items, intItem)
		}
		f := strings.Fields(l[2])
		m.operation = f[4:]
		f = strings.Fields(l[3])
		divisor, _ := strconv.Atoi(f[3])
		f = strings.Fields(l[4])
		truthMonkey, _ := strconv.Atoi(f[5])
		f = strings.Fields(l[5])
		falseMonkey, _ := strconv.Atoi(f[5])
		m.test = []int{divisor, truthMonkey, falseMonkey}
		t = append(t, m)
	}
	return t
}

// multiply the inspected items of the two busiest monkeys
func (t Troop) shenaniganLevel() int {
	b := [2]int{0, 0}
	for _, m := range t {
		//if monkey is leaderboard worthy replace the 2nd best
		if m.business > b[1] {
			b[1] = m.business
		}
		//rearrange leaderboard
		if b[1] > b[0] {
			b[1], b[0] = b[0], b[1]
		}
	}
	return b[1] * b[0]
}

// simulates a troop of monkeys
// accept custom functions, to be applicable to both exercises
func (t Troop) keepAway(tally int, op func(int) int) int {
	//make copy of troop so the original isnt manipulated for the second iteration
	t = append(Troop{}, t...)
	for i := 0; i < tally; i++ {
		// round: all monkeys take turn ;
		for i, m := range t {
			// turn: inspect and throw every item in order
			for _, item := range m.items {
				//inspect and apply custom function
				newVal := op(m.evaluate(item))
				//test and add item to the corresponding monkey
				if newVal%m.test[0] == 0 {
					t[m.test[1]].items = append(t[m.test[1]].items, newVal)
				} else {
					t[m.test[2]].items = append(t[m.test[2]].items, newVal)
				}
			}
			//increase the counters and empty the item array
			t[i].business = m.business + len(m.items)
			t[i].items = []int{}
		}
	}

	return t.shenaniganLevel()
}

func (m Monkey) evaluate(item int) int {
	if m.operation[1] == "old" {
		return item * item
	}
	term, _ := strconv.Atoi(m.operation[1])
	if m.operation[0] == "*" {
		return item * term
	}
	if m.operation[0] == "+" {
		return item + term
	}
	return item
}
