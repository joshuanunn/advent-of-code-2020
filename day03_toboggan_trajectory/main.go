package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Read inputs into slice of []byte
	grid := readInputs("input.txt")

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

// Parse input file into slice of []byte
func readInputs(filename string) [][]byte {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open %s", filename)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var inputs [][]byte
	for scanner.Scan() {
		// Represent each line as a slice of bytes
		line := []byte(scanner.Text())
		inputs = append(inputs, line)
	}
	return inputs
}
