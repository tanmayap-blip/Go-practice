package main

import (
	"fmt"
	"sync"
)

// A struct to represent a read operation.
type readOp struct {
	key  int
	resp chan int
}

// A struct to represent a write operation.
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var wg sync.WaitGroup

	// Create channels for read and write operations.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// The stateful goroutine
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				// Process a read request
				read.resp <- state[read.key]
			case write := <-writes:
				// Process a write request
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Simulate concurrent read requests
	for r := 0; r < 5; r++ {
		wg.Add(1)
		go func(r int) {
			defer wg.Done()
			read := readOp{
				key:  r,
				resp: make(chan int)}
			reads <- read
			value := <-read.resp
			fmt.Printf("Read request %d, value: %d\n", r, value)
		}(r)
	}

	// Simulate concurrent write requests
	for w := 0; w < 5; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			write := writeOp{
				key:  w,
				val:  w,
				resp: make(chan bool)}
			writes <- write
			<-write.resp
			fmt.Printf("Write request %d complete.\n", w)
		}(w)
	}

	wg.Wait()
}
