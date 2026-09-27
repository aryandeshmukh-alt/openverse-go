package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main () {

	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the sentence whose words' frequency you want to count: ")
	input, _ = reader.ReadString('\n')

	input = strings.TrimSpace(input)
	words := strings.Fields(input)
	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}
	maxFrequency := 0
	for _, word := range words {
		if count[word] > maxFrequency {
			maxFrequency = count[word]
		}
	}

	result := []string{}
	exists := make(map[string]bool)
	for _, word := range words {
		if count[word]==maxFrequency && !exists[word] {
			result = append(result, word)
			exists[word] = true
		}
	}

	fmt.Println("Most frequent words: ", result)
}