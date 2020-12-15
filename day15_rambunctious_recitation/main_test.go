package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []int{1, 20, 8, 12, 0, 14}

	want := 492
	got := solve(input, 2020)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []int{1, 20, 8, 12, 0, 14}

	want := 63644
	got := solve(input, 30000000)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := []int{1, 20, 8, 12, 0, 14}
	for n := 0; n < b.N; n++ {
		solve(input, 2020)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := []int{1, 20, 8, 12, 0, 14}
	for n := 0; n < b.N; n++ {
		solve(input, 30000000)
	}
}
