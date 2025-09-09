package main

import "fmt"

func main() {
	array := []int{1, 4, 5, 64, 4, 5, 6, 4, 3, 3, 4, 5, 6}
	fmt.Print(quickSort(array))

}

func quickSort(array []int) []int {
	var greater, less []int
	if len(array) <= 1 {
		return array
	}
	pivot := array[0]
	for i := 1; i <= len(array)-1; i++ {
		if array[i] < pivot {
			less = append(less, array[i])
		} else {
			greater = append(greater, array[i])
		}
	}
	return append(append(quickSort(less), pivot), quickSort(greater)...)

}
