package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		i := 0
		for {
			ch <- i
			fmt.Printf("Отпраяем данные:%d\n", i)
			i++
			time.Sleep(545 * time.Millisecond)

		}
	}()
	timeout := time.After(time.Duration(5) * time.Second)
	for {
		select {
		case num := <-ch:
			fmt.Println("получаем данные", num)
		case <-timeout:
			fmt.Println("остановка программы")
			return
		}

	}
}
