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
	blocks := readInputs("input.txt")

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

// Read in file and split into passenger groups
func readInputs(filename string) [][]uint {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}
	contents := string(b)

	var blocks [][]uint
	for _, group := range strings.Split(contents, "\n\n") {
		var block []uint
		for _, line := range strings.Split(group, "\n") {
			block = append(block, toBinary(line))
		}
		blocks = append(blocks, block)
	}
	return blocks
}
