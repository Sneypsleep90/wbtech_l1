package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	done := make(chan struct{})
	go func() {
		timer := time.NewTimer(d)
		<-timer.C
		close(done)
	}()
	<-done
}

func main() {
	fmt.Println("начало")
	sleep(2 * time.Second)
	fmt.Println("конец суеты")
}
