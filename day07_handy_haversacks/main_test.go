package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	rules := parseInputs(input)

	want := 302
	var bags []*Tree
	for desc := range rules {
		bags = append(bags, createTree(rules, desc))
	}
	got := searchTrees(bags, "shiny gold")

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	rules := parseInputs(input)

	want := 4165
	shinyGold := createTree(rules, "shiny gold")
	got := total(shinyGold.root) - 1

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
		rules := parseInputs(input)
		var bags []*Tree
		for desc := range rules {
			bags = append(bags, createTree(rules, desc))
		}
		searchTrees(bags, "shiny gold")
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")

	for n := 0; n < b.N; n++ {
		rules := parseInputs(input)
		shinyGold := createTree(rules, "shiny gold")
		total(shinyGold.root)
	}
}
