package main

import (
	"os"
	"testing"

	"gobox/cmd/internal/solutions"
)

var (
	input, _  = os.Open("./small.txt")
	output, _ = os.ReadFile("./small.out")
)

func Test1BRC(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		result, cities := ExecuteSolution(small)

		if result != string(output)[:len(output)-1] || cities != small.cities {
			t.Fatalf("test failed")
		}
	})

	defer input.Close()
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solutions.Solution1(input, small.rows)
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solutions.Solution2(input, small.rows)
	}
}
