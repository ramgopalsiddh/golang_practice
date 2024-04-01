package main

import (
	"fmt"
)

// LoginToken's L is capital beacuse this is a public variable
// that access by any other file OR program in same folder
const LoginToken string = "logintokentest"  //Public

func main() {
	fmt.Println("Variables")

	var username string = "Ram gopal"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.5735459565
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "ramgopal.dev"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style 
	// this method can't allowed out side of main function
	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	// Access Public variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
