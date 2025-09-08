package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	var result []string
	m := map[string]struct{}{}

	for _, word := range words {
		m[word] = struct{}{}
	}
	for val := range m {
		result = append(result, val)
	}
	fmt.Println(result)

}
