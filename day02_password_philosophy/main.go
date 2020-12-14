package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]+)`)

type Password struct {
	low       int
	high      int
	char      byte
	plaintext string
}

func main() {
	input := readInputs("input.txt")
	data := parseInputs(input)

	// Part 1 - count number of valid passwords using ruleset 1
	count := validatePart1(data)
	fmt.Printf("Part 1 - found a total of %d/%d valid passwords.\n", count, len(data))

	// Part 2 - count number of valid passwords using ruleset 2
	count = validatePart2(data)
	fmt.Printf("Part 2 - found a total of %d/%d valid passwords.\n", count, len(data))
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

func parseInputs(inputs []string) []*Password {
	var data []*Password
	for _, line := range inputs {
		if len(line) > 0 {
			elements := pattern.FindAllStringSubmatch(line, -1)[0]
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
				data = append(data, pwd)
			} else {
				log.Fatalf("failed to parse regexp for line")
			}
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
