package main

import (
	"fmt"
	"sync"
	"time"
)

// The worker function now takes a unique ID to distinguish each worker.
func work(id int, wg *sync.WaitGroup) {
	// Defer the call to wg.Done() so it runs right before the function exits.
	defer wg.Done()

	fmt.Printf("Worker %d is working...\n", id)
	time.Sleep(time.Second) // Simulate a task taking time.
	fmt.Printf("Worker %d is done with work.\n", id)
}

func main() {
	var wg sync.WaitGroup

	// We will launch 3 workers, so we add 3 to the WaitGroup counter.
	const numWorkers = 3
	wg.Add(numWorkers)

	// Use a for loop to launch multiple goroutines.
	for i := 1; i <= numWorkers; i++ {
		// Launch a goroutine for each worker.
		// The loop variable 'i' is passed as an argument to avoid a common bug
		// where all goroutines would use the same 'i' value.
		go work(i, &wg)
	}

	// Wait here until the WaitGroup's counter is zero.
	fmt.Println("Waiting for all workers to finish...")
	wg.Wait()

	fmt.Println("\nMain goroutine has detected that all workers are complete.")
}
