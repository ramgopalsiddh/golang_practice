package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome in Slices")

	var fruitList = []string{"Apple", "Mango", "Banan"}
	fmt.Printf("Type of fruitList is : %T \n", fruitList)

	// Add more values in slices
	fruitList = append(fruitList, "Kivi", "Peach")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // only first 2 iteams print because last range is not including
	fmt.Println(fruitList)


	highScores := make([]int, 5) // if you allocate 5 and create 6 value then got error

	highScores[0] = 246
	highScores[1] = 257
	highScores[2] = 167
	highScores[3] = 940
	highScores[4] = 567
	// highScores[5] = 367


	highScores = append(highScores, 555, 666, 777) // append realocate memory for more data

	fmt.Println(highScores)

	// Sort slices in increasing order
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}