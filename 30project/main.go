package main

import "fmt"

func main() {
	fmt.Println("Please Enter your name")

	var name string
	fmt.Scanln(&name) // & use for memory location address refrence
	fmt.Printf("Hello,%v welcome to quize game!", name)
}