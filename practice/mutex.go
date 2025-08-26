package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // Declare a Mutex
	var counter int

	fmt.Println("Running mutex-protected counter example...")

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Acquire the lock before accessing the shared resource
			mu.Lock()

			// Critical section: only one goroutine can be in here at a time
			counter++

			// Release the lock after we're done with the shared resource
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Final (Correct) Mutex Counter: %d\n", counter)
}
