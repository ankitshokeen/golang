package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name   string `json"name"`
	ID     int    `json"id_no"`
	Age    int    `json"age"`
	Target bool   `json"target_achieved"`
}

func jsonData() {
	fmt.Println("welcome to class of JSON")

	emp1 := Employee{"jhon", 118801, 28, true}
	fmt.Println("Employee data is: ", emp1)

	// encoding to json format
	jsonFormat, err := json.Marshal(emp1)

	if err != nil {
		fmt.Println("nothing to display")
		panic(err)
	}

	fmt.Println("Employee data in json format ", string(jsonFormat))

	// decoded from json format
	var empDec Employee
	json.Unmarshal(jsonFormat, &empDec)

	fmt.Println("Decoded data from json: ", empDec)
}
