package main

import (
	"fmt"
	"errors"
)

func accessSlice (slice []int, index int) error {
	if index > len(slice)-1 {
		return errors.New("Error: Trying to access to element out of index")
	}
	fmt.Println("Index: ", index, " & Value: ", slice[index])
	return nil
}

func main () {
	slice := []int{1, 2, 4, 6}
	var index int
	fmt.Println("Enter index: ")
	fmt.Scanf("%d", &index)
	err := accessSlice(slice, index)
	if err != nil {
		fmt.Println(err)
	}
}