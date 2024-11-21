package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("cmd/app/measurements.txt")
	defer file.Close()

	execute(file)
}

func execute(file *os.File) string {
	scanner := bufio.NewScanner(file)
	data := make(map[string][]float32)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		name := parts[0]
		floatTemp, _ := strconv.ParseFloat(parts[1], 32)
		temp := float32(floatTemp)

		_, ok := data[name]
		if !ok {
			data[name] = make([]float32, 3)
		}

		// min
		if data[name][0] == 0 || data[name][0] > temp {
			data[name][0] = temp
		}
		// max
		if data[name][2] < temp {
			data[name][2] = temp
		}
		// mean
		data[name][1] = (data[name][0] + data[name][2]) / 2
	}

	strs := make([]string, len(data))
	index := 0
	for key, value := range data {
		strs[index] = fmt.Sprintf("%v=%v/%v/%v", key, value[0], value[1], value[2])
		index++
	}
	slices.Sort(strs)
	result := fmt.Sprintf("{%v}", strings.Join(strs, ", "))

	return result
}
