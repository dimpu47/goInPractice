package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	// Create a new app
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print Hello World"
	// Set up a global flag
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:	"n, name",
			Value:	"World",
			Usage:	"Who to say hello to.",
		},
	}
	// Define an action to run
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s! \n", name)
		return nil
	}
	// Run the app.
	app.Run(os.Args)
}
