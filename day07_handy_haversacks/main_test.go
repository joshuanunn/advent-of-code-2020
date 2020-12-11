package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")

	want := 302
	var bags []*Tree
	for desc := range input {
		bags = append(bags, createTree(input, desc))
	}
	got := searchTrees(bags, "shiny gold")

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")

	want := 4165
	shinyGold := createTree(input, "shiny gold")
	got := total(shinyGold.root) - 1

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		var bags []*Tree
		for desc := range input {
			bags = append(bags, createTree(input, desc))
		}
		searchTrees(bags, "shiny gold")
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		shinyGold := createTree(input, "shiny gold")
		total(shinyGold.root)
	}
}
