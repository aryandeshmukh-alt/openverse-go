package main

import (
	"fmt"
	"runtime"
)

func main() {

	var input string
	fmt.Println("Enter the string which you want to reverse:")
	fmt.Scan(&input)
	done := make(chan bool)

	go func() {

		runes := []rune(input)
		left := 0
		right := len(runes) - 1

		for left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}

		fmt.Println(string(runes), runtime.NumGoroutine())
		done <- true

	}()

	<-done
}