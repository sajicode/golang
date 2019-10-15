package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	//* Declare & invoke an anonymous function
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again")

	//* Yield to the scheduler
	runtime.Gosched()
}

//* runtime.Gosched() is a way to indicate to the go runtime that
//* we are at a point where we could pause and yield to the scheduler.
//* If the scheduler has other tasks queued up, e.g. other goroutines,
//* it may then run one of them before coming back to this function
//* if we omit the runtime.Gosched() line, our goroutine will never run
//* This is bcos the main function returns (terminates) before the scheduler has a chance to run the goroutine.
