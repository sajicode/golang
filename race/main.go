//* Race conditions
//* The program below demonstrates race conditions
//! Testing Purposes Only - DO NOT TRY THIS AT WORK!

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//* counter is a variable incremented by all goroutines
	counter int

	//* wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	//* Add a count of two, one for each goroutine
	wg.Add(2)

	//* create two goroutines
	go incCounter(1)
	go incCounter(2)

	//* wait for the goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

//* incCounter increments the package level counter variable
func incCounter(id int) {
	//* Schedule the call to Done to tell main we are done
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//* capture the value of Counter
		value := counter

		//* Yield the thread to give the other goroutine a chance
		//* and be placed back in queue
		//* this forces the scheduler to swap between two goroutines
		//* to exaggerate the effects of the race condition
		runtime.Gosched()

		//* Increment our local value of counter
		value++

		//* store the value back into Counter
		counter = value
	}
}

//* In the above program, the counter variable is read & written
//* to four times, twice by each goroutine, but the value of
//* the counter variable when the program terminates is 2
//* This is because each goroutine overrides the work of the other
//* This happens when the goroutine swap is taking place
//* Each goroutine makes its own copy of the counter variable and
//* then is swapped out for the other goroutine

//* we can use go's built in race detector
//* go build -race & ./race
