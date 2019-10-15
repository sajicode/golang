//* compressing files in parallel with waitgroup
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	//* a waitgroup doesn't need to be initialized
	var wg sync.WaitGroup

	//* because we want to refernce <i> outside the loop, we declare the variables here
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		//* for every file we add, we tell the wait group that we are waiting for one more compress operation
		wg.Add(1)
		//* the function calls <compress> & then notifies the wait group it is done
		go func(filename string) {
			compress(filename)
			wg.Done()
			//* because we are calling a goroutine in a for loop, we need to do a little trickery with the parameter passing
		}(file)
		//* we want to ensure the value of <file> is passed into each goroutine as it is scheduled
	}
	//* the outer goroutine (main) waits until all the compressing goroutines have called wg.Done
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	//* opens the source file for reading
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	//* opens a destination file, with the .gz extension added to the source file's name.
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	//* The gzip Writer compresses data & then writes it to the underlying file
	gzout := gzip.NewWriter(out)
	//* the io.Copy() func does all the copying for us
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}

//* A waitgroup is a message-passing facility that signals a waiting goroutine when it's safe to proceed.

//* To use it, you tell the wait group when you want to wait for something & then you signal it again when that is done.

//* A wait group doesn't need to know more about the things it's waiting for other than (a) the number of things it's waiting for and (b) when each is done
