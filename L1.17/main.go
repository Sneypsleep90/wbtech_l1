package main

import "fmt"

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7}
	for i := 1; i <= len(array)-1; i++ {
		fmt.Println(array[i])
	}
}
