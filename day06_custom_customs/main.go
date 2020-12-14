package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/bits"
	"strings"
)

func main() {
	// Read inputs into slice of []uint
	input := readInputs("input.txt")
	blocks := parseInputs(input)

	// Part 1 - count number of customs declarations using ruleset 1
	count := countPart1(blocks)
	fmt.Printf("Part 1 - total declaration count of %d from all groups.\n", count)

	// Part 2 - count number of customs declarations using ruleset 2
	count = countPart2(blocks)
	fmt.Printf("Part 2 - total declaration count of %d from all groups.\n", count)
}

func countPart1(blocks [][]uint) int {
	var count int
	for _, block := range blocks {
		unique := block[0]
		for i := 1; i < len(block); i++ {
			unique |= block[i]
		}
		count += bits.OnesCount(unique)
	}
	return count
}

func countPart2(blocks [][]uint) int {
	var count int
	for _, block := range blocks {
		unique := block[0]
		for i := 1; i < len(block); i++ {
			unique &= block[i]
		}
		count += bits.OnesCount(unique)
	}
	return count
}

// Convert chars a-z to unique binary representation
func toBinary(chars string) uint {
	var binary uint
	for _, char := range []byte(chars) {
		binary |= 1 << (char - 'a')
	}
	return binary
}

func parseInputs(inputs []string) [][]uint {
	var blocks [][]uint
	var block []uint
	for i, line := range inputs {
		if line != "" {
			block = append(block, toBinary(line))
		}
		if line == "" || i == len(inputs)-1 {
			blocks = append(blocks, block)
			block = nil
		}
	}
	return blocks
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
