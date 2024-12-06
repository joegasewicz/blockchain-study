package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func PrintSomething2(s string, wg *sync.WaitGroup) {
	defer wg.Done() // decrements the wait group by 1
	fmt.Println(s)
}

// Prints
// 8: epsilon
// 1: beta
// 2: delta
// 3: gamma
// 4: pi
// 7: theta
// 5: zeta
// 6: eta
// 0: alpha
// This is the second thing to be printed
func WaitGroups1() {
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x))
	}

	time.Sleep(time.Second)
	printSomething("This is the second thing to be printed")
}

// Results:
// 8: epsilon
// 5: zeta
// 6: eta
// 7: theta
// 2: delta
// 3: gamma
// 0: alpha
// 4: pi
// 1: beta
// This is the second thing to be printed
func WaitGroups2() {
	// add 1 entree for each goroutine you need to wait for
	// Do not copy or pass them around as values, always use pointers!
	// Do not go below zero
	// The Go schedular decides what order they are completed in
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go PrintSomething2(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	printSomething("This is the second thing to be printed")
}
