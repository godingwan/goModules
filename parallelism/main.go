package main

import (
	"fmt"
	"sync"
	"time"
)

/* sync package allows synchronization */
var wg sync.WaitGroup

func main() {
	wg.Add(2) // when the number reaches 0 the goroutines on wait are released
	go first()
	go other()
	wg.Wait()
}

func first() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		// waits 3 millisecond before calling the next interation
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func other() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		// waits 20 millisecond before calling the next iteration
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}
