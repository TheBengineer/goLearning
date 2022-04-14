package main

import (
	"fmt"
	"time"
)

func drill(belt chan string, ore string) {
	for {
		time.Sleep(1 * time.Second)
		belt <- ore
	}
}

func Inserter1(belt chan string) {
	item := ``
	for {
		time.Sleep(1 * time.Second)
		belt <- item
	}
}

func Inserter2(belt chan string) {

	for {
		select {
		case item := <-belt:
			fmt.Printf("[Inserter2] I got %s\n", item)
		}
	}
}

func main() {
	belt := make(chan string)
	go drill(belt, "Iron")
	go drill(belt, "Iron")
	go drill(belt, "Iron")
	go drill(belt, "Iron")
	go Inserter2(belt)
	for {
		// this loops forever so our "factory" never stops
	}
}
