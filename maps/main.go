package main

import "fmt"

func main() {
	//* declaring a map
	//* 1
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	// }

	//* 2 empty map
	// var colors map[string]string

	//* 3
	colors := make(map[string]string)

	colors["white"] = "#fff"

	//* delete keys off map

	delete(colors, "white")

	fmt.Println(colors)
}
