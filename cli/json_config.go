package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//* JSON configuration parser

//* a type capable of holding JSON values
//* the properties in the struct must be similar to the json file
type configuration struct {
	Enabled bool
	Path    string
	Version int
}

func main() {
	file, _ := os.Open("conf.json") //* open config file
	defer file.Close()

	//* parse json into a variable w/ the variables
	decoder := json.NewDecoder(file)
	//* conf is an instance of configuration
	conf := configuration{}
	//* json content is decoded into conf
	err := decoder.Decode(&conf)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(conf.Version)
}
