package main

import (
	"fmt"
	url2 "net/url"
)

const myGitUrl string = "https://github.com/ankitshokeen"
const classSourceUrl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=27"

func urlS() {
	fmt.Println("welcome to class of handling urlS")

	fmt.Println(myGitUrl)

	// this will parse string into url
	result, err := url2.Parse(classSourceUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qParams := result.Query()
	fmt.Printf("type of query params is %T\n", qParams)

	fmt.Println(qParams["index"])

	for _, v := range qParams {
		fmt.Println("param is : ", v)
	}

	// these are parts to generate a url
	partsOfUrl := &url2.URL{
		Scheme:  "https",
		Host:    "github.com",
		Path:    "/ankitshokeen/golang",
		RawPath: "user=ankitshokeen",
	}

	// here it will merge all parts of string into single string
	compUrl := partsOfUrl.String()

	fmt.Println(compUrl)
}
