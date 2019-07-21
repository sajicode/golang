//* Race conditions
//* Using atomic functions to fix race conditions
//* The sync/atomic package gives us safe access to numeric types

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	//* counter is a variable incremented by all goroutines
	counter int64

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
		//* Safely Add One to counter
		atomic.AddInt64(&counter, 1)

		//* Yield the thread and be placed back in queue
		runtime.Gosched()
	}
}

//* the AddInt64 function synchronizes the adding of integer values
//* by enforcing that only one goroutine can perform & complete the
//* add operation at a time
//* when goroutines attempt to call any atomic function, they are
//* automatically synchronized against the variable that's referenced
