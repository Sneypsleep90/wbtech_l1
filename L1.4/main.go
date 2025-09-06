package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Worker(i int, ch <-chan int, done chan struct{}) {
	for {
		select {
		case c := <-ch:
			fmt.Printf("%d:%d\n", i, c)
		case <-done:
			fmt.Println("завершаем работу")
			return
		}
	}
}

func main() {
	var n int
	done := make(chan struct{})
	fmt.Println("укажите колво воркеров")
	fmt.Scan(&n)
	if n < 0 {
		return
	}
	ch := make(chan int)
	for i := 1; i <= n; i++ {
		go Worker(i, ch, done)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		count := 1
		for {
			select {
			case <-done:
				close(ch)
				return
			default:
				ch <- count
				count++
				time.Sleep(544 * time.Millisecond)
			}
		}
	}()
	<-c
	fmt.Println("получен сигнал,стопаемся!!!")
	close(done)

	time.Sleep(2 * time.Second)
	fmt.Println("!!!стоп работы!!!")

}
