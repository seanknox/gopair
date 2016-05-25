package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gopair"
	app.Usage = "commit code as a pair"
	app.Action = func(c *cli.Context) error {
		fmt.Println("paired!")
		return nil
	}

	app.Run(os.Args)
}
