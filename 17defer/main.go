package main

import "fmt"

func main() {

	// Wel come message
	fmt.Println("Welcome in Defer")

	// defer execuite after surrounding statememt execuite
	defer fmt.Println("World")
	// Defer follow LIFO (list in first out) 
	// print output like :-  hello, Two, One, World
	defer fmt.Println("One")
	defer fmt.Println("Two")

	fmt.Println("hello!")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}