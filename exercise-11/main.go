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

	switch {
	case x < 60:
		fmt.Println("The steak is ALIVE")
	case x >= 60 && x <= 90:
		fmt.Println("The steak is RARE")
	case x >= 90 && x < 120:
		fmt.Println("The steak is MEDIUM RARE")
	case x >= 120 && x < 180:
		fmt.Println("The steak is MEDIUM")
	case x >= 180 && x < 240:
		fmt.Println("The steak is MEDIUM WELL")
	case x >= 240 && x < 320:
		fmt.Println("The steak is WELL DONE")
	default:
		fmt.Println("The steak is BURNT")
	}
}
