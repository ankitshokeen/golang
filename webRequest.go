package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const myUrl string = "https://fakestoreapi.com/products/1"

// handling web requests
func webRequest() {
	fmt.Println("welcome to web request handling class")

	fmt.Println("web request class")

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("type of response data is %T \n", response)
	fmt.Println("response data is: ", response)

	resByteData, _ := ioutil.ReadAll(response.Body)
	fmt.Println("byte data to response is ", resByteData)
	fmt.Println(string(resByteData))
}
