package main

import "fmt"

func main() {
	slice := []int{1, 23, 45, 43, 452, 245}
	fmt.Println(slice)
	i := 2
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
