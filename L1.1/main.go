package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human
}

func (h Human) SayMyName() {
	fmt.Printf("My name is %s \n", h.Name)
}
func (h Human) SayMyAge() {
	fmt.Printf("My age is %v", h.Age)
}

func main() {
	action := Action{Human{
		Age:  19,
		Name: "Хайзенберг",
	}}
	action.SayMyName()
	action.SayMyAge()

}
