package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)`)
var statsKeys = []string{
	"departure date",
	"departure location",
	"departure platform",
	"departure station",
	"departure time",
	"departure track",
}

type Constraint struct {
	minA int
	maxA int
	minB int
	maxB int
}

type Solver struct {
	descriptors   []string
	columns       []int
	constraints   map[string]*Constraint
	yourTicket    []int
	nearbyTickets [][]int
}

type Result struct {
	column int
	count  int
}

func main() {
	input := readInputs("input.txt")
	solver := parseInputs(input)

	// Part 1 - calculate ticket scanning error rate, and filter out bad tickets.
	total := len(solver.nearbyTickets)
	errorRate := solver.solvePart1()
	final := len(solver.nearbyTickets)
	fmt.Printf("Part 1 - %d/%d valid tickets found. Ticket scanning error rate = %d\n", final, total, errorRate)

	// Part 2 - assign all columns, and return product of departure related fields on your ticket.
	order := orderConstraints(solver)
	assignments := solver.solvePart2(nil, order)
	product := solver.summarise(assignments)
	fmt.Printf("Part 2 - product of departure related fields on your ticket = %d\n", product)
}

func (s *Solver) summarise(assignment map[string]int) int {
	product := 1
	for _, key := range statsKeys {
		column := assignment[key]
		product *= s.yourTicket[column]
	}
	return product
}

// Solve using backtracking algorithm. Requuires order information to be efficient.
func (s *Solver) solvePart2(assignment map[string]int, order []*Result) map[string]int {
	if assignment == nil {
		assignment = make(map[string]int)
	}

	// Base case - if we have assigned all fields, then complete
	if len(assignment) == len(s.descriptors) {
		return assignment
	}

	// Get list of unassigned descriptors at this point, and take first value
	unassignedDesc := s.getUnassigned(assignment)
	first := unassignedDesc[0]

	// Consider every possible ccolumn descriptor combo
	for _, column := range order {
		// Check constraint
		if s.consistent(first, column.column) {
			assignmentLocal := copy(assignment)
			assignmentLocal[first] = column.column
			result := s.solvePart2(assignmentLocal, order)
			if result != nil {
				return result
			}
		}
	}
	return nil
}

// Check constraints met for every ticket in specified column
func (c *Constraint) satisfied(column int, tickets [][]int) bool {
	for _, ticket := range tickets {
		if !c.validate(ticket[column]) {
			return false
		}
	}
	return true
}

func (s *Solver) consistent(descriptor string, assignment int) bool {
	// Lookup constraint
	constraint := s.constraints[descriptor]
	if constraint.satisfied(assignment, s.nearbyTickets) {
		return true
	}
	return false
}

func copy(original map[string]int) map[string]int {
	var target = make(map[string]int)
	for k, v := range original {
		target[k] = v
	}
	return target
}

func (s *Solver) getUnassigned(assignment map[string]int) []string {
	var unassignedDesc []string

	// Find unassigned descriptors and columns
	for _, descriptor := range s.descriptors {
		if _, ok := assignment[descriptor]; !ok {
			unassignedDesc = append(unassignedDesc, descriptor)
		}
	}

	return unassignedDesc
}

func orderConstraints(solver *Solver) []*Result {
	// Consider each column of every ticket to see which constraints fit
	// If we get a non-match then abort
	var order = []*Result{}

	for col := 0; col < len(solver.constraints); col++ {
		// Check all constraints
		var count int
		for _, constraint := range solver.constraints {
			valid := true
			for _, ticket := range solver.nearbyTickets {
				if !constraint.validate(ticket[col]) {
					valid = false
				}
			}
			if valid {
				count++
			}
		}
		order = append(order, &Result{column: col, count: count})
	}

	// Sort order ascending (significantly improves peformance of backtracking algorithm)
	sort.Slice(order, func(i, j int) bool {
		if order[i].count < order[j].count {
			return true
		}
		return order[i].count < order[j].count
	})
	return order
}

func (s *Solver) solvePart1() int {
	var errorRate int
	var validTickets [][]int
	for _, ticket := range s.nearbyTickets {
		valid := 0
		for _, attribute := range ticket {
			if s.validateAll(attribute) {
				valid++
			} else {
				errorRate += attribute
			}
		}
		if valid == len(ticket) {
			validTickets = append(validTickets, ticket)
		}
	}
	s.nearbyTickets = validTickets
	return errorRate
}

func (c *Constraint) validate(val int) bool {
	conA := (c.minA <= val) && (val <= c.maxA)
	conB := (c.minB <= val) && (val <= c.maxB)
	if conA || conB {
		return true
	}
	return false
}

func (s *Solver) validateAll(val int) bool {
	var valid bool
	for _, desc := range s.descriptors {
		c := s.constraints[desc]
		if c.validate(val) {
			valid = true
		}
	}
	return valid
}

func parseInt(line string) int {
	val, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("could not parse int")
	}
	return val
}

func parseTicket(line string) []int {
	var ticket []int
	values := strings.Split(line, ",")
	if len(values) == 0 {
		log.Fatalf("ticket format incorrect")
	}
	for _, value := range values {
		ticket = append(ticket, parseInt(value))
	}
	return ticket
}

func parseInputs(inputs []string) *Solver {
	solver := &Solver{}
	solver.constraints = make(map[string]*Constraint)

	for i := 0; i < len(inputs); i++ {
		match := pattern.FindStringSubmatch(inputs[i])
		if len(match) == 6 {
			desc := match[1]
			minA := parseInt(match[2])
			maxA := parseInt(match[3])
			minB := parseInt(match[4])
			maxB := parseInt(match[5])

			solver.descriptors = append(solver.descriptors, desc)
			solver.columns = append(solver.columns, i)
			solver.constraints[desc] = &Constraint{minA, maxA, minB, maxB}
			continue
		}
		if strings.Contains(inputs[i], "your ticket:") {
			i++
			solver.yourTicket = parseTicket(inputs[i])
			continue
		}
		if strings.Contains(inputs[i], "nearby tickets:") {
			i++
			for i < len(inputs) {
				solver.nearbyTickets = append(solver.nearbyTickets, parseTicket(inputs[i]))
				i++
			}
			continue
		}
	}
	return solver
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
