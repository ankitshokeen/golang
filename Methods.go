package main

import "fmt"

// no inheritance in golang; super keyword or parent class
func methods() {
	fmt.Println("welcome to the class of structs")

	// assigning values to struct
	usr1 := Usr{"ankit", 24, "ankitshokeen.work@gmail.com", true}

	fmt.Println(usr1)

	// print structure with entries
	fmt.Printf("Entry of User %+v \n", usr1)

	// get desired value
	fmt.Printf("name of user1 %v, email of user1 %v \n", usr1.name, usr1.email)

	// using getEmail() method here
	usr1.getEmail()

	// it is just creating copy of struct obj and modifying the copy
	usr1.newEmail()
}

type Usr struct {
	name   string
	age    int
	email  string
	status bool
}

// creating new method
func (u Usr) getEmail() {
	fmt.Println("user Email is: ", u.email)
}

// updating the copy value
func (u Usr) newEmail() {
	u.email = "ankit@go.dev"
	fmt.Println(u.email)
}
