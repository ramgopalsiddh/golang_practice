package main

import (
	"fmt"
	"sync"
)

// In channels multiple goroutines talk to each other
func main() {
	fmt.Println("Welcome in Channels Example")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)
	// Recive ONLY channel(<-chan )
	go func(ch <-chan int, wg *sync.WaitGroup){

		// Identify channel is close or open
		val, isChanelOpen := <-myChannel
		fmt.Println(isChanelOpen)

		// Print channel value
		fmt.Println(val)

		//fmt.Println(<-myChannel)

		wg.Done()
	}(myChannel, wg)

	// Send ONLY channel (chan<- )
	go func(ch chan<- int, wg *sync.WaitGroup){
		myChannel <- 5
		myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

}