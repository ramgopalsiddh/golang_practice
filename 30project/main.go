package main

import "fmt"

func main() {
	fmt.Println("Welcome in Quize, Please Enter your name")

	var name string
	fmt.Scan(&name) // & use for memory location address reference
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

	score := 0
	num_question := 2

	fmt.Printf("What is better, RTX 3080 or RTX 3090? ")
	var answer1 string
	var answer2 string
	fmt.Scan(&answer1, &answer2)

	if answer1 + " " + answer2 == "RTX 3090" || answer1 + " " + answer2 == "rtx 3090" {
		fmt.Println("Correct !")
		score++
	} else {
		fmt.Println("Incorrect!")
	}


	fmt.Printf("How many cores does the Ryzen 9 3090x have? ")

	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You Scored %v out of %v\n",score, num_question)
	percent := (float64(score) / float64(num_question)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}