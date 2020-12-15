package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var maskPattern = regexp.MustCompile(`^mask = ([01X]{36})$`)
var memPattern = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

type Program map[uint64]uint64

type Parser func(uint64, uint64, string, map[uint64]uint64)

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")

	// Part 1 - calculate number of 1 and 3 sequential diffs, and return product.
	program1 := parseInputs(input, parserPart1)
	fmt.Printf("Part 1 - Sum of allocated memory values = %d\n", program1.sum())

	// Part 2 - find number of valid combinations of adaptors in input.
	program2 := parseInputs(input, parserPart2)
	fmt.Printf("Part 2 - Sum of allocated memory values = %d\n", program2.sum())
}

func (p Program) sum() uint64 {
	var total uint64
	for _, value := range p {
		total += value
	}
	return total
}

func parserPart2(address, value uint64, line string, program map[uint64]uint64) {
	floatingBits := strings.Count(line, "X")
	addressCount := 1 << floatingBits
	bits := len(line) - 1

	for i := 0; i < addressCount; i++ {
		var position uint64
		mask := address
		for j := 0; j <= bits; j++ {
			if line[bits-j] == '1' {
				mask |= 1 << j
			}
			if line[bits-j] == 'X' {
				X := (uint64(i) >> position) & 1
				if X == 1 {
					mask |= (1 << j)
				} else {
					mask &= ^(1 << j)
				}
				position++
			}
		}
		program[mask] = value
	}
}

func parserPart1(address, value uint64, line string, program map[uint64]uint64) {
	var mask0, mask1 uint64
	bits := len(line) - 1
	for i := 0; i <= bits; i++ {
		if line[bits-i] == '0' {
			mask0 |= 1 << i
		}
		if line[bits-i] == '1' {
			mask1 |= 1 << i
		}
	}
	program[address] = (value & ^mask0) | mask1
}

func parseInt(line string) uint64 {
	val, err := strconv.ParseUint(line, 10, 64)
	if err != nil {
		log.Fatalf("could not parse mem int")
	}
	return val
}

func parseInputs(inputs []string, parser Parser) Program {
	program := make(Program)

	var mask string
	for _, line := range inputs {
		maskMatch := maskPattern.FindStringSubmatch(line)
		memMatch := memPattern.FindStringSubmatch(line)
		if len(maskMatch) == 2 {
			mask = maskMatch[1]
			continue
		}
		if len(memMatch) == 3 {
			address := parseInt(memMatch[1])
			value := parseInt(memMatch[2])
			// Calculate set of (address, value) and add to program
			parser(address, value, mask, program)
		}
	}
	return program
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
