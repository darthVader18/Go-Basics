package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait groups
	var wg sync.WaitGroup
	// add, done and wait

	//Mutex (critical section)

	mut := &sync.Mutex{}
	mut.Lock()
	wg.Add(2)
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()

	fmt.Println("Fin")
	wg.Wait()
	fmt.Println("Exit")
}
