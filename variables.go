package main

import (
	"fmt"
)

// const keywords are immutable their value cannot be modified ,if want to create public variable, name of the variable,s first letter should be in capital letter
const LoginToken string = "login_token"

func variables() {

	// print statements, placeholders, implicit type variables go decide itself hat type of variable is by given value
	var name = "ankit"
	var age = 24
	fmt.Println(name)
	fmt.Printf("my name is %v and age is: %v \n", name, age)
	fmt.Printf("my name type is %T and age type is: %T \n", name, age)

	// default values
	var id int
	var class string
	fmt.Printf("id is %v, class is %v \n", id, class)

	// no var declaration style, not var style declaration only inside method if want to create globally var keyword or var type is must
	token := 355577
	fmt.Println(token)
	// token = "244k67" if no var style is using you cannot change variable type by changing its value because what value given to it first it will be treated same later

	// using global variable
	fmt.Println(LoginToken)
	// LoginToken = "changed_value" // trying to modify the LoginToken value but it cannot be done, if we declare a variable as const its value cannot be modified

}
