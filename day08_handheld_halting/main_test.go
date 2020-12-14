package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	instructions := parseInputs(input)

	want := 2003
	got, _ := execute(instructions)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	instructions := parseInputs(input)

	want := 1984
	trials := mutateInstructions(instructions)
	got, _ := executor(trials)

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
	instructions := parseInputs(input)

	for n := 0; n < b.N; n++ {
		execute(instructions)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	instructions := parseInputs(input)

	for n := 0; n < b.N; n++ {
		trials := mutateInstructions(instructions)
		executor(trials)
	}
}
