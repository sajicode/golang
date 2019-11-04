package main

import (
	"testing"
	"time"
)

func TestParallel_1(t *testing.T) {
	//* call parallel function to run test cases in parallel
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

//* go test -v -short -parallel 3
//* The parallel flag indicates that we want to run a maximum of three test cases in parallel
