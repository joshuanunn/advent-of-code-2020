package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInputs("input.txt")
	foodSlice := parseInputs(input)

	want := 2573
	allergens := getAllergens(foodSlice)
	got := countNonAllergens(allergens, foodSlice)

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInputs("input.txt")
	foodSlice := parseInputs(input)

	want := "bjpkhx,nsnqf,snhph,zmfqpn,qrbnjtj,dbhfd,thn,sthnsg"
	allergens := getAllergens(foodSlice)
	assignments := assignAllergens(allergens)
	got := summarise(assignments)

	if got != want {
		t.Errorf("got %s; want %s", got, want)
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
		foodSlice := parseInputs(input)
		allergens := getAllergens(foodSlice)
		countNonAllergens(allergens, foodSlice)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		foodSlice := parseInputs(input)
		allergens := getAllergens(foodSlice)
		assignments := assignAllergens(allergens)
		summarise(assignments)
	}
}
