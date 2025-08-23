package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["tanmay"] = 22
	fmt.Println(m1)

	m2 := map[string]int{
		"tanmaypanigrahi": 22,
	}
	fmt.Println(m2)

	m3 := map[string]int{}
	m3["hello"] = 1
	fmt.Println(m3)

	for key, value := range m1 {
		fmt.Println(key, value)
	}

}
