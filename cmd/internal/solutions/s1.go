package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type data1 struct {
	min, max, sum float64
	count         int64
}

func Solution1(file *os.File, rows int) (string, int) {
	scanner := bufio.NewScanner(file)
	data := make(map[string]*data1)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		name := parts[0]
		temp, _ := strconv.ParseFloat(parts[1], 64)

		d := data[name]
		if d == nil {
			data[name] = &data1{
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

	return result, len(strs)
}
