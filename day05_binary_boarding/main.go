package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read inputs into slice of string, then parse into slice of ints
	passes := readInputs("input.txt")
	sids := parseInputs(passes)
	sidMin, sidMax := intRange(sids)

	// Part 1 - calculate highest boarding pass seat ID
	fmt.Printf("Part 1 - found highest seat id of %d.\n", sidMax)

	// Part 2 - find missing boarding pass seat ID
	missing := findMissing(sids, sidMin, sidMax)
	fmt.Printf("Part 2 - found missing seats: %v.\n", missing)
}

func findMissing(ints []int, min, max int) []int {
	var missing []int
	// Loop over range of all expected ints
	for i := min; i <= max; i++ {
		match := false
		// Loop over actual slice of ints
		for _, sid := range ints {
			if i == sid {
				match = true
				break
			}
		}
		if !match {
			missing = append(missing, i)
		}
	}
	return missing
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

// Parse passport string blocks into slice of *Passport
func parseInputs(passes []string) []int {
	var sids []int
	for _, pass := range passes {
		pass = strings.ReplaceAll(pass, "F", "0")
		pass = strings.ReplaceAll(pass, "B", "1")
		pass = strings.ReplaceAll(pass, "L", "0")
		pass = strings.ReplaceAll(pass, "R", "1")

		// Convert boarding pass string to binary representation
		if i, err := strconv.ParseInt(pass, 2, 16); err != nil {
			log.Fatalf("failed to parse boarding pass number!")
		} else {
			row := int((i & 0b1111111000) >> 3)
			col := int(i & 0b0000000111)
			sid := row*8 + col

			sids = append(sids, sid)
		}
	}
	return sids
}

// Read in file and split into boarding pass strings
func readInputs(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passes []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 10 {
			log.Fatalf("line has incorrect number of characters!")
		}
		passes = append(passes, line)
	}
	return passes
}
