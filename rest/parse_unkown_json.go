//* parsing JSON with unknown schema
//* Parse the JSON into an interface instead of a struct. After the JSON is in an interface, we can inspect it and use it.package rest

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//* a JSON document to be parsed and unmarshalled
var ks = []byte(
	`{
		"firstName":"Jean",
		"lastName":"Grey",
		"age": 75,
		"education": [
			{
				"institution": "University of Mannitoba",
				"degree": "Masters in Robotics"
			},
			{
				"institution": "Caltech University",
				"degree": "Masters in Engineering"
			}
		],
		"spouse": "William Hill",
		"children": [
			"Thomas Edison",
			"Jane Austen",
			"Mary Poppins"
		]
	}
`)

func main() {
	//* A variable instance of type interface{} to hold the JSON data
	var f interface{}
	//* parse the JSON data and put it into the interface{} type variable
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f)

	//* accessing data
	//* strings
	m := f.(map[string]interface{})
	fmt.Println(m["firstName"], m["lastName"], m["age"])

}
