package main

import "fmt"

type slices []int

func newSlice() slices {
	slice := slices{}

	return slice
}

func (d slices) mergeSlice(slice1 []int, slice2 []int) {

	fmt.Printf("%v\n", append(slice1, slice2...))
}
