package main

import (
	"fmt"
	"math/rand"
	"time"
)

func switchCase() {
	fmt.Println("welcome to switchCase class")

	// this is used to get random number between range at current time
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	// getting output as per these cases if no one case matches it will return default value
	switch diceNumber {
	case 1:
		fmt.Println("dice number is 1 now you can open")
	case 2:
		fmt.Println("dice number is 2 now you can move 2 spot")
	case 3:
		fmt.Println("dice number is 3 now you can move 3 spot")
	case 4:
		fmt.Println("dice number is 4 now you can move 4 spot")
	case 5:
		fmt.Println("dice number is 5 now you can move 5 spot")
	case 6:
		fmt.Println("dice number is 6 now you can move 6 spot and re roll dice")

	default:
		fmt.Println("dice lost")
	}
}
