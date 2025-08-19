package main

import (
	"fmt"
)

var message string = "Hello, World!"

const greeting = "Welcome to Go programming!"

func main() {
	s1 := "its a simple Go program."
	s2 := "that is easy to understand."

	fmt.Println(message, s1)
	fmt.Println(greeting, s2)
}
