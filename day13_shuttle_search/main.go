package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Bus struct {
	id   int
	init int
}

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")
	time, buses := parseInputs(input)

	// Part 1 - calculate number of 1 and 3 sequential diffs, and return product.
	timeID := solvePart1(time, buses)
	fmt.Printf("Part 1 - ID of the earliest bus * wait time = %d\n", timeID)

	// Part 2 - find number of valid combinations of adaptors in input.
	t, ok := solvePart2(buses)
	if ok {
		fmt.Printf("Part 2 - first time for sequential bus times at %d minutes\n", t)
	}
}

func solvePart1(time int, buses []*Bus) int {
	// Calculate earliest arrival time of any bus after current time
	var firstArrivalID int
	firstArrivalTime := time * 1e9
	for _, bus := range buses {
		nextArrivalTime := ((time / bus.id) + 1) * bus.id
		if nextArrivalTime < firstArrivalTime {
			firstArrivalTime = nextArrivalTime
			firstArrivalID = bus.id
		}
	}

	waitMinutes := firstArrivalTime - time
	return waitMinutes * firstArrivalID
}

func solvePart2(buses []*Bus) (int, bool) {
	// Calculate distance between each solution (span)
	span := 1
	for _, bus := range buses {
		span *= bus.id
	}
	// Set initial looping vars based on bus 1 (with largest bus id)
	pointer := 0
	init := buses[pointer].init
	step := buses[pointer].id

	pointer++
	for i := init; i <= span; i += step {
		// Check next bus for time alignment, and then step by product of bus ids (LCM)
		adjust := buses[pointer].init
		busID := buses[pointer].id
		if (i-adjust)%busID == 0 {
			step *= buses[pointer].id
			pointer++
		}
		if pointer == len(buses) {
			return i, true
		}
	}
	return 0, false
}

func parseInputs(inputs []string) (int, []*Bus) {
	time, err := strconv.Atoi(inputs[0])
	if err != nil {
		log.Fatalf("could not parse int on line 1")
	}
	ids := inputs[1]
	var buses []*Bus
	for i, id := range strings.Split(ids, ",") {
		if id == "x" {
			continue
		}
		val, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalf("could not parse int on line 2")
		}
		buses = append(buses, &Bus{id: val, init: val - i})
	}

	// Sort data descending (cheap, but improves part 2 algorithm performance)
	sort.Slice(buses, func(i, j int) bool {
		if buses[j].id < buses[i].id {
			return true
		}
		return buses[j].id < buses[i].id
	})

	return time, buses
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
