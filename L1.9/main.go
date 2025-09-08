package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 5, 6, 7, 6, 5}
	firstCh := make(chan int)
	secondChan := make(chan int)
	go func() {
		for _, i := range nums {
			firstCh <- i
		}
		close(firstCh)
	}()
	go func() {
		for _, i := range nums {
			secondChan <- i * 2
		}
		close(secondChan)
	}()
	for result := range secondChan {
		fmt.Println(result)
	}
}
