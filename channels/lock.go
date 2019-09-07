//* Simple locking with channels

package main

import (
	"fmt"
	"time"
)

func main() {
	//* create a buffered channel with one space
	lock := make(chan bool, 1)
	//* start up to six goroutines sharing the locking channel
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

//* A worker acquires the lock by sending it a message. The first worker to hit this will get the one space, and thus own the lock. The rest will block
func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond)
	//* the space between the lock <-true and the <-lock is "locked"
	fmt.Printf("%d is releasing the lock\n", id)
	//* Releases the lock by reading a value, which then opens that one space on the buffer again so that the next function can lock it
	<-lock
}
