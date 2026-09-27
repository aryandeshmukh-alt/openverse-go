package main

import (
	"fmt"
)

func main () {

	days := map[int]string {
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var index int
	fmt.Println("Enter the index of element you want to find: ")
	fmt.Scanf("%d", &index)

	day, found := days[index]

	if found {
		fmt.Printf("Day for index %d: %s", index, day)
	} else {
		fmt.Println("No day found for given index")
	}

}