package main

import "fmt"

// no inheritance in golang; super keyword or parent class
func structs() {
	fmt.Println("welcome to the class of structs")

	// assigning values to struct
	user1 := User{"ankit", 24, "ankitshokeen.work@gmail.com", true}

	fmt.Println(user1)

	// print structure with entries
	fmt.Printf("Entry of User %+v \n", user1)

	// get desired value
	fmt.Printf("name of user1 %v, email of user1 %v \n", user1.name, user1.email)
}

type User struct {
	name   string
	age    int
	email  string
	status bool
}
