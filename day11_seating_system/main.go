package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	floor    = byte('.')
	empty    = byte('L')
	occupied = byte('#')
)

type rule func([][]byte, int, int, int, int) int

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")
	data := parseInputs(input)

	// Part 1 - calculate number of 1 and 3 sequential diffs, and return product.
	count, rounds := converge(data, rulePart1, 4, 1000)
	fmt.Printf("Part 1 - total occupied seats: %d (after %d rounds)\n", count, rounds)

	// Part 2 - find number of valid combinations of adaptors in input.
	count, rounds = converge(data, rulePart2, 5, 1000)
	fmt.Printf("Part 2 - total occupied seats: %d (after %d rounds)\n", count, rounds)
}

func converge(data [][]byte, f rule, tolerance, rounds int) (int, int) {
	for round := 1; round <= rounds; round++ {
		data2, converged := sweep(data, f, tolerance)
		if converged {
			count := countOccupiedGrid(data)
			return count, round
		}
		data = data2
	}
	return 0, rounds
}

func generateGrid(rows, cols int) [][]byte {
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]byte, cols)
	}
	return grid
}

func duplicateGrid(dst, src [][]byte) {
	for y, row := range src {
		for x := range row {
			dst[y][x] = src[y][x]
		}
	}
}

func checkGrid(gridA, gridB [][]byte) bool {
	var total, same int
	for y, row := range gridB {
		for x := range row {
			total++
			if gridA[y][x] == gridB[y][x] {
				same++
			}
		}
	}
	if total == same {
		return true
	}
	return false
}

func countOccupiedGrid(grid [][]byte) int {
	var total int
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == occupied {
				total++
			}
		}
	}
	return total
}

func rulePart1(initial [][]byte, x, y, cols, rows int) int {
	var countOccupied int
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i == j && j == 0 {
				continue
			}
			xx := x + i
			yy := y + j
			if xx < 0 || xx >= cols {
				continue
			}
			if yy < 0 || yy >= rows {
				continue
			}
			if initial[yy][xx] == occupied {
				countOccupied++
			}

		}
	}
	return countOccupied
}

func rulePart2(initial [][]byte, x, y, cols, rows int) int {
	var countOccupied int
	// span all 8 compass directions
	for dj := -1; dj <= 1; dj++ {
		for di := -1; di <= 1; di++ {
			if di == dj && dj == 0 {
				continue
			}
			xx := x
			yy := y
			for {
				xx += di
				yy += dj

				if xx < 0 || xx >= cols {
					break
				}
				if yy < 0 || yy >= rows {
					break
				}
				if initial[yy][xx] == empty {
					break
				}
				if initial[yy][xx] == occupied {
					countOccupied++
					break
				}
			}
		}
	}
	return countOccupied
}

// Function to calculate seating arrangements based on rule
func sweep(initial [][]byte, f rule, tolerance int) ([][]byte, bool) {
	rows := len(initial)
	cols := len(initial[0])

	final := generateGrid(rows, cols)
	duplicateGrid(final, initial)

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			// Check current seat using appropriate ruleset
			countOccupied := f(initial, x, y, cols, rows)
			if initial[y][x] == empty && countOccupied == 0 {
				final[y][x] = occupied
			} else if initial[y][x] == occupied && countOccupied >= tolerance {
				final[y][x] = empty
			}
		}
	}
	// Convergence check - has grid changed?
	status := checkGrid(initial, final)
	return final, status
}

func show(initial [][]byte) {
	for y := 0; y < len(initial); y++ {
		for x := 0; x < len(initial[0]); x++ {
			fmt.Printf("%c", initial[y][x])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func parseInputs(inputs []string) [][]byte {
	var data [][]byte
	for _, line := range inputs {
		if len(line) > 0 {
			data = append(data, []byte(line))
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
