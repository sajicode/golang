package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name") //* this has to be same w/ Name:
		fmt.Printf("Hello %s!\n", name)
		return nil
	}
	app.Run(os.Args)
}

//* run <go get gopkg.in/urfave/cli.v1> if package does not imprt auto
//* cli.StringFlag props => name flag, default value, description
//* a default action is also set to print "Hello ..."
