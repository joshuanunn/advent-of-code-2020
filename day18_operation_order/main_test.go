package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")

	want := 9535936849815
	precedence := map[string]int{
		LPAREN:   -1, // needed for parseInfix algorithm
		RPAREN:   -1, // needed for parseInfix algorithm
		PLUS:     1,
		MINUS:    1,
		MULTIPLY: 1,
		DIVIDE:   1,
	}
	got := processPart(input, precedence)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")

	want := 472171581333710
	precedence := map[string]int{
		LPAREN:   -1, // needed for parseInfix algorithm
		RPAREN:   -1, // needed for parseInfix algorithm
		PLUS:     5,
		MINUS:    1,
		MULTIPLY: 1,
		DIVIDE:   1,
	}
	got := processPart(input, precedence)

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
		precedence := map[string]int{
			LPAREN:   -1, // needed for parseInfix algorithm
			RPAREN:   -1, // needed for parseInfix algorithm
			PLUS:     1,
			MINUS:    1,
			MULTIPLY: 1,
			DIVIDE:   1,
		}
		processPart(input, precedence)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		precedence := map[string]int{
			LPAREN:   -1, // needed for parseInfix algorithm
			RPAREN:   -1, // needed for parseInfix algorithm
			PLUS:     5,
			MINUS:    1,
			MULTIPLY: 1,
			DIVIDE:   1,
		}
		processPart(input, precedence)
	}
}
