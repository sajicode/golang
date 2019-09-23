package main

import (
	"log"
	"net"
)

func main() {
	//* connects to the log server
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("failed to connect to localhost:1902")
	}
	//* make sure you clean up by closing the connection, even on panic
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	//* Sends log messages to the network connection
	logger := log.New(conn, "example ", f)

	logger.Println("This is a regular message.")
	//* logs a message and then panics - don't use Fatalln here
	logger.Panicln("This is a panic")
}

//* A tcp connection is created with net.Dial.
//* when we use log.Fatal, the deferred function - conn.Close() - isn't called. This is because log.Fatal calls os.Exit which immediately terminates the program without unwinding the function stack

//* to run this app, run <nc -lk -p 1902> on a separate terminal.
//* then run <network_logger.go>
