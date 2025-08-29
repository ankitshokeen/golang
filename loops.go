package main

import "fmt"

func loop() {
	fmt.Println("lets start loop class")

	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	fmt.Println(days)

	// for loop
	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	// for each loop
	//for i := range days {
	//	fmt.Println(days[i])
	//}

	// enhanced for each loop
	//for i, day := range days {
	//	fmt.Printf("day number is %v, day name is %v \n", i+1, day)
	//}

	// for/while loop
	//t := 5
	//c := 1
	//
	//for c <= 10 {
	//	fmt.Println(t * c)
	//	c++
	//}

}
