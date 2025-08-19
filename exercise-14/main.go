package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {

		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("x is %d\n", x)
		fmt.Printf("y is %d\n", y)

		if x+y < 4 {
			fmt.Println("X + Y is less than 4")
		} else if x+y > 6 {
			fmt.Println("X + Y is greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("X is between 4 and 6")
		} else if y != 5 {
			fmt.Println("Y is not equal to 5")
		} else {
			fmt.Println("none of the above is true")
		}
	}
}
