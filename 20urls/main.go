package main

import (
	"fmt"
	"net/url"
)

// Url for operation
const myurl string = "https://ramgopal.dev:3000/about?certificates=courses&github=ramgopal"

func main() {

	// welcome message
	fmt.Println("Welcome in Url handeling demo")

	// Print Url
	fmt.Println(myurl)

	// Url Parse
	result, _ := url.Parse(myurl)

	// extract url and get data
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// identify type of Query params
	qparams := result.Query()
	fmt.Printf("The Type of query params are: %T\n", qparams)

	//print a specific params by use index
	fmt.Println(qparams["certificates"])

	// print all params useing for loop
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}


	// Construct custom url 
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "ramgopal.dev",
		Path: "/about",
		RawQuery: "certificates=courses&github=ramgopal",
	}

	anotherurl := partsOfUrl.String()
	// print custom url
	fmt.Println(anotherurl)

}