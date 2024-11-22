package solutions

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func Solution2(file *os.File, rows int) (string, int) {
	bufferSize := 2048 * 2048

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	data := make(DataMap)

	workers := runtime.GOMAXPROCS(runtime.NumCPU())
	chunkSize := rows / workers
	chunks := make(chan []string, workers)

	fmt.Printf("Workers: %v, Chunk Size: %v\n", workers, chunkSize)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, bufferSize), bufferSize)

	for range workers {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for chunk := range chunks {
				processChunk2(chunk, &data, &mu)
			}
		}()
	}

	lines := make([]string, 0, chunkSize)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		if len(lines) == chunkSize {
			chunks <- lines
			lines = nil
		}
	}

	if len(lines) > 0 {
		fmt.Printf("Left lines: %v\n", len(lines))
		chunks <- lines
	}

	close(chunks)
	wg.Wait()

	index := 0
	dataArr := make([]string, len(data))

	for key, value := range data {
		mean := value.sum / float64(value.count)
		dataArr[index] = fmt.Sprintf("%v=%.1f/%.1f/%.1f", key, value.min, mean, value.max)
		index++
	}

	slices.Sort(dataArr)
	result := fmt.Sprintf("{%v}", strings.Join(dataArr, ", "))

	return result, len(data)
}

func processChunk2(lines []string, data *DataMap, mu *sync.Mutex) {
	// mu.Lock()
	// defer mu.Unlock()

	for _, line := range lines {
		parts := strings.Split(line, ";")
		name := parts[0]
		temp, _ := strconv.ParseFloat(parts[0], 64)

		d := (*data)[name]
		if d == nil {
			(*data)[name] = &Data{
				min:   temp,
				max:   temp,
				sum:   temp,
				count: 1,
			}
		} else {
			d.min = min(d.min, temp)
			d.max = max(d.min, temp)
			d.sum += temp
			d.count++
		}
	}
}
