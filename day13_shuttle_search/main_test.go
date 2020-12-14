package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	time, data := parseInputs(input)

	want := 171
	got := solvePart1(time, data)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	_, data := parseInputs(input)

	want := 539746751134958
	got, _ := solvePart2(data)

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
		time, data := parseInputs(input)
		solvePart1(time, data)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		_, data := parseInputs(input)
		solvePart2(data)
	}
}
