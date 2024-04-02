package main

import "fmt"

// func main() {}
// Not allowed to write function inside function
// main function is entry point of program
func main() {
	
	// Welcom message
	fmt.Println("Welcome in Functions")

	// call the greeting function
	greeting()

	// call the adder function
	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	// call the proAdder function
	proRes := proAdder(2, 6, 8, 9)
	fmt.Println("Value of proresult is: ", proRes)
}


// define adder function for add 2 number
func adder(val1 int, val2 int) int {
	return val1 + val2
}


// define proAdder function for add multiple value 
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}


// define greeting function
func greeting() {
	fmt.Println("Hello i am ram gopal")
}