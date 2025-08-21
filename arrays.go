package main

import "fmt"

func arr() {
	fmt.Println("Welcome to the class of Arrays")

	// creating array of size [4]
	var fruits [4]string

	// initializing values at positions [1], [2], [3], [4]
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fruits[3] = "Grape"

	// printing array table
	fmt.Println(fruits)

	// getting length of fruits array
	fmt.Println(len(fruits))

	// initializing values while creating the array
	var veges = [4]string{"potato", "chilli", "onion", "garlic"}

	fmt.Println(veges)
}
