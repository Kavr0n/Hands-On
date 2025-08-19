package main

import (
	"fmt"
)

var xi = []int{42, 43, 44, 45, 46, 47}

func main() {
	for i, v := range xi {
		fmt.Printf("Iteration %d: X is %d\n", i, v)
	}
}
