package main

import (
	"fmt"
)

type Message struct {
	name    string
	message string
}

func main() {

	input := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	alice := make(chan string)
	bob := make(chan string)
	output := make(chan Message)

	go func() {

		message := ""

		for _, ch := range input {

			if ch == '^' {
				break
			}

			if ch == '$' {
				alice <- message
				message = ""
			} else if ch == '&' {
				bob <- message
				message = ""
			} else {
				message += string(ch)
			}
		}

		close(alice)
		close(bob)
	}()

	go func() {

		aliceOpen := true
		bobOpen := true

		for aliceOpen || bobOpen {

			select {

			case message, ok := <-alice:
				if ok {
					output <- Message{
						name:    "alice",
						message: message,
					}
				} else {
					aliceOpen = false
				}

			case message, ok := <-bob:
				if ok {
					output <- Message{
						name:    "bob",
						message: message,
					}
				} else {
					bobOpen = false
				}
			}
		}

		close(output)
	}()

	for message := range output {
		fmt.Printf("%s : %s\n", message.name, message.message)
	}
}