package search

import (
	"container/list"
)

func bfs(lines []string, start Pos) {
	queue := list.New()

	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		element := queue.Front()
		elementPos := element.Value.(Pos)

		queue.Remove(element)

		draw()

		for _, neighbor := range getNeighbors(lines, elementPos) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}
}
