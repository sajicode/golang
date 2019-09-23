//* logging to a network via udp instead of tcp to deal with back pressure -- Go in Practice pp121

package main

import (
	"log"
	"net"
	"time"
)

func main() {
	//* adds an explicit timeout
	timeout := 30 * time.Second
	//* Dials a UDP connection instead of a TCP one
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)

	if err != nil {
		panic("failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}

//* to start this app, run <nc -luk 1902> to start a UDP server.
//* then run file

//* Advantages of UDP
//* 1. The app is resistant to back pressure and log server outages. If the log server hiccups, it may lose some UDP packets, but the client won’t be impacted.
//* 2. Sending logs is faster even without back pressure.
//* 3. The code is simple.

//! Disadvantages

//* Log messages can get lost easily. UDP doesn’t equip you to know whether a message was received correctly.
//* Log messages can be received out of order. Large log messages may be packetized and then get jumbled in transition. Adding a timestamp to the message (as you’ve done) can help with this, but not totally resolve it.
//* Sending the remote server lots of UDP messages may turn out to be more likely to overwhelm the remote server, because it can’t manage its connections and slow down the data intake. Although your app may be immune to back pressure, your log server may be worse off.
