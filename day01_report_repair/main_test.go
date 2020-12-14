package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 197451
	elements, _ := findIn(data, 0, 2020, 2)
	got := summarise(elements)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 138233720
	elements, _ := findIn(data, 0, 2020, 3)
	got := summarise(elements)

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
		elements, _ := findIn(data, 0, 2020, 2)
		summarise(elements)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		data := parseInputs(input)
		elements, _ := findIn(data, 0, 2020, 3)
		summarise(elements)
	}
}
