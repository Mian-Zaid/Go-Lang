package main

import (
	"fmt"
	"net/url"
)

const (
	myUrl = "https://lco.dev:9000/learn/?coursename=golangtuts&userid=12lji2"
)

func main() {
	fmt.Println("Welcome to handling URLs in Go Lang!")
	//parsing the url
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)
	// fmt.Println("Result: ", result.Scheme)
	// fmt.Println("Result: ", result.Host)
	// fmt.Println("Result: ", result.Path)
	fmt.Println("Result: ", result.RawQuery)
	// fmt.Println("Result: ", result.Port())

	qparams := result.Query()
	// fmt.Println("Query Params : ", qparams)

	for key, value := range qparams {
		fmt.Println("Param is: ", key, ":", value)
	}

	//constructing the URL

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=zaidali",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url: ", anotherUrl)

}
