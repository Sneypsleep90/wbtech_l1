package main

import (
	"fmt"
	"sync"
)

type Number struct {
	num int
	mu  sync.Mutex
}

func (n *Number) Increment() {
	n.mu.Lock()
	n.num++
	n.mu.Unlock()
}

func (n *Number) Number() int {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.num
}

func main() {
	var wg sync.WaitGroup
	n := Number{}

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.Increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Итог: %d", n.Number())

}
