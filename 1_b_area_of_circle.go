package main

import (
	"fmt"
	"math"
)

func main () {
	var r float64

	fmt.Println("Enter radius of the circle: ")
	fmt.Scanf("%f", &r)

	a := math.Pi * math.Pow(r, 2)
	fmt.Printf("Area of circle: %.2f", a)
}