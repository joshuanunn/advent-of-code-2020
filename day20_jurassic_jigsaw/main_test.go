package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")

	tiles, count, _ := parseInputs(input)
	groupedPieces := definePieces(tiles, count)

	want := 51214443014783
	got := 1
	for _, cornerID := range groupedPieces[2] {
		got *= cornerID.id
	}

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")

	tiles, count, links := parseInputs(input)
	groupedPieces := definePieces(tiles, count)

	want := 2065
	firstCorner := groupedPieces[2][0]
	grid := buildGrid(tiles, firstCorner.id, links)
	total := countMonsters(grid)
	imageCorrection := total * 15
	got := grid.count() - imageCorrection

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
		tiles, count, _ := parseInputs(input)
		groupedPieces := definePieces(tiles, count)
		product := 1
		for _, cornerID := range groupedPieces[2] {
			product *= cornerID.id
		}

	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		tiles, count, links := parseInputs(input)
		groupedPieces := definePieces(tiles, count)
		firstCorner := groupedPieces[2][0]
		grid := buildGrid(tiles, firstCorner.id, links)
		total := countMonsters(grid)
		imageCorrection := total * 15
		_ = grid.count() - imageCorrection
	}
}
