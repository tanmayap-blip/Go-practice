package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	n := 200
	var wg sync.WaitGroup
	ch := make(chan int, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // random delay
			r := rand.Intn(5) + 1                                        // rating 1–5
			ch <- r
		}()
	}

	wg.Wait()
	close(ch)

	sum := 0
	cnt := 0
	for r := range ch {
		sum += r
		cnt++
	}
	fmt.Printf("Avg rating: %.2f\n", float64(sum)/float64(cnt))
}
