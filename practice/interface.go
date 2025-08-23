package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Man struct {
	name string
}
type Dog struct {
	Breed string
}

func (m Man) Speak() string {
	str := "I am a man" + m.name
	return str
}
func (d Dog) Speak() string {
	str := "I am a dog" + d.Breed
	return str
}

func main() {
	fmt.Println("Start of code ")
	var i Speaker
	i = Man{"Tanmaya"}
	fmt.Println(i.Speak())
	i = Dog{"Jackie"}
	fmt.Println(i.Speak())

}
