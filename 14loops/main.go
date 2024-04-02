package main

import "fmt"

func main() {
	fmt.Println("Welcome in loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// Structrer of for loop
	for d:=0; d< len(days); d++{
		fmt.Println(days[d])
	}

	// for each loop
	for d := range days{
		fmt.Println(days[d])
	}

	// use index and print index position with relevent day
	// if want to remove index use _  in loop and remove index & value from print statement 
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day )
	}


	// similer like while loop
	rougueValue := 1

	for rougueValue < 10 {

		// use continue for continue loop
		if rougueValue == 6 {
			rougueValue++
			continue
		}

		// tranfer to any other lable
		if rougueValue == 7 {
			goto visitportfolio
		}
		// Break use for terminate loop 
		if rougueValue == 8 {
			break
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++ // prevent from infinite loop

	}


	// Lable for tranfer from loop
	visitportfolio:
		fmt.Println("Visit https://www.ramgopal.dev")

}