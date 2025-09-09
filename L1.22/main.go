package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)
	a.SetString("2000000", 10)
	b.SetString("3000000", 10)
	sum := new(big.Int).Add(a, b)
	minus := new(big.Int).Sub(a, b)
	umnozhit := new(big.Int).Mul(a, b)
	razdelit := new(big.Int).Div(a, b)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("a+b: ", sum)
	fmt.Println("a-b: ", minus)
	fmt.Println("a*b: ", umnozhit)
	fmt.Println("a / b", razdelit)
}
