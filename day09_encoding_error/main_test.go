package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	preamble := 25

	want := 14144619
	idx, _ := findInvalid(input, preamble)
	got := input[idx]

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	target := 14144619

	want := 1766397
	got, _ := findContiguous(input, target)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := readInputs("input.txt")
	preamble := 25

	for n := 0; n < b.N; n++ {
		findInvalid(input, preamble)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	target := 14144619

	for n := 0; n < b.N; n++ {
		findContiguous(input, target)
	}
}
