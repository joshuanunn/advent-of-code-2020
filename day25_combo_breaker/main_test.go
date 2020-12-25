package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	cardPK, doorPK := parseInputs(input)

	cardLoopSize := extractLoopSize(7, cardPK)
	doorLoopSize := extractLoopSize(7, doorPK)

	want := 6198540
	gotA := extractEncryptionKey(cardLoopSize, doorPK)
	gotB := extractEncryptionKey(doorLoopSize, cardPK)

	if gotA != want {
		t.Errorf("got %d; want %d", gotA, want)
	}
	if gotB != want {
		t.Errorf("got %d; want %d", gotB, want)
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
		cardPK, doorPK := parseInputs(input)
		cardLoopSize := extractLoopSize(7, cardPK)
		doorLoopSize := extractLoopSize(7, doorPK)
		_ = extractEncryptionKey(cardLoopSize, doorPK)
		_ = extractEncryptionKey(doorLoopSize, cardPK)
	}
}
