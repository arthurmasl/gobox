package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"

	"gobox/internal/utils"
)

type loadBalancer struct {
	servers  []*server
	capacity int
}

type server struct {
	id   string
	ch   chan int
	data []*client
}

type client struct {
	id string
}

func newLoadBalancer(capacity int) *loadBalancer {
	return &loadBalancer{
		servers:  make([]*server, 0),
		capacity: capacity,
	}
}

func (balancer *loadBalancer) addServer(capacity int) {
	server := &server{
		id:   strconv.Itoa(len(balancer.servers)),
		ch:   make(chan int, capacity),
		data: make([]*client, 0, capacity),
	}

	balancer.servers = append(balancer.servers, server)
}

func (balancer *loadBalancer) addClient(id string) {
	foundServer := false

	for _, server := range balancer.servers {
		if len(server.data) < balancer.capacity {
			server.data = append(server.data, &client{id})
			foundServer = true
			break
		}
	}

	if !foundServer {
		balancer.addServer(balancer.capacity)
	}
}

func (balancer *loadBalancer) removeClient(id string) {
	for _, server := range balancer.servers {
		for i, client := range server.data {
			if client.id == id {
				server.data = append(server.data[:i], server.data[i+1:]...)
				break
			}
		}
	}
}

func (balancer *loadBalancer) print() {
	utils.ClearConsole()
	fmt.Printf("capacity: %v, servers: %v\n", balancer.capacity, len(balancer.servers))
	for _, server := range balancer.servers {
		fmt.Printf("  server: %v, clients: %v\n", server.id, len(server.data))

		for _, client := range server.data {
			fmt.Printf("    client: %v\n", client.id)
		}
	}
}

func main() {
	balancer := newLoadBalancer(8)

	for {
		id := strconv.Itoa(rand.N(30))
		balancer.addClient(id)
		time.AfterFunc(
			time.Duration(rand.N(10)*int(time.Second)),
			func() { balancer.removeClient(id) },
		)

		balancer.print()
		time.Sleep(time.Duration(rand.N(500) * int(time.Millisecond)))
	}
}
