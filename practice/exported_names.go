package main

import "fmt"

// Exported (public)

func SayHello(name string) {
	fmt.Println("Hello,", name)
}

// Unexported (private)

func whisperSecret() {
	fmt.Println("psst... this is a secret")
}
