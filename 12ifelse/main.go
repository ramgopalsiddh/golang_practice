package main

import (
	"errors"
	"fmt"
)

// Function for print test error
func errorFunction() error {
	return errors.New("You have a test error")
}

func main() {
	// Welcome message 
	fmt.Println("Welcome in If Else loop intro")

	// if else loop example
	loginCount := 24
	var result string 

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount >10 {
		result = "Watch out user"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)


	// direct give conditon and check condition
	if 6%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}


	// Assigin value and check same time in loop
	if num := 5; num < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is Not less then 10")
	}

	
	// comman use this in error print 
	// we print a test error here 
	var err error
	err = errorFunction()

	if err != nil {
		fmt.Println("Error:", err)
	}
}