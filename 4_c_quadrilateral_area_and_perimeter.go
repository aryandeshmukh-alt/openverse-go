package main

import (
	"fmt"
)

type Quadrilateral interface {
	Area () int
	Perimeter () int
}	

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

type Square struct {
	side int
}

func (s Square) Area () int {
	return s.side * s.side
}

func (s Square) Perimeter () int {
	return 4 * s.side
}

func Print (shape Quadrilateral) {
	fmt.Println("Area of the shape: ", shape.Area())
	fmt.Println("Perimeter of the shape: ", shape.Perimeter())
}

func main () {
	var option int
	var length int
	var width int
	var side int

	fmt.Println("Choose a shape: ")
	fmt.Println("Enter 1 for rectangle: ")
	fmt.Println("Enter 2 for square: ")
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		fmt.Println("Enter length of rectangle: ")
		fmt.Scanf("%d", &length)
		fmt.Println("Enter width of rectangle: ")
		fmt.Scanf("%d", &width)

		rectangle := Rectangle {
			length: length,
			width: width,
		}

		Print(rectangle)
	case 2:
		fmt.Println("Enter side of square: ")
		fmt.Scanf("%d", &side)

		square := Square {
			side: side,
		}

		Print(square)
	}
}