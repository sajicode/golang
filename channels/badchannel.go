//* An example of a badly closed channel

package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	//* starts a send goroutine with a sending channel
	go send(msg)

	//* loops over a select that watches for messages from send, or for a time out
	for {
		select {
		//* if a message arrives from <send>, prints it
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			//* when the time-out occurs, shuts things down. You pause to ensure that you see the failure before the main goroutine exits
			close(msg)
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan string) {
	//* sends <hello> to the channel every half-second
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}

//* When we run this program, the program panics because <main> closes the msg channel while <send> is still sending messages to it.

//* In Go, the <close> function should only be called by a sender, and in general, it should be done with some protective guards around it.
