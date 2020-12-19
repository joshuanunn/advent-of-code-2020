package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	inactive = byte('.') // false
	active   = byte('#') // true
)

// Supports 3 and 4 dimensions
type Coord [4]int

type Grid struct {
	gridCurr map[Coord]bool
	gridNext map[Coord]bool
}

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")

	// Part 1 - complete 6 rounds of growth in 3 dimensions.
	grid3D := parseInputs(input)
	grid3D.cycle(6, 3)
	count3D := grid3D.active()
	fmt.Printf("Part 1 - total cubes in active state for 3 dimensions = %d (after 6 rounds)\n", count3D)

	// Part 2 - complete 6 rounds of growth in 4 dimensions.
	grid4D := parseInputs(input)
	grid4D.cycle(6, 4)
	count4D := grid4D.active()
	fmt.Printf("Part 2 - total cubes in active state for 4 dimensions = %d (after 6 rounds)\n", count4D)
}

func (c Coord) extract() (int, int, int, int) {
	return c[0], c[1], c[2], c[3]
}

func (c Coord) compare(coord Coord) bool {
	cmpX := c[0] == coord[0]
	cmpY := c[1] == coord[1]
	cmpZ := c[2] == coord[2]
	cmpW := c[3] == coord[3]
	if cmpX && cmpY && cmpZ && cmpW {
		return true
	}
	return false
}

func (g *Grid) grow(coord Coord, dimension int) {
	if dimension == 0 {
		// Check if coord exists in grid, if not create it
		if _, ok := g.gridCurr[coord]; !ok {
			g.gridCurr[coord] = false
			g.gridNext[coord] = false
		}
		return
	}
	dim := coord[dimension-1]
	for d := dim - 1; d <= dim+1; d++ {
		newCoord := coord
		newCoord[dimension-1] = d
		g.grow(newCoord, dimension-1)
	}
}

func (g *Grid) neighbours(coord, ref Coord, dimension int) int {
	var count int
	if dimension == 0 {
		// Skip cell itself
		if coord.compare(ref) {
			return count
		}
		// If active, include in count
		if g.gridCurr[coord] {
			count++
		}
		return count
	}
	dim := coord[dimension-1]
	for d := dim - 1; d <= dim+1; d++ {
		newCoord := coord
		newCoord[dimension-1] = d
		count += g.neighbours(newCoord, ref, dimension-1)
	}
	return count
}

func (g *Grid) countNeighbours(coord Coord, dimension int) int {
	return g.neighbours(coord, coord, dimension)
}

func (g *Grid) update() {
	for coord, state := range g.gridNext {
		g.gridCurr[coord] = state
	}
	g.gridNext = make(map[Coord]bool)
}

func (g *Grid) cycle(rounds, dimensions int) {
	if dimensions <= 1 {
		log.Fatalf("dimensions must be greater than 1")
	}
	for i := 0; i < rounds; i++ {
		// Create working grid
		g.gridNext = copy(g.gridCurr)
		// Grow grid by 1 cell around all current cells
		for coord := range g.gridCurr {
			g.grow(coord, dimensions)
		}
		// Calculate next state of system using rules
		for coord, active := range g.gridCurr {
			count := g.countNeighbours(coord, dimensions)
			switch {
			case active && (count == 2):
				g.gridNext[coord] = true
			case active && (count == 3):
				g.gridNext[coord] = true
			case !active && (count == 3):
				g.gridNext[coord] = true
			default:
				g.gridNext[coord] = false
			}
		}
		// Store next state of system
		g.update()
	}
}

func (g *Grid) active() int {
	var count int
	for _, active := range g.gridCurr {
		if active {
			count++
		}
	}
	return count
}

func copy(original map[Coord]bool) map[Coord]bool {
	target := make(map[Coord]bool)
	for k, v := range original {
		target[k] = v
	}
	return target
}

func parseInputs(inputs []string) *Grid {
	// Read in 2d plane, assume (x,y); (z,w_=0
	var data = &Grid{}
	data.gridCurr = make(map[Coord]bool)
	data.gridNext = make(map[Coord]bool)

	for y, line := range inputs {
		if len(line) > 0 {
			for x, cube := range []byte(line) {
				xyz := Coord{x, y, 0, 0}
				if cube == active {
					data.gridCurr[xyz] = true
				} else {
					data.gridCurr[xyz] = false
				}
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
