package main

import (
	"testing"
)

func benchmarkProdDiffs(i int, b *testing.B) {
	var input []int
	for x := 0; x < i; x++ {
		input = append(input, x)
	}
	for n := 0; n < b.N; n++ {
		prodDiffs(input)
	}
}

func benchmarkSumBranches(i int, b *testing.B) {
	var input []int
	for x := 0; x < i; x++ {
		input = append(input, x)
	}
	for n := 0; n < b.N; n++ {
		sumBranches(input)
	}
}

func BenchmarkProdDiffsData(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		prodDiffs(input)
	}
}

func BenchmarkSumBranchesData(b *testing.B) {
	input := readInputs("input.txt")
	for n := 0; n < b.N; n++ {
		sumBranches(input)
	}
}

func BenchmarkProdDiffs10(b *testing.B)   { benchmarkProdDiffs(10, b) }
func BenchmarkProdDiffs100(b *testing.B)  { benchmarkProdDiffs(100, b) }
func BenchmarkProdDiffs1000(b *testing.B) { benchmarkProdDiffs(1000, b) }

func BenchmarkSumBranches10(b *testing.B)   { benchmarkSumBranches(10, b) }
func BenchmarkSumBranches100(b *testing.B)  { benchmarkSumBranches(100, b) }
func BenchmarkSumBranches1000(b *testing.B) { benchmarkSumBranches(1000, b) }
