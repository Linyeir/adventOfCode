package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	left  any
	right any
}

func main() {
	//Read input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	iSum, list := 0, []any{}

	for i, pair := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		pair := strings.Split(pair, "\n")
		var first, second any

		json.Unmarshal([]byte(pair[0]), &first)
		json.Unmarshal([]byte(pair[1]), &second)
		list = append(list, first, second)

		if compare([]Pair{{first, second}}) {
			iSum += i + 1
		}

	}
	fmt.Println("The sum of indices of packages in the right order is:", iSum)

	list = append(list, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(list, func(i, j int) bool { return compare([]Pair{{list[i], list[j]}}) })

	product := 1
	for i, p := range list {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			product *= i + 1
		}
	}
	fmt.Println("The product of the indices of the divider packets in the ordered list is:", product)

}

// true: right order, false: wrong order
func compare(stack []Pair) bool {
	//make stacks of things to be evaluated
	var left, right any

	//as long as there is a Stack
	for len(stack) > 0 {
		//pop the last item from stack
		left, right, stack = stack[len(stack)-1].left, stack[len(stack)-1].right, stack[:len(stack)-1]

		//check if these are arrays
		leftSlice, lOK := left.([]any)
		rightSlice, rOK := right.([]any)

		switch {
		case !lOK && !rOK:
			if int(left.(float64)) < int(right.(float64)) {
				return true
			} else if int(left.(float64)) > int(right.(float64)) {
				return false
			} else {
				continue
			}
		case !lOK:
			leftSlice = []any{left}
		case !rOK:
			rightSlice = []any{right}
		case len(leftSlice) == 0 && len(rightSlice) == 0:
			continue
		case len(leftSlice) == 0:
			return true
		case len(rightSlice) == 0:
			return false
		}

		//if slices are of unequal length, get them to the same length& add remainder to the stack
		if len(leftSlice) < len(rightSlice) {
			stack, rightSlice = append(stack, Pair{[]any{}, rightSlice[len(leftSlice):]}), rightSlice[:len(leftSlice)]
		} else if len(leftSlice) > len(rightSlice) {
			stack, leftSlice = append(stack, Pair{leftSlice[len(rightSlice):], []any{}}), leftSlice[:len(rightSlice)]
		}

		//add items of the slices to the stack
		for i := len(leftSlice) - 1; i >= 0; i-- {
			stack = append(stack, Pair{leftSlice[i], rightSlice[i]})
		}
	}

	return true
}
