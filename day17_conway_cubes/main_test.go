package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	grid := parseInputs(input)

	want := 317
	grid.cycle(6, 3)
	got := grid.active()

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	grid := parseInputs(input)

	want := 1692
	grid.cycle(6, 4)
	got := grid.active()

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
		grid := parseInputs(input)
		grid.cycle(6, 3)
		grid.active()
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		grid := parseInputs(input)
		grid.cycle(6, 4)
		grid.active()
	}
}
