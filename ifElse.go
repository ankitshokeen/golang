package main

import "fmt"

func ifElse() {
	fmt.Println("class of if else")

	var num int = 13

	if num < 2 {
		fmt.Println("invalid number")
	} else if num%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}
