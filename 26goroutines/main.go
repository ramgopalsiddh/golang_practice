package main

import (
	"fmt"
	"net/http"
	"time"
	"sync"
)

var signals = []string{"test"}

var mut sync.Mutex // ponter

// this work for wait 
var wg sync.WaitGroup   // pointer

func main() {
	fmt.Println("Welcome in go routines")

	// example of wait group
	websitelist := []string{
		"https://ramgopal.dev",
		"https://suraj.dev",
		"https://google.com",
		"https://fb.com",
	}

	// example of wait group
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) // this number define how many wg at wait outside the main
	}

	// this wait group stop exit to main (this send signal main and say my friend wg waight inside main please wait for that)
	wg.Wait()
	fmt.Println(signals)


	// Concurrency and Goroutines example
	// Goroutines :- use in microservices
	// first print hello * 5  then loop tranfer to 2nd
	go greeter("hello")
	// 2nd print world * 5
	greeter("world")
}


// Concurrency and Goroutines example
func greeter(s string) {
	
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}


// waight group example
func getStatusCode(endpoint string){

	// this wait group send signal for done
	defer wg.Done()

	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops problem in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n ", result.StatusCode, endpoint)
	}
}