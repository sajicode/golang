//* telling the difference between two errors naively
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// ErrTimeout... the timeout error instance
var ErrTimeout = errors.New("The request timed out")

// ErrRejected... the rejection error instance
var ErrRejected = errors.New("The request was rejected")

// a random number generator with a fixed source
var random = rand.New(rand.NewSource(35))

func main() {
	// calls the stubbed-out SendRequest function
	response, err := sendRequest("Halos")
	// Handles the time-out condition with retries
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying...")
		response, err = sendRequest("Hello")
	}
	// handles any other error as a failure
	if err != nil {
		fmt.Println(err)
	} else {
		// if there is no error, prints the result
		fmt.Println(response)
	}
}

// defines a function that superficially behaves like a message sender
func sendRequest(req string) (string, error) {
	// instead of sending a message, randomly generates behavior
	switch random.Int() % 3 {
	// handles the timeout with retries
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
