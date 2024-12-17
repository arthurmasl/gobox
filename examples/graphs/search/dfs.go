package search

import (
	"container/list"
)

func dfs(lines []string, start Pos) {
	stack := list.New()
	stack.PushBack(start)

	for stack.Len() > 0 {
		element := stack.Back()
		elementPos := element.Value.(Pos)

		stack.Remove(element)

		if !visited[elementPos] {
			visited[elementPos] = true
			draw()
		}

		for _, neighbor := range getNeighbors(lines, elementPos) {
			if !visited[neighbor] {
				stack.PushBack(neighbor)
			}
		}
	}
}
