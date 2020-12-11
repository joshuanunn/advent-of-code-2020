package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	grid := parseInputs(input)

	want := 167
	got := slopeRun(grid, 3, 1)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	grid := parseInputs(input)

	want := 736527114
	trees1 := slopeRun(grid, 1, 1)
	trees2 := slopeRun(grid, 3, 1)
	trees3 := slopeRun(grid, 5, 1)
	trees4 := slopeRun(grid, 7, 1)
	trees5 := slopeRun(grid, 1, 2)
	got := trees1 * trees2 * trees3 * trees4 * trees5

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := readInputs("input.txt")
	grid := parseInputs(input)

	for n := 0; n < b.N; n++ {
		slopeRun(grid, 3, 1)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	grid := parseInputs(input)

	for n := 0; n < b.N; n++ {
		trees1 := slopeRun(grid, 1, 1)
		trees2 := slopeRun(grid, 3, 1)
		trees3 := slopeRun(grid, 5, 1)
		trees4 := slopeRun(grid, 7, 1)
		trees5 := slopeRun(grid, 1, 2)
		_ = trees1 * trees2 * trees3 * trees4 * trees5
	}
}
