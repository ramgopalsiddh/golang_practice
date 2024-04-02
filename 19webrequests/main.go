package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//  url for get responce
const url = "https://ramgopal.dev"

func main() {

	// Welcome message
	fmt.Println("Welcome in Handle web requests")

	// Get responce 
	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// Print responce type
	fmt.Printf("Type of responce: %T\n", responce)
	// close request
	defer responce.Body.Close() // caller's responsibility to close the connection

	// Get read all responce body of website
	databytes, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	// Print content
	content := string(databytes)
	fmt.Println(content)
}
