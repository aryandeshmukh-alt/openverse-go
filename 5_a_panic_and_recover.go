package main

import (
	"fmt"
)

func accessSlice (slice []int, index int) {
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("Internal error: ", r)
		}
	}()
	fmt.Println("Item: ", index, " & Value: ", slice[index])
}

func main () {
	slice := []int{1, 2, 4, 6}
	var index int
	fmt.Println("Enter index: ")
	fmt.Scanf("%d", &index)
	accessSlice(slice, index)
	fmt.Println("Testing panic and recover")
}