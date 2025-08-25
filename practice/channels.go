package main

import "fmt"

func multiplyWithChannel(ch chan int) {
	fmt.Println(100 * <-ch)
}

func main() {

	ch := make(chan int)
	fmt.Println("Hello from main")
	go multiplyWithChannel(ch)
	ch <- 9
	fmt.Println("Bye from main")
}
