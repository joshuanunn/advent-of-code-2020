package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	pattern = regexp.MustCompile(`(acc|jmp|nop)\s([+-])([0-9]+)`)
)

type Instruction struct {
	op   string
	val  uint16
	sign bool
}

func (i *Instruction) value() int {
	if i.sign {
		return int(i.val)
	}
	return -1 * int(i.val)
}

func main() {
	// Read inputs into slice of *Instruction
	input := readInputs("input.txt")
	instructions := parseInputs(input)

	// Part 1 - find value of accumulator at point we get a repeated instruction.
	acc, ok := execute(instructions)
	fmt.Printf("Part 1 - accumulator value is %d and status should be false: [%v].\n", acc, ok)

	// Part 2 - mutate nop|jmp instructions to find way to break infinite loop.
	trials := mutateInstructions(instructions)
	acc, ok = executor(trials)
	if ok {
		fmt.Printf("Part 2 - solution found. The accumulator value at escape is %d.\n", acc)
	}
}

func executor(trials [][]*Instruction) (int, bool) {
	for _, trial := range trials {
		if acc, ok := execute(trial); ok {
			return acc, true
		}
	}
	return 0, false
}

// Generate a collection of mutated instructions, by replacing nop|jmp.
// Returns slice of *Instructions of length equal to the number of nop+jmp.
func mutateInstructions(instructions []*Instruction) [][]*Instruction {
	var mutated [][]*Instruction
	for x := 0; x < len(instructions); x++ {
		op := instructions[x].op
		if op == "nop" || op == "jmp" {
			mutated = append(mutated, mutate(instructions, op, x))
		}
	}
	return mutated
}

// Make a copy of the original slice of Instruction pointers, and replace
// selected position with a new mutated *Instruction.
func mutate(original []*Instruction, ins string, pos int) []*Instruction {
	mutated := make([]*Instruction, len(original))
	copy(mutated, original)
	if ins == "nop" {
		ins = "jmp"
	} else {
		ins = "nop"
	}
	mutated[pos] = &Instruction{ins, original[pos].val, original[pos].sign}
	return mutated
}

// Execute instruction set according to defined rules. Returns
// the value of the accumulator, along with the status (below).
func execute(instructions []*Instruction) (int, bool) {
	if len(instructions) < 1 {
		log.Fatalf("instructions are zero length")
	}
	pointer := 0
	accumulator := 0
	history := map[int]bool{}
	for pointer < len(instructions) {
		current := instructions[pointer]
		switch current.op {
		case "nop":
			pointer++
		case "acc":
			accumulator += current.value()
			pointer++
		case "jmp":
			pointer += current.value()
		default:
			log.Fatalf("instruction [%v] not defined", current)
		}
		if _, ok := history[pointer]; ok {
			break
		}
		history[pointer] = true
	}
	// If pointer has escaped the instruction stack, then success!
	if pointer >= len(instructions) {
		return accumulator, true
	}
	// Else we have hit an instruction twice == infinite loop.
	return accumulator, false
}

// Read in file and parse into a slice of *Instruction.
func parseInputs(inputs []string) []*Instruction {
	var instructions = []*Instruction{}
	for _, line := range inputs {
		extract := pattern.FindAllStringSubmatch(line, -1)[0]
		if len(extract) != 4 {
			log.Fatalln("failed to parse instruction line")
		}
		val, err := strconv.Atoi(extract[3])
		if err != nil {
			log.Fatalf("failed to parse int")
		}
		sign := false
		if extract[2] == "+" {
			sign = true
		}
		ins := &Instruction{extract[1], uint16(val), sign}
		instructions = append(instructions, ins)
	}
	return instructions
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
