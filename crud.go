package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GETrequest() {
	fmt.Println("GET request class")

	content, err := http.Get("https://jsonplaceholder.typicode.com/todos/5")

	if err != nil {
		fmt.Println("Incorrect URL")
		panic(err)
	}

	fmt.Println(content)

	defer content.Body.Close()

	// readURL, err := ioutil.ReadAll(content.Body)

	// if err != nil {
	// 	fmt.Println("Empty URL")
	// 	panic(err)
	// }

	// fmt.Println(string(readURL))

	var todo Todo
	err = json.NewDecoder(content.Body).Decode(&todo)
	if err != nil {
		panic(err)
	}
	fmt.Println(todo)
}

func POSTrequest() {
	fmt.Println("POST request class")

	todo := Todo{
		UserId:    1,
		Title:     "SDE",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("no content found")
		panic(err)
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	postURL := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(postURL, "application/json", jsonReader)

	if err != nil {
		fmt.Println("error 404")
		panic(err)
	}

	fmt.Println("POST status: ", res.Status)
	defer res.Body.Close()

	POSTres, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("not able to give response")
		panic(err)
	}

	fmt.Println("post request response: ", string(POSTres))

}

func UPDATErequest() {
	fmt.Println("UPDATE request class")

	todo := Todo{
		UserId:    14367,
		Title:     "SDE architect",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("error while Marshal to json")
		panic(err)
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	updateURL := "https://jsonplaceholder.typicode.com/todos/8"

	res, err := http.NewRequest(http.MethodPut, updateURL, jsonReader)

	if err != nil {
		fmt.Println("error creating update request")
		panic(err)
	}

	res.Header.Set("Content-type", "application/json")

	client := http.Client{}
	response, err := client.Do(res)

	if err != nil {
		fmt.Println("error sending update request")
	}

	fmt.Println(response.Status)

	defer response.Body.Close()

	updateRes, _ := ioutil.ReadAll(response.Body)

	fmt.Println("update response: ", string(updateRes))
}

func deleteReq() {
	fmt.Println("delete request class ")

	deleteURL := "https://jsonplaceholder.typicode.com/todos/2"

	res, _ := http.NewRequest(http.MethodDelete, deleteURL, nil)

	client := http.Client{}
	response, err := client.Do(res)

	if err != nil {
		fmt.Println("error sending delete request")
	}

	fmt.Println(response.Status)

	defer response.Body.Close()

	deleteRes, _ := ioutil.ReadAll(response.Body)

	fmt.Println("delete response: ", string(deleteRes))
}
