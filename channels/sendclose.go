//* Closing a channel from a sender

package main

import "time"

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)

	//* loops over a select with two channels and a default
	for {
		select {
		//* if you get a message over your main channel, print something
		case <-ch:
			println("Got message.")
		case <-timeout:
			println("Time out")
			return
		default:
			//* by default, sleeps for a bit. this makes the example easier to work with
			println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	//* sends a single message over the channel and then closes the channel
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Sent and closed")
}

//* we do not get the expected response here because a closed channel always returns the channel's nil value, so <send> sends one true value and then closes the channel.

//* Each time the <select> examines ch after ch is closed, it'll receive a false value (the nil value on a bool channel).
