package main

import (
	"fmt"
	"time"
)

func Worker(i int, ch <-chan int) {
	for c := range ch {
		fmt.Printf("%d:%d\n", i, c)
	}
}

func main() {
	var n int
	fmt.Println("укажите колво воркеров")
	fmt.Scan(&n)
	if n < 0 {
		return
	}
	ch := make(chan int)
	for i := 1; i <= n; i++ {
		go Worker(i, ch)
	}
	cnt := 0
	for {
		ch <- cnt
		cnt++
		time.Sleep(444 * time.Millisecond)
	}
}
