package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("This is where initialization for my program happens")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(401)
	fmt.Println(x)

	switch {
	case x < 60:
		fmt.Printf("After %v seconds of cook the steak is ALIVE\n", x)
	case x >= 60 && x <= 90:
		fmt.Printf("After %v seconds of cook the steak is RARE\n", x)
	case x >= 90 && x < 120:
		fmt.Printf("After %v seconds of cook the steak is MEDIUM RARE\n", x)
	case x >= 120 && x < 180:
		fmt.Printf("After %v seconds of cook the steak is MEDIUM\n", x)
	case x >= 180 && x < 240:
		fmt.Printf("After %v seconds of cook the steak is MEDIUM WELL\n", x)
	case x >= 240 && x < 320:
		fmt.Printf("After %v seconds of cook the steak is WELL DONE\n", x)
	default:
		fmt.Printf("After %v seconds of cook the steak is BURNT\n", x)
	}
}
