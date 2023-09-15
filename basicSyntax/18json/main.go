package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in Golang!")
	EncodeJson()
}

// encode json
func EncodeJson() {
	myCourses := []course{
		{
			"React JS Bootcamp",
			299,
			"zaidcodes.com",
			"abc123",
			[]string{"reactjs", "react", "bootcamp"},
		},
		{
			"Go lang Bootcamp",
			200,
			"zaidcodes.com",
			"abc123",
			[]string{"go", "Golang", "bootcamp"},
		},
		{
			"Full stack Bootcamp",
			500,
			"zaidcodes.com",
			"abc123",
			nil,
		},
	}

	//package this data as JSON

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	handleErr(err)

	fmt.Printf("%s\n", string(finalJson))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
