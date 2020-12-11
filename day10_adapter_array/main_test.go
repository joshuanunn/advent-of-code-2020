package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")

	want := 2210
	got := prodDiffs(input)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")

	want := 7086739046912
	got := sumBranches(input)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		prodDiffs(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		sumBranches(input)
	}
}

func benchmarkPart1(i int, b *testing.B) {
	var input []int
	for x := 0; x < i; x++ {
		input = append(input, x)
	}
	for n := 0; n < b.N; n++ {
		prodDiffs(input)
	}
}

func benchmarkPart2(i int, b *testing.B) {
	var input []int
	for x := 0; x < i; x++ {
		input = append(input, x)
	}
	for n := 0; n < b.N; n++ {
		sumBranches(input)
	}
}

func Benchmark10Part1(b *testing.B)   { benchmarkPart1(10, b) }
func Benchmark100Part1(b *testing.B)  { benchmarkPart1(100, b) }
func Benchmark1000Part1(b *testing.B) { benchmarkPart1(1000, b) }

func Benchmark10Part2(b *testing.B)   { benchmarkPart2(10, b) }
func Benchmark100Part2(b *testing.B)  { benchmarkPart2(100, b) }
func Benchmark1000Part2(b *testing.B) { benchmarkPart2(1000, b) }
