package main

import (
	"fmt"
	"strings"
	"sync"
)

func count(s string, ch chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()
	m := make(map[rune]int)
	for _, c := range strings.ToLower(s) {
		if c >= 'a' && c <= 'z' {
			m[c]++
		}
	}
	ch <- m
}

func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	ch := make(chan map[rune]int, len(words))
	var wg sync.WaitGroup

	for _, w := range words {
		wg.Add(1)
		go count(w, ch, &wg)
	}

	wg.Wait()
	close(ch)

	res := make(map[rune]int)
	for m := range ch {
		for k, v := range m {
			res[k] += v
		}
	}

	for k, v := range res {
		fmt.Printf("%c: %d\n", k, v)
	}
}
