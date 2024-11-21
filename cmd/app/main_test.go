package main

import (
	"os"
	"testing"
)

var (
	input, _  = os.Open("./testdata.txt")
	output, _ = os.ReadFile("./testdata.out")
)

func Test1BRC(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		result := execute(input)

		if result != string(output)[:len(output)-1] {
			t.Fatalf("test failed")
		}
	})
}

func Benchmark1BRC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		execute(input)
	}
}
