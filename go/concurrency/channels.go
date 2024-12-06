package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s

}

func Channels1() {
	msg = "Hello!"

	wg.Add(2)
	go updateMessage("Goodbye!")
	go updateMessage("Howdy!")
	wg.Wait()

	fmt.Println(msg)
}
