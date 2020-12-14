package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")

	want := 926
	sids := parseInputs(input)
	_, got := intRange(sids)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")

	want := 657
	sids := parseInputs(input)
	sidMin, sidMax := intRange(sids)
	got := findMissing(sids, sidMin, sidMax)[0]

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
		sids := parseInputs(input)
		intRange(sids)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		sids := parseInputs(input)
		sidMin, sidMax := intRange(sids)
		findMissing(sids, sidMin, sidMax)
	}
}
