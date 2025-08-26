package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch1 <- "hey this is first channel"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "hey this is second channel"

	}()

	// for once condition exexcution

	//select {
	//case msg1 := <-ch1:
	//	fmt.Println("there is a message from channel 1", msg1)
	//case msg2 := <-ch2:
	//	fmt.Println("there is a message from channel 2 :", msg2)
	//default:
	//	fmt.Println(" No channel is ready till now ")
	//
	//}

	// for mutiple command execution

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("there is a message from channel 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("there is a message from channel 2 :", msg2)
		}
	}

}
