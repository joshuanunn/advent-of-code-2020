package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)
	preamble := 25

	want := 14144619
	idx, _ := findInvalid(data, preamble)
	got := data[idx]

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)
	target := 14144619

	want := 1766397
	got, _ := findContiguous(data, target)

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
	data := parseInputs(input)
	preamble := 25

	for n := 0; n < b.N; n++ {
		findInvalid(data, preamble)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	data := parseInputs(input)
	target := 14144619

	for n := 0; n < b.N; n++ {
		findContiguous(data, target)
	}
}
