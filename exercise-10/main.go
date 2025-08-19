package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(401)
	fmt.Println(x)

	if x < 60 {
		fmt.Println("The steak is ALIVE")
	} else if x >= 60 && x <= 90 {
		fmt.Println("The steak is RARE")
	} else if x >= 90 && x < 120 {
		fmt.Println("The steak is MEDIUM RARE")
	} else if x >= 120 && x < 180 {
		fmt.Println("The steak is MEDIUM")
	} else if x >= 180 && x < 240 {
		fmt.Println("The steak is MEDIUM WELL")
	} else if x >= 240 && x < 320 {
		fmt.Println("The steak is WELL DONE")
	} else {
		fmt.Println("The steak is BURNT")
	}
}
