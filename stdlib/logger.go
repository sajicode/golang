package main

import "log"

//* demo for using the base log package

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// * Println writes to the standard logger
	log.Println("message")

	// * Fatalln is println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// * Panicln is a Println() followed by a call to panic()
	log.Panicln("panic message")
}
