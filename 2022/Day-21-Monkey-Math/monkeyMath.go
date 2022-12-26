package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	left, right any
	operation   string
}

type DeducedMonkey struct {
	name  string
	value int
}

type Listeners map[string]Monkey
type Callers map[string]int

type Troop struct {
	listeners Listeners
	callers   Callers
}

func main() {
	//Read input file
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println("The monkey 'root' should scream:", solvePart1(parseInput(string(input))))
	fmt.Println("The monkey 'humn' should scream:", solvePart2(parseInput(string(input))))
}

func solvePart1(t Troop) int {
	for {
		for caller := range t.callers {
			t.monkeyScream(caller)
			if value, valueOk := t.callers["root"]; valueOk {
				return value
			}
		}
	}
}

func solvePart2(t Troop) int {
	// remove me from the list
	delete(t.callers, "humn")
	//solve as far as possible
	for len(t.callers) > 0 {
		for caller := range t.callers {
			t.monkeyScream(caller)
		}
	}
	fmt.Println(t.listeners)
	//deduce the first number a monkey has to scream from root
	root := t.listeners["root"]
	lInt, lOk := root.left.(int)
	var dm DeducedMonkey
	if lOk {
		dm = DeducedMonkey{root.right.(string), lInt}
	} else {
		dm = DeducedMonkey{root.left.(string), root.right.(int)}
	}
	delete(t.listeners, "root")
	for {
		dm = t.deduceNextMonkey(dm)
		if dm.name == "humn" {
			return dm.value
		}
	}
}

func (t Troop) deduceNextMonkey(dIn DeducedMonkey) (dOut DeducedMonkey) {
	for lisName, lisMonkey := range t.listeners {
		if lisName == dIn.name {
			return deduce(dIn.value, lisMonkey)
		}
	}
	return DeducedMonkey{}
}

func deduce(result int, m Monkey) (dm DeducedMonkey) {
	lInt, lIsInt := m.left.(int)
	if lIsInt {
		dm.name = m.right.(string)
		switch m.operation {
		case "+":
			dm.value = result - lInt
		case "-":
			dm.value = lInt - result
		case "*":
			dm.value = result / lInt
		case "/":
			dm.value = lInt / result
		}
		return
	}
	dm.name = m.left.(string)
	rInt := m.right.(int)
	switch m.operation {
	case "+":
		dm.value = result - rInt
	case "-":
		dm.value = result + rInt
	case "*":
		dm.value = result / rInt
	case "/":
		dm.value = result * rInt
	}
	return
}

func (t Troop) monkeyScream(caller string) {
	for nameListener, monkey := range t.listeners {
		if monkey.left == caller {
			monkey.left = t.callers[caller]
		}
		if monkey.right == caller {
			monkey.right = t.callers[caller]
		}
		mValue, monkeyIsReady := monkey.evaluate()
		if monkeyIsReady {
			delete(t.listeners, nameListener)
			t.callers[nameListener] = mValue
		} else {
			t.listeners[nameListener] = monkey
		}
	}
	delete(t.callers, caller)
}

func (m Monkey) evaluate() (int, bool) {
	lInt, leftInt := m.left.(int)
	rInt, rightInt := m.right.(int)

	if leftInt && rightInt {
		switch m.operation {
		case "+":
			return lInt + rInt, true
		case "-":
			return lInt - rInt, true
		case "*":
			return lInt * rInt, true
		case "/":
			return lInt / rInt, true
		}
	}
	return 0, false
}

func parseInput(input string) Troop {
	t := Troop{make(Listeners), make(Callers)}
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if len(line) <= 10 {
			name := line[:4]
			value, _ := strconv.Atoi(strings.TrimSpace(line[6:]))
			t.callers[name] = value
		} else if len(line) == 17 {
			t.listeners[line[:4]] = Monkey{line[6:10], line[13:], string(line[11])}
		} else {
			panic("waaah")
		}
	}
	return t
}
