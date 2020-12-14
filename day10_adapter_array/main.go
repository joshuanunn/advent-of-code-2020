package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readInputs("input.txt")
	data := parseInputs(input)

	// Part 1 - calculate number of 1 and 3 sequential diffs, and return product.
	diffs := prodDiffs(data)
	fmt.Printf("Part 1 - number of differences of ones * threes = %d\n", diffs)

	// Part 2 - find number of valid combinations of adaptors in input.
	count := sumBranches(data)
	fmt.Printf("Part 2 - number of valid combinations found: %d.\n", count)
}

func sumBranches(adaptors []int) int {
	var count int
	var start = map[int]int{0: 1}

	// Sort inputs and get max
	sort.Ints(adaptors)
	target := adaptors[len(adaptors)-1]

	for i := range adaptors {
		level := make(map[int]int)
		for initial, branches := range start {
			for _, jolt := range adaptors[i:] {
				diff := jolt - initial
				if diff < 1 {
					continue
				}
				if diff > 3 {
					break
				}
				if jolt == target {
					count += branches
				}
				level[jolt] += branches
			}
		}
		start = level
	}
	return count
}

func prodDiffs(adaptors []int) int {
	// Tweak copy of input slice to describe problem
	data := make([]int, len(adaptors))
	copy(data, adaptors)

	// Append start point (0), sort, and append end point (max+3)
	data = append(data, 0)
	sort.Ints(data)
	max := data[len(data)-1]
	data = append(data, max+3)

	// Calculate and count number of 1's and 3's
	var ones, threes int
	for i := 0; i < len(data)-1; i++ {
		diff := data[i+1] - data[i]
		if diff == 1 {
			ones++
		}
		if diff == 3 {
			threes++
		}
	}
	return ones * threes
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
