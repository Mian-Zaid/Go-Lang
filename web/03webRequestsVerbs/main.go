package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Request VERBS in Golang!")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:8000/get"

	response, err := http.Get(url)
	handleErr(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content is: ", string(content))

	var responseString strings.Builder
	_, _ = responseString.Write(content)
	fmt.Println("Content is: ", responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"zaid ali",
			"country":"pakistan",
			"number":923121686494
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	handleErr(err)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	handleErr(err)
	fmt.Println(string(dataBytes))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
