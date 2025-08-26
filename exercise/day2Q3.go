package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Bank struct {
	bal int
	mu  sync.Mutex
}

func (b *Bank) deposit(amt int, wg *sync.WaitGroup) {
	defer wg.Done()
	b.mu.Lock()
	b.bal += amt
	fmt.Printf("Deposited: %d | Balance: %d\n", amt, b.bal)
	b.mu.Unlock()
}

func (b *Bank) withdraw(amt int, wg *sync.WaitGroup) {
	defer wg.Done()
	b.mu.Lock()
	if b.bal >= amt {
		b.bal -= amt
		fmt.Printf("Withdrew: %d | Balance: %d\n", amt, b.bal)
	} else {
		fmt.Printf("Failed withdrawal of %d | Balance: %d\n", amt, b.bal)
	}
	b.mu.Unlock()
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	b := Bank{bal: 500}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go b.deposit(rand.Intn(200), &wg)
		go b.withdraw(rand.Intn(200), &wg)
	}

	wg.Wait()
	fmt.Println("Final Balance:", b.bal)
}
