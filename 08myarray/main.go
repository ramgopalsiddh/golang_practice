package main

import "fmt"

func main(){
	fmt.Println("Welcome in array class")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "pineapple"
	fruitList[3] = "mango"

	fmt.Println("Fruits in list is: ", fruitList)
	fmt.Println("Fruits in list is: ", len(fruitList))

	var vegList = [5]string{"potato", "Beans", "Mushroom"}
	fmt.Println("List of vegitable : ", vegList)
	fmt.Println("List of vegitable : ", len(vegList))


}