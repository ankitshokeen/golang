package main

import (
	"fmt"
)

func maps() {
	fmt.Println("welcome to maps class")

	// creating map assigning with key value, map[key] = value
	languages := make(map[string]string)
	languages["JDK"] = "java"
	languages["CPP"] = "c++"
	languages["GO"] = "golang"
	languages["RB"] = "ruby"
	languages["JS"] = "node"

	fmt.Println(languages)

	// get value using key
	fmt.Println(languages["JDK"])

	// delete entity using key
	delete(languages, "JS")

	fmt.Println(languages)

	// iterate over map using for loop
	for key, value := range languages {
		fmt.Printf("key is %v, value of key is %v\n", key, value)
	}
}
