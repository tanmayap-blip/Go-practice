package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, sum is ", add(1, 2))
}
