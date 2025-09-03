package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://ankit.dev"

// handling web requests
func webRequest() {
	fmt.Println("welcome to web request handling class")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is: %T \n", response)

	dataByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)

	fmt.Println(content)
	fmt.Println(dataByte)

	defer response.Body.Close()
}
