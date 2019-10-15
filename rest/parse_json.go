//* Parsing JSON data structure into struct
package main

import (
	"encoding/json"
	"fmt"
)

//* A struct that also represents information in JSON. The json tag maps the Name property to name in the JSON
type Person struct {
	Name   string `json:"name"`
	Weapon string `json:"weapon"`
}

//* json represented as a string
var JSON = `{
	"name": "Mad Max",
	"weapon": "Colt 45"
}`

func main() {
	//* an instance of the Person struct to hold the parsed JSON data
	var p Person
	//* parse the JSON data into the instance of the Person struct
	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}
