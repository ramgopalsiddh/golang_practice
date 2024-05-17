package main

import (
	"encoding/json"
	"fmt"
)

// Struct/class 
type course struct {
	Name string `json:"coursename"` // Now index name is changed to Name to coursename
	Price int
	Platform string `json:"website"` // Rename as website
	Password string `json:"-"` // - this prevent to show this field to API consumer
	Tags []string `json:"tags,omitempty"`  // This use for stop show nil field
}

func main() {
	// welcome message
	fmt.Println("welcome in  encode json data practice")

	// call function
	EncodeJson()

	// call function
	DecodeJson()
}


func EncodeJson() {

	// Add Data
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "lco.in", "js123", []string{"web-dev", "js"}},
		{"Go Bootcamp", 199, "lco.in", "go123", []string{"backend", "go"}},
		{"MERN Bootcamp", 99, "lco.in", "mern123", nil },
	}

	// Package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // \t use for space

	// error handeling
	if err != nil {
		panic(err)
	}

	// print json data
	fmt.Printf("%s\n", finalJson)
}


func DecodeJson(){

	// Add Json data
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "lco.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourses course

	// check validation of JSon
	checkValid := json.Valid(jsonDataFromWeb)

	// If valid then perform these operation
	if checkValid {
		fmt.Println("JSON was Vaild")
		// Unmarshal data
		if err := json.Unmarshal(jsonDataFromWeb, &lcoCourses); err != nil {
			panic(err)
		}  // use & and reference data for extra validation

		fmt.Printf("%#v\n", lcoCourses)
		
	} else {
		fmt.Printf("JSON WAS NOT VALID\n")
	}

	// some case where you just want to add data to key value

	var myOnlineData map[string]interface{}
	if err := json.Unmarshal(jsonDataFromWeb, &myOnlineData); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", myOnlineData)

	// print data by use for loop
	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v )
	}

}