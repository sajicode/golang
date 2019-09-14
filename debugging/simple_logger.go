//* creating a log file & using it as a destination for log messages

package main

import (
	"log"
	"os"
)

func main() {
	//* create a log file
	logfile, _ := os.Create("./log.txt")
	//* make sure it gets closed
	defer logfile.Close()

	//* create a logger
	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	//* send logger some messages
	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
	//* the next line never gets called bcos Fatalln causes the program to exit
	logger.Println("This is the end of the function")
}

//* when creating a logger, we can pass three pieces of info to it
//* 1. The io.Writer where we want to send messages.
//* 2. The prefix for log messages.
//* 3. The list of flags that determine the format of the log message
//* log.LstdFlags | log.Lshortfile sets the date format and then instructs the logger to show the file & line info.
