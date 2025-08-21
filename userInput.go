package main

import (
	"bufio"
	"fmt"
	"os"
)

func userInput() {

	var welcome = "welcome to user input"
	fmt.Println(welcome)

	// creating reader variable
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for my learning skills")

	// taking input form user
	rating, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating: ", rating)
	fmt.Printf("type of rating is %T \n", rating)
}
