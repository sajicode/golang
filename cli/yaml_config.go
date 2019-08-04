package main

//* sample app to parse and print config from a YAML file

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml" //* import 3rd party yaml
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")

	if err != nil {
		fmt.Println(err)
	}

	//* obtain the value of a string with <Get>
	fmt.Println(config.Get("path"))

	//* obtain the value of a boolean with <GetBool>
	fmt.Println(config.GetBool("enabled"))

	//* obtain the value of an integer with <GetInt>
	fmt.Println(config.GetInt("version"))
}

//* we use an external package to read and parse a yaml file
