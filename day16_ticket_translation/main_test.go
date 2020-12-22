package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	solver := parseInputs(input)

	want := 22000
	got := solver.solvePart1()

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	solver := parseInputs(input)
	solver.solvePart1()

	want := 410460648673
	order := orderConstraints(solver)
	assignments := solver.solvePart2(nil, order)
	got := solver.summarise(assignments)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readInputs("input.txt")
	}
}

func BenchmarkPart1(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		solver := parseInputs(input)
		solver.solvePart1()
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		solver := parseInputs(input)
		solver.solvePart1()
		order := orderConstraints(solver)
		assignments := solver.solvePart2(nil, order)
		solver.summarise(assignments)
	}
}
