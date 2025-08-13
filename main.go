package main

import (
	"fmt"

	"github.com/Kavr0n/puppy"
)

func main() {
	// Simple bark
	bark := puppy.Bark()
	fmt.Println(bark)

	// Multiple barks
	barks := puppy.Barks()
	fmt.Println(barks)

	// Loud bark
	bigBark := puppy.BigBark()
	fmt.Println(bigBark)

	// Multiple loud barks
	bigBarks := puppy.BigBarks()
	fmt.Println(bigBarks)
}
