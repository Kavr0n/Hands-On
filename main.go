package main

import (
	"fmt"

	"github.com/Kavr0n/puppy"
)

func main() {
	// Simple bark
	bark := puppy.Bark()
	fmt.Println("Bark:", bark)

	// Multiple barks
	barks := puppy.Barks()
	fmt.Println("Barks:", barks)

	// Loud bark
	bigBark := puppy.BigBark()
	fmt.Println("Big Bark:", bigBark)

	// Multiple loud barks
	bigBarks := puppy.BigBarks()
	fmt.Println("Big Barks:", bigBarks)
}
