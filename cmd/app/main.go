package main

import (
	"fmt"
	"os"
	"time"

	"gobox/cmd/internal/solution1"
	"gobox/cmd/internal/solution2"
)

type Solution struct {
	name   string
	rows   int
	cities int
}

var (
	small   = Solution{rows: 12, name: "small", cities: 10}
	tenmil  = Solution{rows: 10_000_000, name: "tenmils", cities: 413}
	billion = Solution{rows: 1_000_000_000, name: "billion", cities: 413}
)

var (
	solutionFn   = solution2.Execute
	solutionCase = tenmil
)

func main() {
	fmt.Fprintln(os.Stdout, []any{solution1.Execute, solution2.Execute}...)

	ExecuteSolution(solutionCase)
}

func ExecuteSolution(solution Solution) (string, int) {
	file, _ := os.Open(fmt.Sprintf("cmd/app/%v.txt", solution.name))
	defer file.Close()

	fmt.Printf("Name: %v, Rows: %v\n", solution.name, solution.rows)

	t1 := time.Now()
	str, cities := solutionFn(file, solution.rows)
	if cities != solution.cities {
		fmt.Printf("-Wrong solution!, got %v cities, when expected %v\n", cities, solution.cities)
	}
	fmt.Printf("Cities: %v, Expected: %v\n", cities, solution.cities)
	fmt.Printf("Execution time: %v\n", time.Since(t1))

	return str, cities
}
