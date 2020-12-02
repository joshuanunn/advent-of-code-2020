package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read inputs into int slice
	inputs := readInputs("input.txt")

	// Part 1 - find two elements that sum to 2020 and calculate product
	elements, ok := findIn(inputs, 0, 2020, 2)
	if ok {
		summarise(elements)
	}

	// Part 2 -find three elements that sum to 2020 and calculate product
	elements, ok = findIn(inputs, 0, 2020, 3)
	if ok {
		summarise(elements)
	}
}

func summarise(elements []int) {
	product := 1
	for _, val := range elements {
		product *= val
	}
	log.Printf("Found elements: %v. Product = %d.\n", elements, product)
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

// Parse input file into slice of ints
func readInputs(filename string) []int {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var inputs []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("error parsing line in inputs as int")
		}
		inputs = append(inputs, number)
	}
	return inputs
}
