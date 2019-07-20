package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//* Allocate one logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	//* wg is used to wait for the program to finish
	//* Add a count of two, one for each goroutine
	wg.Add(2) //* for the two goroutines we intend to run

	//* create two goroutines
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	//* wait for the goroutines to finish
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

//* printPrime displays prime numbers for the first 5000 numbers
func printPrime(prefix string) {
	//* Schedule the call to Done to tell main we are done
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

//* =======================================================

//* the program creates two goroutines that print any prime numbers
//* between 1 & 5000 that can be found.
//* This takes a bit of time and so at some point, there is a swap
//* between goroutine A & B - one gives the other time on the thread
//* this shows how the scheduler runs goroutiines concurrently
//* within a logical processor
