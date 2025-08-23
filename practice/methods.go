package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) greet() {
	fmt.Println("Hello, this is " + h.Name)
}
func (h *Human) birthday() {
	h.Age++
}

func main() {
	fmt.Println("Start of the code ")
	h1 := Human{"James", 20}
	h2 := Human{"Bond", 30}
	h1.greet()
	h2.birthday()
	fmt.Println(h1)
	fmt.Println(h2)
}
