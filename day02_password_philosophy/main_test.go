package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 483
	got := validatePart1(data)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 482
	got := validatePart2(data)

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
		data := parseInputs(input)
		validatePart1(data)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		data := parseInputs(input)
		validatePart2(data)
	}
}
