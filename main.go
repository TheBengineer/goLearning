package main

import (
	"fmt"
	"time"
)

// Pulled from here:
// https://blog.ropnop.com/learning-go-concurrency-from-factorio/

func Inserter1(belt chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		belt <- i
		i++
	}
}

func Inserter2(belt chan int) {
	for {
		select {
		case i := <-belt:
			fmt.Printf("[Inserter2] I got %d\n", i)
		}
	}
}

func main() {
	belt := make(chan int)
	go Inserter1(belt)
	go Inserter2(belt)
	for {
		// this loops forever so our "factory" never stops
	}
}
