package utils

import (
	"fmt"
	"iter"
	"os"
	"os/exec"
	"strings"
)

func UNUSED(x ...any) {}

func Assert(condition bool) {
	if !condition {
		panic("Assertion failed")
	}
}

func Window[Slice ~[]E, E any](slice Slice, size int) iter.Seq[Slice] {
	return func(yield func(Slice) bool) {
		for i := range slice[:len(slice)-size+1] {
			if !yield(slice[i : i+size]) {
				return
			}
		}
	}
}

func GetSafeValue(arr []string, x, y int) (byte, bool) {
	if y >= 0 && y < len(arr) && x >= 0 && x < len(arr[y]) {
		return arr[y][x], true
	}

	return 0, false
}

func GetLines(inputDir string, args ...string) []string {
	input, err := os.ReadFile("assets/" + inputDir + ".txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil
	}

	sep := "\n"
	if len(args) > 0 {
		sep = args[0]
	}

	lines := strings.Split(strings.TrimSpace(string(input)), sep)
	return lines
}

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
