package main

import "fmt"

func main() {
	x := 2
	switch x {
	case 1:
		fmt.Println("Number 1 ")

	case 2:
		fmt.Println("Number 2 ")

	case 3:
		fmt.Println("Number 3 ")
	default:
		fmt.Println("Number default")
	}

}
