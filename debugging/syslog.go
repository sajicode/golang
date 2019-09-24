//* generating a go logger that is backed to syslog
package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	//* tells the logger how to appear to syslog
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	//* sets the log flags
	flags := log.Ldate | log.Lshortfile
	//* creates a new syslog logger
	logger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		fmt.Printf("Can't attach to syslog: %s", err)
		return
	}
	//* sends a simple message
	logger.Println("This is a test log message")
}

//* run <nc -luk 1902> and then run file
//! file doesn't work for some reason
