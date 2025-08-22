package main

import "fmt"

func main() {
	counter := 0

	increment := func() int {
		counter++
		return counter
	}

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
}
