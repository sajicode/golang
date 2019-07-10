package main

import "fmt"

func main() {
	fmt.Println("Go there>>>")

	// s1 := []int{1, 2, 3}
	// s2 := []int{4, 5, 6}

	// slice1 := newSlice()

	// slice1.mergeSlice(s1, s2)
	slice := newSlice()

	slice = append(slice, 4, 5, 6, 7, 8)

	slice.iterateSlice()
}
