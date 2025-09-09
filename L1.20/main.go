package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	b := []byte(s)
	reverse(b, 0, len(b)-1)
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}
	return string(b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) > 0 && text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}
	result := reverseWords(text)
	fmt.Println(result)
}
