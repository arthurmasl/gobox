package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PlantMap struct {
	destination int
	source      int
	length      int
}

var (
	seeds  []int
	groups [][]PlantMap
)

func main() {
	dat, _ := os.ReadFile("src/input.txt")
	text := string(dat)

	scanner := bufio.NewScanner(strings.NewReader(text))
	index := 0
	groupIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		// collect seeds
		if index == 0 {
			chunks := strings.Split(line, " ")
			for _, chunk := range chunks[1:] {
				seed, _ := strconv.Atoi(chunk)
				seeds = append(seeds, seed)
			}
		}

		// create group
		if strings.Contains(line, "map:") {
			groupIndex++
			groups = append(groups, []PlantMap{})
			continue
		}

		// fill group
		if groupIndex > 0 && len(line) > 0 {
			nums := strings.Split(line, " ")

			destination, _ := strconv.Atoi(nums[0])
			source, _ := strconv.Atoi(nums[1])
			length, _ := strconv.Atoi(nums[2])

			plantMap := PlantMap{destination, source, length}

			groups[groupIndex-1] = append(groups[groupIndex-1], plantMap)
		}

		index++
	}

	fmt.Printf("%v\n", seeds)
	fmt.Printf("%v\n", groups)
}
