package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	//"time"
)

func main() {
	fmt.Println("Welcome to maths example in golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4

	fmt.Println("The sum is: ", mynumberOne + int(mynumberTwo))


	// Random Number
	//rand.Seed(45) // Seed is use by algo for genarate random number don't use seed always algo pick seed value automatic
	//rand.Seed(time.Now().UnixNano()) // use time for generate random number
	//fmt.Println(rand.Intn(5) + 1)


	// Genarate random number by crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}