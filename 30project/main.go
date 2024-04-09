package main

import "fmt"

func main() {
	fmt.Println("Welcome in Quize, Please Enter your name")

	var name string
	fmt.Scan(&name) // & use for memory location address refrence
	fmt.Printf("Hello,%v welcome to quize game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	fmt.Println(age >= 10)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You can't play!")
		return
	}

}