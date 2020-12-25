package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	floor := parseInputs(input)

	want := 266
	coords := playPart1(floor)
	got := len(coords)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	floor := parseInputs(input)

	want := 3627
	coords := playPart1(floor)
	newCoords := playPart2(coords, 100)
	got := len(newCoords)

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
		floor := parseInputs(input)
		coords := playPart1(floor)
		_ = len(coords)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	floor := parseInputs(input)
	coords := playPart1(floor)
	for n := 0; n < b.N; n++ {
		newCoords := playPart2(coords, 100)
		_ = len(newCoords)
	}
}
