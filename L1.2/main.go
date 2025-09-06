package main

import (
	"fmt"
	"sync"
)

func main() {
	var nums = []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var wg sync.WaitGroup
	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			ch <- num * num
			defer wg.Done()

		}(n)

	}
	wg.Wait()

	for res := range ch {
		fmt.Println(res)
	}

}
