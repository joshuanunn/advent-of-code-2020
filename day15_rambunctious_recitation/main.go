package main

import (
	"fmt"
)

type Log struct {
	curr int
	prev int
}

func main() {
	input := []int{1, 20, 8, 12, 0, 14}

	// Part 1 - calculate number at end of round 2020.
	number1 := solve(input, 2020)
	fmt.Printf("Part 1 - number at end of round 2020 = %d\n", number1)

	// Part 2 - calculate number at end of round 30000000.
	number2 := solve(input, 30000000)
	fmt.Printf("Part 2 - number at end of round 30000000 =  %d\n", number2)
}

func solve(initial []int, rounds int) int {
	// map of number to last two rounds number spoken
	curr := make([]int, rounds, rounds)
	prev := make([]int, rounds, rounds)

	// Start game with submitted numbers
	for i, num := range initial {
		curr[num] = i + 1
		prev[num] = -1
	}
	// Play out rest of game
	var last, speak int
	for x := len(initial); x < rounds; x++ {
		// If last number not spoken before, then add and speak 0
		if curr[last] == 0 {
			last = 0
			curr[0] = x + 1
			prev[0] = -1
			continue
		}
		// If last number only spoken once, speak 0, else speak diff
		if prev[last] == -1 {
			speak = 0
		} else {
			speak = curr[last] - prev[last]
		}
		last = speak
		if curr[speak] == 0 {
			curr[speak] = x + 1
			prev[speak] = -1
		} else {
			prev[speak] = curr[speak]
			curr[speak] = x + 1
		}
	}
	return last
}
