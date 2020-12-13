package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 2453
	got, _ := converge(data, rulePart1, 4, 1000)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 2159
	got, _ := converge(data, rulePart2, 5, 1000)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		data := parseInputs(input)
		converge(data, rulePart1, 4, 1000)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		data := parseInputs(input)
		converge(data, rulePart2, 5, 1000)
	}
}
