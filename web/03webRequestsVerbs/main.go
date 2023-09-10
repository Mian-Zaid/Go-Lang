package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Request VERBS in Golang!")
	PerformGetRequest()
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

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
