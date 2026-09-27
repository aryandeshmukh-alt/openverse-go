package main

import (
	"fmt"
	"os"
	"bufio"
)

func main () {

	arr := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	var index1, index2 int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter index1 and index2 for making slices: ")
	fmt.Fscan(reader, &index1, &index2)

	if index1<0 || index2>len(arr) || index1>index2 {
		fmt.Println("Incorrect indexes")
		return
	}

	slice1 := arr[: index1]
	slice2 := arr[index1: index2]
	slice3 := arr[index2: ]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}