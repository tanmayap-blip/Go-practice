package main

import "fmt"

func multiplyWithChannel(ch chan int) {
	fmt.Println(100 * <-ch)
}

func main() {

	//ch := make(chan int)
	//fmt.Println("Hello from main")
	//go multiplyWithChannel(ch)
	//ch <- 9
	//fmt.Println("Bye from main")

	//	 close and ok syntax

	//ch := make(chan int)
	//elem, ok := <-ch
	//close(ch)
	//
	//fmt.Println("the element and ok staus of  the channel ", elem, ok)

	//	 using for loop with the channel

	//ch := make(chan string)
	//go inflateChannel(ch)
	//for {
	//	elem, ok := <-ch
	//	if ok == false {
	//		fmt.Println("Closing Channel", ok)
	//		break
	//
	//	}
	//	fmt.Println("Channel open ", elem, ok)
	//}

	//	buffered channels

	ch := make(chan string, 4)

	//fmt.Println("Length of the channel is ", len(ch))
	//fmt.Println("Capacity of the Channel is ", cap(ch))
	go func() {

		ch <- "hello "
		ch <- "hi"
		ch <- "hww"
		ch <- "heee"
		close(ch)
	}()

	for it := range ch {
		//fmt.Println("Length of the channel is ", len(ch))

		fmt.Println(it)
	}

	//	 for range loop

}
func inflateChannel(ch chan string) {
	for v := 0; v < 3; v++ {

		ch <- "hello world"

	}
	close(ch)

}
