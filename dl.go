package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var app *cli.App

func main() {

	app = cli.NewApp()

	app.Name = "dl"
	app.Usage = "Downloads files from the command line"
	app.Version = "0.0.0"

	app.Commands = []cli.Command{
		fromFileCommand(),
		fromArgsCommand(),
	}

	app.Run(os.Args)

}

// fromArgs
// fromFile
