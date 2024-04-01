package main

import "fmt"

func main() {
	fmt.Println("Welcome in pointers")

	// var ptr *int 
	// fmt. Println("Default value of pointer ptr is", ptr)

	myNumber := 23

	// reference means &
	var ptr = &myNumber

	fmt.Println("Vaule of my pointer is ", ptr) // This print memory address like 0xc0000120b8
	fmt.Println("Vaule of my pointer is ", *ptr) // This print vaule of pointer like 23

	*ptr = *ptr * 2
	fmt.Println("vaule after multiply by 2 is", myNumber)

}
