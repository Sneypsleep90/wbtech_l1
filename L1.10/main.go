package main

import (
	"fmt"
	"math"
)

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, t := range temp {
		key := int(math.Floor(t/10) * 10)
		groups[key] = append(groups[key], t)
	}

	for k, vals := range groups {
		fmt.Printf("%d: %v\n", k, vals)
	}
}
