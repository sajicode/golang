//* program to show how scheduler behaves with two logical
//* processors

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//* Allocate two logical processors for the scheduler to use
	runtime.GOMAXPROCS(2)

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
	// fmt.Println(runtime.NumCPU()) //* print number of CPU cores
}

//* with two logical processors, the goroutines run in parallel
//* this can be seen from the way the results are logged
//* the letters are mixed
