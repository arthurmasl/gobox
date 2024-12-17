package search

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type AlgType int

const (
	BFS AlgType = iota
	DFS
)

type Pos struct {
	x, y int
}

var (
	directions = []Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited    = make(map[Pos]bool)
	canvas     *[]string
)

func Start(alg AlgType) {
	input, _ := os.ReadFile("draft/graphs/example")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	canvas = createCanvas(lines)

	for y, row := range lines {
		for x := range row {
			pos := Pos{x, y}
			if visited[pos] {
				continue
			}

			if alg == BFS {
				bfs(lines, pos)
			} else {
				dfs(lines, pos)
			}
		}
	}
}

func getNeighbors(lines []string, pos Pos) []Pos {
	id := string(lines[pos.y][pos.x])
	neighbors := make([]Pos, 0)

	for _, dir := range directions {
		neighborPos := Pos{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := getSafeValue(lines, neighborPos)

		if ok && neighborId == id {
			neighbors = append(neighbors, neighborPos)
		}
	}

	return neighbors
}

func getSafeValue(arr []string, pos Pos) (string, bool) {
	if pos.y >= 0 && pos.y < len(arr) && pos.x >= 0 && pos.x < len(arr[0]) {
		return string(arr[pos.y][pos.x]), true
	}
	return "", false
}

func createCanvas(lines []string) *[]string {
	return &lines
}

var reset = "\033[0m"

func draw() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println()
	for y, row := range *canvas {
		for x, char := range row {
			if visited[Pos{x, y}] {
				color := getColorByLetter(char)
				fmt.Print(color + string(char) + reset)
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}

	time.Sleep(time.Millisecond * 10)
}

func getColorByLetter(letter rune) string {
	colorCode := (int(letter) - int('A')) + 16
	return fmt.Sprintf("\033[38;5;%dm", colorCode%256)
}
