package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Coord struct {
	x int
	y int
	z int
}

type Coords []*Coord

var (
	// Vector along x,y,z axes for a move in a cube-based hex coordinate system
	NORTHEAST = &Coord{x: 1, y: -1}
	EAST      = &Coord{x: 1, z: -1}
	SOUTHEAST = &Coord{y: 1, z: -1}
	SOUTHWEST = &Coord{x: -1, y: 1}
	WEST      = &Coord{x: -1, z: 1}
	NORTHWEST = &Coord{y: -1, z: 1}
)

func main() {
	input := readInputs("input.txt")
	floor := parseInputs(input)

	// Part 1 - calculate initial state of flooring and count number of black tiles.
	coords := playPart1(floor)
	fmt.Printf("Part 1 - intial number of black tiles = %d\n", len(coords))

	// Part 2 - using ruleset, iterate over 100 days from initial state and count black tiles.
	newCoords := playPart2(coords, 100)
	fmt.Printf("Part 2 - number of black tiles after 100 days = %d\n", len(newCoords))
}

func playPart2(coords map[string]*Coord, rounds int) map[string]*Coord {
	for x := 0; x < rounds; x++ {
		var newCoords = make(map[string]*Coord)

		for _, coord := range coords {
			checkAdjacent(coord, coords, newCoords)
		}

		// Update coords with new coords
		coords = make(map[string]*Coord)
		for key, value := range newCoords {
			coords[key] = value
		}
	}
	return coords
}

func checkAdjacent(coord *Coord, coords, newCoords map[string]*Coord) {
	// Construct keys for cells around current cell
	checkCoords := getOrbit(coord)

	// Count number of adjacent black cells
	var black int
	for _, crd := range checkCoords {
		// If cell exists in map, then is black
		if _, ok := coords[crd.getKey()]; ok {
			black++
			continue
		}
		// Else cell is white and should be checked
		var count int
		whiteCheck := getOrbit(crd)
		for k := 0; count <= 2 && k < 6; k++ {
			if _, ok := coords[whiteCheck[k].getKey()]; ok {
				count++
			}
		}
		// If count is 2 then "flip" [create] centre as a black cell
		if count == 2 {
			newCoords[crd.getKey()] = crd
		}
	}
	// Retain cell as black
	if !(black == 0 || black > 2) {
		newCoords[coord.getKey()] = coord
	}
}

func getOrbit(coord *Coord) []*Coord {
	// Construct the 6 orbiting cells around cell of interest
	return []*Coord{
		coord.shiftBy(NORTHEAST),
		coord.shiftBy(EAST),
		coord.shiftBy(SOUTHEAST),
		coord.shiftBy(SOUTHWEST),
		coord.shiftBy(WEST),
		coord.shiftBy(NORTHWEST),
	}
}

func playPart1(coordSets []Coords) map[string]*Coord {
	var mapCoords = make(map[string]*Coord)
	for _, coords := range coordSets {
		final := &Coord{}
		for _, shift := range coords {
			final.x += shift.x
			final.y += shift.y
			final.z += shift.z
		}
		// Build floor state in a map
		// Only record black tiles - remove if state flipped to white
		key := final.getKey()
		if _, ok := mapCoords[key]; ok {
			delete(mapCoords, key)
		} else {
			mapCoords[key] = final
		}
	}
	return mapCoords
}

func (c *Coord) getKey() string {
	return fmt.Sprintf("%d,%d,%d", c.x, c.y, c.z)
}

func (c *Coord) shiftBy(shift *Coord) *Coord {
	return &Coord{c.x + shift.x, c.y + shift.y, c.z + shift.z}
}

func decode(line string) Coords {
	var coords Coords
	for i := 0; i < len(line); i++ {
		switch {
		case line[i] == 'e':
			coords = append(coords, EAST)
		case line[i] == 'w':
			coords = append(coords, WEST)
		case line[i] == 'n':
			if i+1 >= len(line) {
				log.Fatalf("parsing failed - something wrong with input")
			}
			if line[i+1] == 'e' {
				coords = append(coords, NORTHEAST)
			} else if line[i+1] == 'w' {
				coords = append(coords, NORTHWEST)
			} else {
				log.Fatalf("parsing failed - something wrong with input")
			}
			i++
		case line[i] == 's':
			if i+1 >= len(line) {
				log.Fatalf("parsing failed - something wrong with input")
			}
			if line[i+1] == 'e' {
				coords = append(coords, SOUTHEAST)
			} else if line[i+1] == 'w' {
				coords = append(coords, SOUTHWEST)
			} else {
				log.Fatalf("parsing failed - something wrong with input")
			}
			i++
		default:
			log.Fatalf("parsing failed - something wrong with input")
		}
	}
	return coords
}

func parseInputs(inputs []string) []Coords {
	var coords []Coords

	for _, line := range inputs {
		if len(line) == 0 {
			continue
		}
		coords = append(coords, decode(line))
	}
	return coords
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
