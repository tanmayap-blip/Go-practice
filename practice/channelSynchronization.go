package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool) {
	fmt.Println("Worker is working ...")
	time.Sleep(time.Second)
	fmt.Println("Worker is done with work ")
	ch <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	fmt.Println("has worker completed the work: ", <-done)
}
