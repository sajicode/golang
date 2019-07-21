package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//* counter is a variable incremented by all goroutines
	counters int

	//* wg is used to wait for programs to finish
	wgs sync.WaitGroup

	//* mutex is used to define a critical section of code
	mutex sync.Mutex
)

func main() {
	//* Add a count of two, one for each goroutine
	wgs.Add(2)

	//* create two goroutines
	go incrCounter(1)
	go incrCounter(2)

	//* wait for the goroutines to finish
	wgs.Wait()
	fmt.Printf("Final Counter: %d\n", counters)
}

//* incCounter increments the package level counter variable
//* using the Mutex to synchronize and provide safe access
func incrCounter(id int) {
	//* Schedule the call to Done to tell main we are done
	defer wgs.Done()

	for count := 0; count < 2; count++ {
		//* Only allow one goroutine thru this critical section at a time
		mutex.Lock()
		//* the curly braces below are not necessary
		{
			//* Capture the value of counter
			value := counters

			//* Yield the thread and be placed back in queue
			runtime.Gosched()

			//* Increment our local value of counter
			value++

			//* Store the value back into counter
			counters = value
		}
		mutex.Unlock()
		//* Release the lock and allow any waiting goroutine through
	}
}

//* run alone not with main.go
