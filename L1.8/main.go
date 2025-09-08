package main

import "fmt"

func changeBit(num int64, i uint, bit uint) int64 {
	if bit == 1 {
		num = num | (1 << i)
	} else {
		num = num &^ (1 << i)
	}
	return num

}

func main() {
	var num int64 = 5
	changeBit(num, 1, 0)
	fmt.Println(num)

}
