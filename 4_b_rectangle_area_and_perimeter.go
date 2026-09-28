package main

import (
	"fmt"
)

type Rectangle struct {
	length int
	width int
}

func (r Rectangle) Area () int {
	return r.length * r.width
}

func (r Rectangle) Perimeter () int {
	return 2 * (r.length + r.width)
}

func main () {
	var length int
	var width int

	fmt.Println("Enter length of the rectangle: ")
	fmt.Scanf("%d", &length)
	fmt.Println("Enter width of the rectangle: ")
	fmt.Scanf("%d", &width)

	rectangle := Rectangle{
		length: length,
		width: width,
	}

	fmt.Println("Area of the rectangle: ", rectangle.Area())
	fmt.Println("Perimeter of the rectangle: ", rectangle.Perimeter())
}