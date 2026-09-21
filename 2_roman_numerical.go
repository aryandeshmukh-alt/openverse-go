package main

import (
	"fmt"
)

func main () {
	var roman string
	fmt.Println("Enter the roman numerical: ")
	fmt.Scanf("%s", &roman)

	values := map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}

	total := 0

	for i:=0; i<len(roman); i++ {
		current := values[roman[i]]
		if i+1<len(roman) && current < values[roman[i+1]] {
			total -= current
		} else {
			total += current
		}
	}

	fmt.Println(total)
}