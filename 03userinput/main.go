package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "WelCome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok  OR  error ok syntex
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type ot this rating is %T \n", input)

}