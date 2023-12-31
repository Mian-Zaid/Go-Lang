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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {

	jsonData := []byte(`
		{
			"coursename": "React JS Bootcamp",
			"price": 299,
			"website": "zaidcodes.com",       
			"tags": ["reactjs","react","bootcamp"]
		}
	`)

	//save as struct
	var myCourse course

	isValid := json.Valid(jsonData)
	if isValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("\nDecoded JSON is\n %#v\n", myCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//save as key value pairs
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("\nDecoded JSON is\n %#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("\nKey is %v, Value is %v, Type is %T\n", k, v, v)
	}

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
