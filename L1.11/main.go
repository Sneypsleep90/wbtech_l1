package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	m := make(map[int]bool)
	var result []int
	for _, val := range a {
		m[val] = true
	}
	for _, val := range b {
		if m[val] {
			result = append(result, val)
			delete(m, val)
		}
	}
	fmt.Println(result)
}
