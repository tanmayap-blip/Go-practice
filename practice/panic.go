package main

import (
	"fmt"
)

// This function will cause a panic if the input is zero.
func safeDivision(numerator, denominator int) int {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("A panic was recovered! Reason: %v\n", r)
		}
	}()

	if denominator == 0 {
		panic("cannot divide by zero")
	}
	return numerator / denominator
}

func main() {
	fmt.Println("Attempting to divide 10 by 2...")
	result1 := safeDivision(10, 2)
	fmt.Printf("Result is: %d\n\n", result1)

	fmt.Println("Attempting to divide 10 by 0...")
	result2 := safeDivision(10, 0) // This will cause a panic
	fmt.Printf("This line is never reached. Result is: %d\n", result2)

	fmt.Println("Attempting to divide 10 by 3...")
	result3 := safeDivision(10, 3)
	fmt.Printf("Result is: %d\n\n", result3)
}
