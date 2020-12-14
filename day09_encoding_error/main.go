package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")
	data := parseInputs(input)
	preamble := 25

	// Sense check inputs
	if len(data) <= preamble {
		log.Fatalf("data slice is too short for supplied preamble")
	}

	// Part 1 - find first int in slice not a sum of the previous PREAMBLE number of ints.
	idx, ok := findInvalid(data, preamble)
	if ok {
		fmt.Printf("Part 1 - number %d at index %d is not sum of previous %d.\n", data[idx], idx, preamble)
	}

	// Part 2 - find "weakness" number in data (from contiguous subarray which sums to target).
	target := data[idx]
	weakness, ok := findContiguous(data, target)
	if ok {
		fmt.Printf("Part 2 - weakness number found: %d.\n", weakness)
	}
}

func findInvalid(data []int, preamble int) (int, bool) {
	for i := preamble; i < len(data); i++ {
		window := data[i-preamble : i]
		// Check number at index i expressable as sum of rolling window
		valid := false
		for j := 0; j < len(window); j++ {
			check := data[i] - window[j]
			for k := 0; k < len(window); k++ {
				if (check == window[k]) && (window[j] != window[k]) {
					valid = true
					break
				}
			}
		}
		if !valid {
			return i, true
		}
	}
	return 0, false
}

func findContiguous(data []int, target int) (int, bool) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 2; j < len(data)-1; j++ {
			check := sumRange(data[i:j])
			if check > target {
				break
			}
			if check == target {
				min, max := intRange(data[i:j])
				return min + max, true
			}
		}
	}
	return 0, false
}

func sumRange(data []int) int {
	var total int
	for _, val := range data {
		total += val
	}
	return total
}

func intRange(ints []int) (int, int) {
	var minVal int = 1e9
	var maxVal int
	for _, i := range ints {
		if i < minVal {
			minVal = i
		}
		if i > maxVal {
			maxVal = i
		}
	}
	return minVal, maxVal
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
