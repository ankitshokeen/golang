package main

import (
	"fmt"
	"sort"
)

func slices() {
	fmt.Println("Welcome to the class of slices")

	var fruits = []string{"apple", "peach", "pear"}
	fmt.Printf("The type of fruits is %T \n: ", fruits)

	// adding more values to slices
	fruits = append(fruits, "banana", "orange")
	fmt.Println(fruits)

	// deleting/slicing values from range
	fruits = append(fruits[1:2])
	fmt.Println(fruits)

	hignScores := make([]int, 5)

	hignScores[0] = 999
	hignScores[1] = 200
	hignScores[2] = 777
	hignScores[3] = 400
	hignScores[4] = 555
	fmt.Println(hignScores)

	// slices is already filled by values if we give it more values using indexing format it will throw error
	// highScores[5] = 353
	// highScores[6] = 463

	// if want to add more value to slices use append method
	hignScores = append(hignScores, 875, 853)
	fmt.Println(hignScores)

	// this will give boolean value is slices are sorted
	fmt.Println(sort.IntsAreSorted(hignScores))

	// sort slices
	sort.Ints(hignScores)
	fmt.Println(hignScores)

	// now this will give true value
	fmt.Println(sort.IntsAreSorted(hignScores))

	var cources = []string{"java", "spring boot", "docker", "linux", "go"}
	fmt.Println(cources)

	// deleting value by index using append method
	cources = append(cources[:3], cources[4:]...)
	fmt.Println(cources)
}
