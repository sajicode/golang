//* implementation of a server that crashes when a panic occurs on a goroutine.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	//* start a new server listening on port 1026
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}

	//* listens for new client connections and handles any connection errors
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		//* when a connection is accepted, passes it to the handle function
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//* The deferred function handles the panic and makes sure that in all cases the connection is closed
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s", err)
		}
		conn.Close()
	}()
	//* tries to read a line of data from the connection
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	//* if we fail to read a line, prints an error and closes the connection
	if err != nil {
		fmt.Println("Failed to read from socket.")
		conn.Close()
	}
	//* once you get a line of text, passes it to response
	response(data, conn)
}

//* writes the data back out to the socket, echoing it to the client; then closes the connection
func response(data []byte, conn net.Conn) {
	// defer func() {
	// 	conn.Close()
	// }()
	conn.Write(data)
	//* simulate a panic
	panic(errors.New("Pretend I'm a real error\n"))
}

//* now, the server doesn't crash because we have handled possible errors in the <handle> function, no pun intended :-)
