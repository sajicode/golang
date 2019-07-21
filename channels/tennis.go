package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//* create an unbuffered channel
	court := make(chan int)

	//* Add a count of two, one for each goroutine
	wg.Add(2)

	//* Launch two players. At this point, both goroutines are locked
	//* waiting to receive the ball
	go player("Nadal", court)
	go player("Federer", court)

	//* start the set by passing a value to the channel
	court <- 1

	//* Wait for the game to finish
	wg.Wait()
}

//* player simulates a person playing the game of tennis
func player(name string, court chan int) {
	//* Schedule the call to Done to tell main we are done
	defer wg.Done()

	//* the game is played within the loop
	for {
		//* wait for the ball to be hit back to us by receiving frm channel
		ball, ok := <-court
		//* checks the ok value for false i.e. was the ball returned ?
		//* we get false if the channel was closed
		if !ok {
			//* If the channel was closed i.e. no ball returned, we won
			fmt.Printf("Player %s Won\n", name)
			return
		}

		//* Pick a random number and see if we miss the ball
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			//* Close the channel to signal we lost
			close(court)
			return
		}

		//* Display and then increment the hit count by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		//* Hit the ball back to the oppossing player
		court <- ball
	}
}

//* run alone not with main.go
//* ignore error on func main declaration
