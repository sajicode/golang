package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//* creates a channel that will receive a message after 30 seconds
	done := time.After(30 * time.Second)
	//* makes a new channel for passing bytes from Stdin to Stdout.
	//* bcos we haven't specified a size, this channel can hold only one message at a time.
	echo := make(chan []byte)
	//* starts a goroutine to read Stdin, passes it our new channel for communicating
	go readStdin(echo)
	//* uses a select statement to pass data from Stdin to Stdout when received, or to shut-down when the time-out event occurs
	for {
		//* select == switch
		//* there is no default value on our select, so it'll block until either a message is received on <-echo or a message is received on <-done.
		//* when the message is received, select will run the case block and then return control
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

//* Takes a write only channel (chan<-) and sends any received input to that channel
func readStdin(out chan<- []byte) {
	//* copies some data from Stdin into data. Note that File.Read blocks until it receives data
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			//* sends the buffered data over the channel
			out <- data
		}
	}
}
