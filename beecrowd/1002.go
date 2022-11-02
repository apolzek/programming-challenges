package main

import (
	"fmt"
	"math"
)

func main() {

	const n = 3.14159
	var raio float64

	fmt.Scanf("%f", &raio)

	area := n * math.Pow(raio, 2)
	fmt.Printf("A=%.4f\n", area)
}
