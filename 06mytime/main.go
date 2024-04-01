package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to Time module")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//  The 01-02-2006 15:04:05 Monday formating way that described in docs.
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Add custom time, date, day etc. setup
	createdDate := time.Date(2020, time.August, 12, 23, 20, 0, 0, time.UTC )
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}