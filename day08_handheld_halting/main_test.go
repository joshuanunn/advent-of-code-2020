package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")

	want := 2003
	got, _ := execute(input)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")

	want := 1984
	trials := mutateInstructions(input)
	got, _ := executor(trials)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		execute(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		trials := mutateInstructions(input)
		executor(trials)
	}
}
