package main

import (
	"fmt"
	"time"
)

func goRoutine() {
	fmt.Println("welcome to the class of goRoutine")

	go sayHello()
	go sayHi()
	time.Sleep(7000 * time.Millisecond)
	fmt.Println("wake up after 7 seconds")

}

func sayHello() {
	fmt.Println("hello everyone!")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("wake up after 2 seconds")
}

func sayHi() {
	fmt.Println("hi coder!")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("wake up after 5 seconds")
}
