package main

import (
	"testing"
)

const tolerance = 0.00000001

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 381.0
	ship := &Ship{compass: 90.0}
	ship.moveCompass(data)
	got := ship.dist(0.0, 0.0)

	if got-want > tolerance {
		t.Errorf("got %.1f; want %.1f", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	data := parseInputs(input)

	want := 28591.0
	ship := &Ship{waypointE: 10, waypointN: 1}
	ship.moveWaypoint(data)
	got := ship.dist(0.0, 0.0)

	if got-want > tolerance {
		t.Errorf("got %.1f; want %.1f", got, want)
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
		ship := &Ship{compass: 90.0}
		ship.moveCompass(data)
		ship.dist(0.0, 0.0)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		data := parseInputs(input)
		ship := &Ship{waypointE: 10, waypointN: 1}
		ship.moveWaypoint(data)
		ship.dist(0.0, 0.0)
	}
}
