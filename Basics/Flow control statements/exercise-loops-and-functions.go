package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var result float64 = x
	for (x - result * result) * (x - result * result) > 1e-6 {
		result -= (result * result - x) / (2 * result)
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
}

