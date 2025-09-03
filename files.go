package main

import (
	"fmt"
	"os"
)

func files() {
	fmt.Println("welcome to files class")

	content := "files class in golang"

	// in os package it has create method which is used to create new file os.Create("address")
	file, err := os.Create("./testFile.txt")

	if err != nil {
		// panic keyword is used to throw an error and print the error
		panic(err)
	}

	// in io package has WriteString() which is used to write content into file in string format
	length, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("file has length: ", length)

	defer file.Close()

	readFile("./testFile.txt")
}

func readFile(filename string) {
	// ReadFile is used to read file it returns the data in bytecode
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// if want to read real data you have to parse it into string format using string()
	fmt.Println(string(databyte))
}
