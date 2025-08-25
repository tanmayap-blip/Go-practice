package main

import (
	"fmt"
	"time"
)

func doSomethingElse() {
	for i := 0; i < 5; i++ {
		fmt.Println("this is the go routine  ")
	}
}

func doSomething() {
	for i := 0; i < 5; i++ {
		fmt.Println("this is the normal ")
	}
}

func main() {
	go doSomethingElse() // this is basically a new go routine , other than main which runs in background
	time.Sleep(time.Second)
	doSomething()

	//anonymous  go routines

	//fmt.Println("Start of the program ")
	//go func() {
	//	fmt.Println("hi is this is anonymous")
	//}()
	//time.Sleep(time.Second)

}
