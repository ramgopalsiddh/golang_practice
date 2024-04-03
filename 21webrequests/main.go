package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// welcome message
	fmt.Println("Task for create get request in golang")

	// call the function for Get request
	PerformGetRequest()

	// Call the function for Post request
	PerformPostJsonRequest()

	// call the function for Post Form request
	PerformPostFormRequest()
}

// create function for get request
func PerformGetRequest() {

	// Define url
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	// for close request
	defer response.Body.Close()

	// for get Status code & content length
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("content length is: ", response.ContentLength)

	// Get content and data from url by use Strings package
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	// Print responce
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// simple way to get and print data
	//content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))
}

func PerformPostJsonRequest() {

	// Url
	const myurl = "http://localhost:8000/post"

	//Add Fake json data
	requestBody := strings.NewReader(`
		{
			"coursename":"Basic of golang",
			"price": 0,
			"platform":"ramgopal.dev"
		}
	`)

	// Post request for send data
	response, err := http.Post(myurl, "application/json", requestBody)

	// Handel error
	if err != nil {
		panic(err)
	}

	// Close connection after complete request
	defer response.Body.Close()

	// Display content
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest(){

	// Url
	const myurl = "http://localhost:8000/postform"

	// Form data
	data := url.Values{}
	data.Add("firstname", "Ram Gopal")
	data.Add("lastname", "Siddh")
	data.Add("email", "ramgopal@gmail.com")

	// Post Form request
	response, err := http.PostForm(myurl, data)
	// get error if any
	if err != nil {
		panic(err)
	}

	// Close after request
	defer response.Body.Close()

	// Print responce
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}