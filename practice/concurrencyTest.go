package main

//
//import (
//	"fmt"
//	"runtime"
//	"sync"
//	"time"
//)
//
//// The worker function that performs a simple, CPU-intensive task.
//// It receives work from the 'jobs' channel and sends results to the 'results' channel.
//func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
//	// Signal to the WaitGroup that this worker is done when the function exits.
//	defer wg.Done()
//
//	// Loop over the jobs channel until it is closed.
//	for j := range jobs {
//		// Simulate a CPU-intensive task by calculating the sum of numbers up to 'j'.
//		// This makes the task take a measurable amount of time.
//		sum := 0
//		for i := 0; i < j; i++ {
//			sum += i
//		}
//
//		// Print a message showing which worker is processing which job.
//		// You'll see these lines printed out of order, demonstrating concurrency.
//		fmt.Printf("Worker %d finished job %d\n", id, j)
//
//		// Send the result back to the results channel.
//		results <- sum
//	}
//}
//
//func main() {
//	// Get the number of CPU cores on the machine.
//	// This ensures the program scales to your specific hardware.
//	numCPUs := runtime.NumCPU()
//	fmt.Printf("Using %d CPU cores for concurrency.\n", numCPUs)
//
//	const numJobs = 100 // Total number of jobs to process.
//
//	// Create channels for jobs and results. We use a buffered channel for the jobs
//	// to avoid blocking the main goroutine as it sends tasks.
//	jobs := make(chan int, numJobs)
//	results := make(chan int, numJobs)
//
//	var wg sync.WaitGroup
//
//	// Record the start time to measure performance.
//	startTime := time.Now()
//
//	// Start the worker pool (fan-out).
//	// We create one worker goroutine for each CPU core.
//	for w := 1; w <= numCPUs; w++ {
//		wg.Add(1) // Add one to the WaitGroup for each worker.
//		go worker(w, jobs, results, &wg)
//	}
//
//	// Send all the jobs to the jobs channel.
//	for j := 1; j <= numJobs; j++ {
//		jobs <- j // The main goroutine acts as the producer.
//	}
//	// Close the jobs channel to signal that no more work is coming.
//	close(jobs)
//
//	// Wait for all the worker goroutines to finish their work.
//	wg.Wait()
//
//	// Close the results channel after all workers are done.
//	close(results)
//
//	// Collect and process all the results from the results channel (fan-in).
//	totalResults := 0
//	for r := range results {
//		totalResults += r
//	}
//
//	// Print the total time taken and the final sum.
//	duration := time.Since(startTime)
//	fmt.Printf("\nTotal time taken: %s\n", duration)
//	fmt.Printf("Total results sum: %d\n", totalResults)
//}
