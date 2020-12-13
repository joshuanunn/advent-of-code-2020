package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	action byte
	value  float64
}

type Ship struct {
	compass   float64
	easting   float64
	northing  float64
	waypointE float64
	waypointN float64
}

var pattern = regexp.MustCompile(`^([NSEWLRF])(\d+)$`)

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")
	data := parseInputs(input)

	// Part 1 - calculate compass route and final manhattan displacement from start.
	ship1 := &Ship{compass: 90.0}
	ship1.moveCompass(data)
	dist1 := ship1.dist(0.0, 0.0)
	fmt.Printf("Part 1 - manhattan distance travelled: %.1f\n", dist1)

	// Part 2 - calculate waypoint route and final manhattan displacement from start.
	ship2 := &Ship{waypointE: 10, waypointN: 1}
	ship2.moveWaypoint(data)
	dist2 := ship2.dist(0.0, 0.0)
	fmt.Printf("Part 2 - manhattan distance travelled: %.1f\n", dist2)
}

// Calculate ship position using waypoint rules
func (s *Ship) moveWaypoint(commands []*Instruction) {
	for _, command := range commands {
		switch command.action {
		case 'N':
			s.waypointN += command.value
		case 'S':
			s.waypointN -= command.value
		case 'E':
			s.waypointE += command.value
		case 'W':
			s.waypointE -= command.value
		case 'L':
			s.waypointE, s.waypointN = rotate(s.waypointE, s.waypointN, command.value)
		case 'R':
			s.waypointE, s.waypointN = rotate(s.waypointE, s.waypointN, -1.0*command.value)
		case 'F':
			s.easting += command.value * s.waypointE
			s.northing += command.value * s.waypointN
		}
	}
}

// Calculate ship position using compass rules
func (s *Ship) moveCompass(commands []*Instruction) {
	for _, command := range commands {
		switch command.action {
		case 'N':
			s.northing += command.value
		case 'S':
			s.northing -= command.value
		case 'E':
			s.easting += command.value
		case 'W':
			s.easting -= command.value
		case 'L':
			s.compass -= command.value
		case 'R':
			s.compass += command.value
		case 'F':
			e, n := compassCoords(s.compass, command.value)
			s.easting += e
			s.northing += n
		}
	}
}

func (s *Ship) dist(easting, northing float64) float64 {
	return math.Abs(s.easting-easting) + math.Abs(s.northing-northing)
}

// convert compass bearing and displacement to rectangular coords
func compassCoords(bearing, displacement float64) (float64, float64) {
	theta := math.Mod(450.0-bearing, 360) * math.Pi / 180.0
	easting := displacement * math.Cos(theta)
	northing := displacement * math.Sin(theta)
	return easting, northing
}

// convert rectangular to polar coords, apply rotation (deg), and convert back
func rotate(easting, northing, rotation float64) (float64, float64) {
	r := math.Sqrt(math.Pow(easting, 2) + math.Pow(northing, 2))
	theta := math.Atan2(northing, easting)
	theta += rotation * math.Pi / 180.0
	eastingNew := r * math.Cos(theta)
	northingNew := r * math.Sin(theta)
	return eastingNew, northingNew
}

func parseInputs(inputs []string) []*Instruction {
	var data []*Instruction
	for _, line := range inputs {
		if len(line) > 0 {
			match := pattern.FindStringSubmatch(line)
			if len(match) != 3 {
				log.Fatalf("no regex match for line")
			}
			action := byte(match[1][0])
			value, err := strconv.ParseFloat(match[2], 64)
			if err != nil {
				log.Fatalf("could not parse int in line")
			}
			data = append(data, &Instruction{action, value})
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
