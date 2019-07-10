package main

import "fmt"

type slices []int

func newSlice() slices {
	slice := slices{}

	return slice
}

//* merge two slices
func (d slices) mergeSlice(slice1 []int, slice2 []int) {

	fmt.Printf("%v\n", append(slice1, slice2...))
}

//* iterating over a slice with a for loop from the 2nd element
func (d slices) iterateSlice() {
	for index := 1; index < len(d); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, d[index])
	}
}
