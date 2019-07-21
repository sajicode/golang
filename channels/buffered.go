package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  //* number of goroutines to use
	taskLoad         = 10 //* Amount of work to process
)

var wg sync.WaitGroup

//* init is called to initialize the package by the go runtime
//* prior to any other code being executed
func init() {
	//* Seed the random number generator
	rand.Seed(time.Now().Unix())
}

func main() {
	//* create a buffered channel to manage the task load
	tasks := make(chan string, taskLoad)

	//* Launch goroutines to handle the work
	wg.Add(numberGoroutines)

	//* create goroutines
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//* Add a bunch of work to get done by sending 10 strings into channel
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	//* close the channel so the goroutines will quit
	//* when all the work is done
	//* when a channel is closed, goroutines can still receive but
	//* cannot send on the channel
	close(tasks)

	//* wait for all the work to get done
	wg.Wait()
}

//* worker is launched as a goroutine to process work
//* from the buffered channel
func worker(tasks chan string, worker int) {
	//* Report that we just returned
	defer wg.Done()

	for {
		//* wait for task to be assigned
		//* this is a blocking operation
		task, ok := <-tasks
		if !ok {
			//* This means the channel is empty and closed
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		//* Display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		//* Randomly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//* Display we finished the work
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

//* run independent of main.go
