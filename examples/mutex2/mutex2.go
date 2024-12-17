package main

import (
	"fmt"
	"sync"
)

type Command struct {
	id   int
	name string
}

type Commands struct {
	mutex sync.Mutex
	wg    sync.WaitGroup
	queue []Command
}

type Commander interface {
	Execute(string)
	Size() int
}

func (commands *Commands) Size() int {
	return len(commands.queue)
}

func (commands *Commands) Execute(name string) {
	commands.mutex.Lock()
	defer commands.mutex.Unlock()
	defer commands.wg.Done()

	command := Command{id: commands.Size(), name: name}
	commands.queue = append(commands.queue, command)
}

func main() {
	commands := Commands{}

	for i := range 20 {
		commands.wg.Add(1)
		go commands.Execute(fmt.Sprintf("command %v", i))
	}
	commands.wg.Wait()

	fmt.Println(commands.queue)
}
