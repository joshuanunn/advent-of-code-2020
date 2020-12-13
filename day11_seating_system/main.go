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

type Seating struct {
	initial [][]byte
	final   [][]byte
	rows    int
	cols    int
}

type rule func(*Seating, int, int) int

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")

	// Part 1 - calculate number of 1 and 3 sequential diffs, and return product.
	seating := parseInputs(input)
	count, rounds := converge(seating, rulePart1, 4, 1000)
	fmt.Printf("Part 1 - total occupied seats: %d (after %d rounds)\n", count, rounds)

	// Part 2 - find number of valid combinations of adaptors in input.
	seating = parseInputs(input)
	count, rounds = converge(seating, rulePart2, 5, 1000)
	fmt.Printf("Part 2 - total occupied seats: %d (after %d rounds)\n", count, rounds)
}

func converge(seating *Seating, f rule, tolerance, rounds int) (int, int) {
	for round := 1; round <= rounds; round++ {
		converged := seating.sweep(f, tolerance)
		if converged {
			count := seating.countOccupiedGrid()
			return count, round
		}
	}
	return 0, rounds
}

// Function to calculate seating arrangements based on rule
func (s *Seating) sweep(f rule, tolerance int) bool {
	for y := 0; y < s.rows; y++ {
		for x := 0; x < s.cols; x++ {
			// Check current seat using appropriate ruleset
			countOccupied := f(s, x, y)
			if s.initial[y][x] == empty && countOccupied == 0 {
				s.final[y][x] = occupied
			} else if s.initial[y][x] == occupied && countOccupied >= tolerance {
				s.final[y][x] = empty
			}
		}
	}
	// Convergence check - has grid changed?
	status := s.checkGrid()
	// update seating state
	duplicateGrid(s.initial, s.final)
	return status
}

func (s *Seating) show() {
	for y := 0; y < len(s.initial); y++ {
		for x := 0; x < len(s.initial[0]); x++ {
			fmt.Printf("%c", s.initial[y][x])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (s *Seating) checkGrid() bool {
	var same int
	for y := 0; y < s.rows; y++ {
		for x := 0; x < s.cols; x++ {
			if s.initial[y][x] == s.final[y][x] {
				same++
			}
		}
	}
	if same == s.rows*s.cols {
		return true
	}
	return false
}

func (s *Seating) countOccupiedGrid() int {
	var total int
	for y, row := range s.initial {
		for x := range row {
			if s.initial[y][x] == occupied {
				total++
			}
		}
	}
	return total
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

func rulePart1(s *Seating, x, y int) int {
	var countOccupied int
	i0, i1 := -1, 1
	j0, j1 := -1, 1
	if x == 0 {
		i0 = 0
	}
	if x == s.cols-1 {
		i1 = 0
	}
	if y == 0 {
		j0 = 0
	}
	if y == s.rows-1 {
		j1 = 0
	}
	for j := j0; j <= j1; j++ {
		for i := i0; i <= i1; i++ {
			if i == j && j == 0 {
				continue
			}
			if s.initial[y+j][x+i] == occupied {
				countOccupied++
			}
		}
	}
	return countOccupied
}

func rulePart2(s *Seating, x, y int) int {
	var countOccupied int
	// span all 8 compass directions
	for dj := -1; dj <= 1; dj++ {
		for di := -1; di <= 1; di++ {
			if di == dj && dj == 0 {
				continue
			}
			xx, yy := x, y
			for {
				xx, yy = xx+di, yy+dj
				if xx < 0 || xx >= s.cols {
					break
				}
				if yy < 0 || yy >= s.rows {
					break
				}
				if s.initial[yy][xx] == empty {
					break
				}
				if s.initial[yy][xx] == occupied {
					countOccupied++
					break
				}
			}
		}
	}
	return countOccupied
}

func parseInputs(inputs []string) *Seating {
	var data [][]byte
	for _, line := range inputs {
		if len(line) > 0 {
			data = append(data, []byte(line))
		}
	}
	seating := &Seating{}
	seating.rows = len(data)
	seating.cols = len(data[0])
	seating.initial = data
	seating.final = generateGrid(seating.rows, seating.cols)
	duplicateGrid(seating.final, seating.initial)
	return seating
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
