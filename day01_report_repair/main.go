package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputs := readInputs("input.txt")
	data := parseInputs(inputs)

	// Part 1 - find two elements that sum to 2020 and calculate product
	elements, ok := findIn(data, 0, 2020, 2)
	if ok {
		product := summarise(elements)
		fmt.Printf("Part 1 - product of elements = %d\n", product)
	}

	// Part 2 -find three elements that sum to 2020 and calculate product
	elements, ok = findIn(data, 0, 2020, 3)
	if ok {
		product := summarise(elements)
		fmt.Printf("Part 2 - product of elements = %d\n", product)
	}
}

func summarise(elements []int) int {
	product := 1
	for _, val := range elements {
		product *= val
	}
	return product
}

func findIn(inputs []int, idx, target, level int) ([]int, bool) {
	// Base case
	if level <= 1 {
		for i := idx; i < len(inputs); i++ {
			val := inputs[i]
			if val == target {
				return []int{val}, true
			}
		}
		return nil, false
	}
	for i := idx; i < len(inputs); i++ {
		val := inputs[i]
		vals, ok := findIn(inputs, i+1, target-val, level-1)
		if ok {
			vals = append(vals, val)
			return vals, true
		}
	}
	return nil, false
}

func parseInputs(inputs []string) []int {
	var data []int
	for _, line := range inputs {
		if len(line) > 0 {
			val, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("could not parse int on line")
			}
			data = append(data, val)
		}
	}
	return data
}

func readInputs(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}
	lines := string(b)

	var inputs []string
	for _, line := range strings.Split(lines, "\n") {
		inputs = append(inputs, line)
	}
	return inputs
}
