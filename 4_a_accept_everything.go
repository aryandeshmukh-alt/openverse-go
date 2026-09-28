package main

import (
	"fmt"
)

type Hello string

func AcceptAnything(value interface{}) {
	switch value := value.(type) {
	case int:
		fmt.Println("This is a value of type Integer,", value)
	case string:
		fmt.Println("This is a value of type String,", value)
	case bool:
		fmt.Println("This is a value of type Boolean,", value)
	case Hello:
		fmt.Println("This is a value of type Hello,", value)
	}
}

func main() {
	var option int
	fmt.Println("Enter your option:")
	fmt.Scan(&option)
	switch option {
	case 1:
		value := 100
		AcceptAnything(value)
	case 2:
		value := "Hello World"
		AcceptAnything(value)
	case 3:
		value := true
		AcceptAnything(value)
	case 4:
		value := Hello("Hello")
		AcceptAnything(value)
	}
}