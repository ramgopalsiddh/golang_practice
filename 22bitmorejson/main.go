package main

import (
	"encoding/json"
	"fmt"
)

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