package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	players := parseInputs(input)

	want := 32199
	_, deck := playPart1(players)
	got := calcScore(deck)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	players := parseInputs(input)

	want := 33780
	_, deck := playPart2(players)
	got := calcScore(deck)

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
		players := parseInputs(input)
		_, deck := playPart1(players)
		calcScore(deck)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		players := parseInputs(input)
		_, deck := playPart2(players)
		calcScore(deck)
	}
}
