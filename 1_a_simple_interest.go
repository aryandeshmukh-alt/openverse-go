package main

import (
	"fmt"
)

func main () {
	var p, r, t float64

	fmt.Println("Please enter principal amount: ")
	fmt.Scanf("%f", &p)
	fmt.Println("Please enter the interest rate: ")
	fmt.Scanf("%f", &r)
	fmt.Println("Please enter time in years: ")
	fmt.Scanf("%f", &t)

	si := (p*r*t)/100
	fmt.Printf("Simple Interest: %.2f", si)
}