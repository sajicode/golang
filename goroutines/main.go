//* program that creates two goroutines that display the
//* English alphabet with lower and uppercase letters
//* in a a concurrent fashion
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//* Allocate one logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	//* wg is used to wait for the program to finish
	//* Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2) //* for the two goroutines we intend to run

	fmt.Println("Start Goroutines")

	//* Declare an anonymous function and create a goroutine
	go func() {
		//* Schedule the call to done to tell the main we are done
		defer wg.Done() //* decrements the value of WaitGroup by one

		//* display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//* Declare an anonymous function and create a goroutine
	go func() {
		//* Schedule the call to done to tell the main we are done
		defer wg.Done()

		//* display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//* wait for both goroutines to finish
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

//* A waitGroup is a counting semaphore that can be used to maintain
//* a record of running goroutines
//* when the value of a WaitGroup is greater than zero - wg.Add(2) -
//* the Wait method will block

//* if we allocate a value to a WaitGroup greater than the number of
//* go routines, the application errors out

//* defer is used to schedule other functions from inside the executing
//* function to be called when the function returns.

//* in our case, we use the 'defer' keyword to guarantee that
//* the method call to 'Done' is made once each goroutine
//* is finished with its work
