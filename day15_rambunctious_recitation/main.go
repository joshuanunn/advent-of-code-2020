package main

import (
	"fmt"
)

type Log struct {
	curr int
	prev int
}

func main() {
	input := []int32{1, 20, 8, 12, 0, 14}

	// Part 1 - calculate number at end of round 2020.
	number1 := solve(input, 2020)
	fmt.Printf("Part 1 - number at end of round 2020 = %d\n", number1)

	// Part 2 - calculate number at end of round 30000000.
	number2 := solve(input, 30000000)
	fmt.Printf("Part 2 - number at end of round 30000000 =  %d\n", number2)
}

func solve(initial []int32, rounds int32) int32 {
	// map of number to last two rounds number spoken
	curr := make([]int32, rounds, rounds)
	prev := make([]int32, rounds, rounds)

	// Start game with submitted numbers
	for i, num := range initial {
		curr[num] = int32(i) + 1
		prev[num] = -1
	}
	// Play out rest of game
	var last int32
	start := int32(len(initial))
	for x := start; x < rounds; x++ {
		// If last not spoken before or spoken once, speak 0, else speak diff
		if curr[last] == 0 || prev[last] == -1 {
			last = 0
		} else {
			last = curr[last] - prev[last]
		}
		if curr[last] == 0 {
			prev[last] = -1
			curr[last] = x + 1
		} else {
			prev[last] = curr[last]
			curr[last] = x + 1
		}
	}
	return last
}
