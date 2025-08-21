package main

import "fmt"

func myPointer() {
	fmt.Println("Welcome to the class of pointers")

	// variable is term a value is stored at a particular address
	// pointer is used to get address of where variable is stored

	// creating integer variable
	var a int
	fmt.Println("value of variable a: ", a)

	// creating integer pointer
	var b *int
	fmt.Println("value of pointer b: ", b)

	// creating variable c and one pointer using & d referring to c variable stores value we assigned and pointer referring to address of variable c
	c := 25
	d := &c
	fmt.Println("value of variable c: ", c)
	fmt.Println("value of variable d: ", d)

	// here we have a pointer named d which is referring to c variable address, doing and operation to multiply c value by using pointer
	*d = *d * 2
	fmt.Println("value of variable c: ", c)
}
