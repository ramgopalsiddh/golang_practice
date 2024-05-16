package main

import (
	"fmt"
	"sync"
)

// Run for detailed info about race condition use command :- `go run --race main.go`


func main() {
	// print welcome message
	fmt.Println("Welcome in Mutex and Await groups example")

	wg := &sync.WaitGroup{}
	// use mutex for resolve race condition
	mut := &sync.RWMutex{}

	// inislize value
	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Three R")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}