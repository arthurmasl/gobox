package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

type Part struct {
	start, end int
}

type StatMap map[string]Stat

type Stat struct {
	min, max, sum, count int
}

func startProfiling() func() {
	fmt.Println("Profiling...")

	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		panic(err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		panic(err)
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}

func main() {
	workers := 12

	input := generateInput(10_000_000)

	stopProfiling := startProfiling()
	defer stopProfiling()

	t1 := time.Now()
	parts := splitString(len(input), workers)

	partChannel := make(chan StatMap)
	for _, part := range parts {
		go processPart(&input, &part, partChannel)
	}

	totals := make(StatMap)
	for range parts {
		chunk := <-partChannel
		processChunk(&chunk, &totals)
	}

	close(partChannel)
	fmt.Println(time.Since(t1), len(totals))
}

func processChunk(chunk *StatMap, totals *StatMap) {
	for name, stat := range *chunk {
		d, ok := (*chunk)[name]

		if !ok {
			d = Stat{stat.min, stat.max, stat.sum, stat.count}
			continue
		}

		d.min = min(d.min, stat.min)
		d.max = min(d.max, stat.max)
		d.sum += stat.sum
		d.count += stat.count

		(*totals)[name] = d
	}
}

func processPart(input *[]string, part *Part, ch chan StatMap) {
	chunk := (*input)[part.start:part.end]
	data := make(StatMap)

	for _, line := range chunk {
		name, numStr, found := strings.Cut(line, ":")
		num, _ := strconv.Atoi(numStr)
		if !found {
			continue
		}

		d, ok := data[name]
		if !ok {
			d = Stat{num, num, num, 1}
		} else {
			d.min = min(d.min, num)
			d.max = min(d.max, num)
			d.sum += num
			d.count++
		}

		data[name] = d
	}

	ch <- data
}

func splitString(inputSize int, numParts int) []Part {
	parts := make([]Part, 0)

	partSize := inputSize / numParts
	leftover := inputSize % numParts

	start := 0
	end := start + partSize
	for i := range numParts {
		if i == numParts-1 {
			end += leftover
		}

		parts = append(parts, Part{start, end})
		start = end
		end = end + partSize
	}

	return parts
}

func generateInput(size int32) []string {
	input := make([]string, size)

	for i := range size {
		id := rand.Int32N(10000)
		num := rand.Int32N(30)
		input[i] = fmt.Sprintf("Id%d:%d", id, num)
	}

	return input
}
