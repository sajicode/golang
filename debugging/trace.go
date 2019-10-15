package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	//* make a buffer
	buf := make([]byte, 1024)
	//* write the stack into the buffer
	runtime.Stack(buf, false)
	//* print the results
	fmt.Printf("Trace:\n %s\n", buf)
}

//* With runtime.Stack we must supply a presized buffer.
//* Stack takes two arguments. The second is a boolean flag. Setting the boolean flag to true will cause `Stack` to print out stacks for all running goroutines.
