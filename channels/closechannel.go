package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	//* add an additional boolean channel that indicates when finished
	done := make(chan bool)
	until := time.After(5 * time.Second)

	//* passes two channels into <send>
	go send(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			//* when you time-out, let <send> know the process is done
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

//* ch is a receiving channel while done is a sending channel
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			//* when done has a message, shut things down
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
