//* Tool that compresses an arbitrary number of individual files to show how we can wait for goroutines to complete before continuing with the main function.

package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	//* collects a list of files passed in on the command line
	for _, file := range os.Args[1:] {
		compress(file)
	}
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
