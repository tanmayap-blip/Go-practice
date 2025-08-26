package main

import (
	"fmt"
	"sync"
	"time"
)

func worke(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("Started the job with worker", id, "with the job", j)
		time.Sleep(time.Second)
		fmt.Println("Completed the job with worker", id, "and completed ", j)
		results <- j
	}

}

func main() {
	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worke(i, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i

	}
	close(jobs)
	wg.Wait()
	close(results)

	for j := range results {
		fmt.Println("results", j)
	}

}
