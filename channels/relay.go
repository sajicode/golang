//* Using an unbuffered channel to simulate a relay race
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//* create an unbuffered channel
	baton := make(chan int)

	//* Add a count of one for the last runner
	wg.Add(1)

	//* first runner to his mark
	go Runner(baton)

	//* Start the race by giving baton to runner
	baton <- 1

	//* wait for the race to finish
	wg.Wait()
}

//* "Runner simulates a runner running in the relay race"
func Runner(baton chan int) {
	var newRunner int

	//* wait to receive the baton, passed from line 23
	runner := <-baton

	//* Start running around the track
	fmt.Printf("Runner %d Running with Baton\n", runner)

	//* New runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	//* Running around the track
	time.Sleep(100 * time.Millisecond)

	//* Is the race over ?
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	//* Exchange the baton for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}

//* Run independent of main.go
