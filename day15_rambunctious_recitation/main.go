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
	history := make([]*Log, rounds, rounds)

	// Start game with submitted numbers
	for i, num := range initial {
		history[num] = &Log{curr: i, prev: -1}
	}
	// Play out rest of game
	var last int
	for x := len(initial); x < rounds; x++ {
		// If last number not spoken before, then add and speak 0
		if history[last] == nil {
			last = 0
			history[last] = &Log{curr: x, prev: -1}
			continue
		}
		if history[last].prev == -1 {
			// If last number has only been spoken once, update, and speak 0
			speak := 0
			last = speak

			if history[speak] == nil {
				history[speak] = &Log{curr: x, prev: -1}
			} else {
				history[speak].prev = history[speak].curr
				history[speak].curr = x
			}
		} else {
			// If last number has been spoken more than once, speak diff
			speak := history[last].curr - history[last].prev
			last = speak

			if history[speak] == nil {
				history[speak] = &Log{curr: x, prev: -1}
			} else {
				history[speak].prev = history[speak].curr
				history[speak].curr = x
			}
		}
	}
	return last
}
