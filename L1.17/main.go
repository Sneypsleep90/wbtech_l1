package main

import "fmt"

func main() {
	var array = []int{1, 4, 5, 5, 643, 3, 53, 2, 4, 52, 4, 5, 6, 7, 8, 8, 9, 11, 2, 33, 45, 65, 47, 44, 27, 24, 24, 24}
	sortedArray := quickSort(array)
	fmt.Println(binarySearch(sortedArray, 643))
	fmt.Println(binarySearch(sortedArray, 3424324))

}

func binarySearch(slice []int, target int) int {
	low := 0
	high := len(slice) - 1
	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			return mid
		}
		if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1

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
