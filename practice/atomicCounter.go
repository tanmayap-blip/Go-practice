package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	// Demonstrate the problem with a non-atomic counter
	var nonAtomicCounter int64
	fmt.Println("Running non-atomic counter example...")

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// This operation is not atomic and can lead to a race condition
			nonAtomicCounter++
		}()
	}

	wg.Wait()
	fmt.Printf("Final (Incorrect) Non-Atomic Counter: %d\n", nonAtomicCounter)
	fmt.Println("----------------------------------------")

	// Demonstrate the correct way with an atomic counter
	var atomicCounter int64
	fmt.Println("Running atomic counter example...")

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// This operation is atomic and is guaranteed to be safe
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Final (Correct) Atomic Counter: %d\n", atomicCounter)
}
