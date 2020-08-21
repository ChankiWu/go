package main

import (
	"fmt"
	"sync"
	"time"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// put message to pings channel
	ping(pings, "passed message")
	// put messages from pings to pongs channel
	pong(pings, pongs)

	fmt.Println(<-pongs)

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
