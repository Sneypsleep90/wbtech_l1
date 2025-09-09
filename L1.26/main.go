package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	m := make(map[rune]bool)
	str = strings.ToLower(str)
	for _, ch := range str {
		if m[ch] {
			return false
		}
		m[ch] = true
	}
	return true
}
func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}
