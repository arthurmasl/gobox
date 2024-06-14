package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	rr := 85
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				cmd.Run()
				rr += 5
				fmt.Printf("%v\n", rr)
			}
		}
	}()

	<-done
	fmt.Println("stop")
}
