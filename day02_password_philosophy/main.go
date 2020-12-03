package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	low       int
	high      int
	char      byte
	plaintext string
}

func main() {
	// Read inputs into slice of *Password
	inputs := readInputs("input.txt")

	// Part 1 - count number of valid passwords using ruleset 1
	count := validatePart1(inputs)
	fmt.Printf("Part 1 - found a total of %d/%d valid passwords.\n", count, len(inputs))

	// Part 2 - count number of valid passwords using ruleset 2
	count = validatePart2(inputs)
	fmt.Printf("Part 2 - found a total of %d/%d valid passwords.\n", count, len(inputs))
}

func validatePart1(passwords []*Password) int {
	var countValid int
	for _, pwd := range passwords {
		count := strings.Count(pwd.plaintext, string(pwd.char))
		if pwd.low <= count && count <= pwd.high {
			countValid++
		}
	}
	return countValid
}

func validatePart2(passwords []*Password) int {
	var countValid int
	for _, pwd := range passwords {
		one := pwd.plaintext[pwd.low-1]
		two := pwd.plaintext[pwd.high-1]
		if (one == pwd.char) != (two == pwd.char) {
			countValid++
		}
	}
	return countValid
}

// Parse input file into slice of *Password
func readInputs(filename string) []*Password {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]+)`)

	var inputs []*Password
	for scanner.Scan() {
		elements := re.FindAllStringSubmatch(scanner.Text(), -1)[0]
		if len(elements) == 5 {
			low, err := strconv.Atoi(elements[1])
			if err != nil {
				log.Fatalf("failed to parse lower bound as int")
			}
			high, err := strconv.Atoi(elements[2])
			if err != nil {
				log.Fatalf("failed to parse higher bound as int")
			}
			char := elements[3][0]
			plaintext := elements[4]

			pwd := &Password{low, high, char, plaintext}

			inputs = append(inputs, pwd)
		} else {
			log.Fatalf("failed to parse regexp for line")
		}
	}
	return inputs
}
