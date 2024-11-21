package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open("cmd/app/testdata2.txt")
	// file, _ := os.Open("cmd/app/measurements.txt")
	defer file.Close()

	t1 := time.Now()
	execute(file)
	fmt.Println(time.Since(t1))
}

type Data struct {
	min, max, sum float64
	count         int64
}

func execute(file *os.File) string {
	scanner := bufio.NewScanner(file)
	data := make(map[string]*Data)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		name := parts[0]
		temp, _ := strconv.ParseFloat(parts[1], 64)

		d := data[name]
		if d == nil {
			data[name] = &Data{
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

	index := 0
	strs := make([]string, len(data))

	for key, value := range data {
		mean := value.sum / float64(value.count)
		strs[index] = fmt.Sprintf("%v=%v/%v/%v", key, value.min, mean, value.max)
		index++
	}

	slices.Sort(strs)
	result := fmt.Sprintf("{%v}", strings.Join(strs, ", "))

	return result
}
