package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "tanmaya", age: 22}
	p2 := Person{"jack", 50}
	var p3 Person

	fmt.Println(p1, p2, p3)

}
