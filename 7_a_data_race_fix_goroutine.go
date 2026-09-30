package main

import (
	"fmt"
	"time"
)

func isEven (n int) bool {
	return n%2 == 0
}

func main () {

	n := 3
	done := make(chan bool)

	go func () {
		
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)

		if nIsEven {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}

		done <- true
	} ()

	<- done
	n++ 
}