package main

import "fmt"

func sayWichType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Printf("Тип int: %d\n", v)
	case string:
		fmt.Printf("Тип string: %s\n", v)
	case bool:
		fmt.Printf("Тип bool: %t\n", v)
	case chan int:
		fmt.Println("Тип: chan int")
	case chan string:
		fmt.Println("Тип: chan string")
	default:
		fmt.Println("Неизвестный тип")
	}
}
func main() {
	sayWichType(42)
	sayWichType("hello")
	sayWichType(true)
	sayWichType(make(chan int))
	sayWichType(make(chan string))
	sayWichType(3.14)
}
