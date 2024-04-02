package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//  print Welcome message
	fmt.Println(" Welcome in Files ")

	// content for insert in file
	content := "This is my portfolio website where you get know more about me please visit https://www.ramgopal.dev"

	//create file
	file, err := os.Create("./portfolio.txt")
	checkNilErr(err)

	// Write coontent in file
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	// calculate length of file
	fmt.Println("Length is: ", length)
	defer file.Close()

	// call the read file function
	readFile("./portfolio.txt")
}


// function for Read files
func readFile(filename string) {
	databytes, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Files data:- \n", string(databytes))
}


// Create Function for check error
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}