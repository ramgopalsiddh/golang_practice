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


	fmt.Printf("What is better, RTX 3080 or RTX 3090? ")
	var answer1 string
	var answer2 string
	fmt.Scan(&answer1, &answer2)

	if answer1 + " " + answer2 == "RTX 3090" {
		fmt.Println("Correct !")
	} else if answer1 + " " + answer2 == "rtx 3090" {
		fmt.Println("Correct !")
	} else {
		fmt.Println("Incorrect!")
	}


	fmt.Printf("How many cores does the Ryzen 9 3090x have? ")

	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}
}