package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readInputs("input.txt")
	cardPK, doorPK := parseInputs(input)

	// Part 1 - calculate encryption key.
	subjectNumber := 7
	cardLoopSize := extractLoopSize(subjectNumber, cardPK)
	doorLoopSize := extractLoopSize(subjectNumber, doorPK)

	encryptionKey := extractEncryptionKey(cardLoopSize, doorPK)
	// Double check same key is derived from second combination
	if encryptionKey != extractEncryptionKey(doorLoopSize, cardPK) {
		log.Fatalln("encryption keys do not match")
	}

	fmt.Printf("Part 1 - handshake encryption key = %d\n", encryptionKey)
}

func extractEncryptionKey(loopSize, PK int) int {
	value := 1
	for x := 0; x < loopSize; x++ {
		value *= PK
		value %= 20201227
	}
	return value
}

func extractLoopSize(subjectNumber, PK int) int {
	value := 1
	loopSize := 0
	for {
		loopSize++
		value *= subjectNumber
		value %= 20201227
		if value == PK {
			break
		}
	}
	return loopSize
}

func parseInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalln("error parsing int")
	}
	return val
}

func parseInputs(inputs []string) (int, int) {
	if len(inputs) != 2 {
		log.Fatalln("problem with input file")
	}
	cardPublicKey := parseInt(inputs[0])
	doorPublicKey := parseInt(inputs[1])
	return cardPublicKey, doorPublicKey
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
