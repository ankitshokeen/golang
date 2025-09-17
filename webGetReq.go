package main

import (
	"fmt"
	"net/http"
)

func getReq() {
	fmt.Println("welcome to web get request class")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8090/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("response status : ", response.StatusCode)
	fmt.Println("content length is : ", response.ContentLength)

	// content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(content)
	// fmt.Println(string(content))

	// var responseString strings.Builder
	// byteLength, _ := responseString.Write(content)
	// fmt.Println("byte length is ", byteLength)
}
