package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {

		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("Iteration %d: X is zero\n", i)
		case 1:
			fmt.Printf("Iteration %d: X is one\n", i)
		case 2:
			fmt.Printf("Iteration %d: X is two\n", i)
		case 3:
			fmt.Printf("Iteration %d: X is three\n", i)
		case 4:
			fmt.Printf("Iteration %d: X is four\n", i)
		default:
			fmt.Printf("Iteration %d: X is out of range\n", i)
		}
	}
}
