package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My favourite number is :", rand.Intn(10))

}

// so a package is basically helps in bundling multiple files together , and in each file we need to add the package name so that it point to that package
// package is of two types - 1) custom made - like the folders that we create   and 2) built in - already provided by the language , which we can directly use like math , fmt
