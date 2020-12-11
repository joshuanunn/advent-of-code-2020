package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Read inputs into slice of []byte
	input := readInputs("input.txt")
	grid := parseInputs(input)

	// Part 1 - count number of trees for run of (3, 1)
	trees := slopeRun(grid, 3, 1)
	fmt.Printf("Part 1 - found a total of %d trees.\n", trees)

	// Part 2 - product of number of trees for 5 different runs
	trees1 := slopeRun(grid, 1, 1)
	trees2 := slopeRun(grid, 3, 1)
	trees3 := slopeRun(grid, 5, 1)
	trees4 := slopeRun(grid, 7, 1)
	trees5 := slopeRun(grid, 1, 2)

	prod := trees1 * trees2 * trees3 * trees4 * trees5
	fmt.Printf("Part 2 - found a total product of %d based on total of 5 runs.\n", prod)
}

func slopeRun(grid [][]byte, right, down int) int {
	var trees, posX, posY, tree int
	for posY < len(grid) {
		posX, posY, tree = moveRightDown(grid, posX, posY, right, down)
		trees += tree
	}
	return trees
}

func moveRightDown(grid [][]byte, posX, posY, right, down int) (int, int, int) {
	// Grid dimensions (fixed height, repeating width)
	width := len(grid[0])
	height := len(grid)

	// Move and check for trees (#) at new location (if in-bounds)
	posX += right
	posY += down
	if posY < height {
		if grid[posY][posX%width] == '#' {
			return posX, posY, 1
		}
	}
	return posX, posY, 0
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
