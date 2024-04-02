package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome in Dice number game by switch case")

	// Generate randum number(Dice) 
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	// Use Switch case for print dice result
	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll dice again")
	default:
		fmt.Println("What was that!")
	}

}