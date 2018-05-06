package main

import (
	"fmt"

	"gopkg.in/gcfg.v1" // third party packge to parse .ini files
)

func main() {
	// structure to hold config values
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
