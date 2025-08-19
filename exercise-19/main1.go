package main

import (
	"fmt"
)

func main() {

	sl := []string{"Almond cafe", "Bistro", "Cafe Mocha", "Deli", "Espresso Bar", "French Bakery"}

	fmt.Println(sl)
	fmt.Println("Length:", len(sl))
	for i, v := range sl {
		fmt.Println("Index:", i, "Value:", v)
	}
}
