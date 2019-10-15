package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

func main() {
	//* create structure to hold the config values
	config := struct {
		Section struct {
			Enabled bool
			Path    string
			Version int
		}
	}{}
	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
	fmt.Println(config.Section.Version)
}
