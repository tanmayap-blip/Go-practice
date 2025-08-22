package main

import "fmt"

func updateValue(num *int) {
	*num = *num + 10
}

func main() {
	x := 10
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 20
	fmt.Println(x)
	fmt.Println(*p)

	nums := 15
	fmt.Println(nums)
	updateValue(&nums)
	fmt.Println(nums)
}
