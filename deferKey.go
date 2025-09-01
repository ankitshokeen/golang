package main

import "fmt"

func deferKey() {
	// defer keyword make that program execute at last first it will be storing in stack using last in first out principle
	defer fmt.Println("ankit")
	defer fmt.Println("is")
	defer fmt.Println("name")
	defer fmt.Println("my")
	fmt.Println("hello world, ")

	deferNum()
}

func deferNum() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

// here is expected output will be (hello world, 54321 my name is ankit
