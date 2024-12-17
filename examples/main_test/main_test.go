package main

import (
	"fmt"
	"sync"
	"testing"
)

const (
	CHANNELS    = 10
	NUMS_TO_SUM = 10000000
)

func main() {
	res := SumNumbers(NUMS_TO_SUM, CHANNELS)
	fmt.Println(res)
}

func SumNumbers(numbers int, threads int) int {
	wg := sync.WaitGroup{}
	ch := make(chan int, threads)

	for range threads {
		wg.Add(1)
		workerNums := numbers / threads

		go worker(&wg, ch, workerNums)
	}

	wg.Wait()
	close(ch)

	total := 0
	for res := range ch {
		total += res
	}

	return total
}

func worker(wg *sync.WaitGroup, c chan int, numsToWork int) {
	defer wg.Done()

	sum := 0
	for range numsToWork {
		sum += 1
	}

	c <- sum
}

// go test ./src
func TestSumNumbers(t *testing.T) {
	tests := []struct {
		num     int
		threads int
	}{
		{100, 1},
		{10000000, 10},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Testing number %v on %v threads", tt.num, tt.threads)
		t.Run(testname, func(t *testing.T) {
			res := SumNumbers(tt.num, tt.threads)
			if res != tt.num {
				t.Errorf("result %v is not equals intial number %v", res, tt.num)
			}
		})
	}
}

// go test ./src -bench=.
func BenchmarkSumNumbers(b *testing.B) {
	tests := []struct {
		num     int
		threads int
	}{
		{10000000, 1},
		{10000000, 3},
		{10000000, 5},
		{10000000, 10},
		{10000000, 12},
		{10000000, 50},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Testing number %v on %v threads", tt.num, tt.threads)
		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SumNumbers(tt.num, tt.threads)
			}
		})
	}
}
