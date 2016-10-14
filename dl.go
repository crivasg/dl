package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var app *cli.App

func fromFileCommand() cli.Command {
	command := cli.Command{
		Name: "fromFile",
		//ShortName: "ff",
		Usage:  "Obtain URLs from filename",
		Action: fromFile,
	}
	return command
}

func fromFile(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		fmt.Printf("Incorrect usage\n")
		return
	}
	fmt.Printf("form File %#v\n", ctx.Args())
}

func fromArgsCommand() cli.Command {
	command := cli.Command{
		Name:      "fromArgs",
		ShortName: "fa",
		Usage:     "Obtain URLs from command line arguments",
		Action:    fromArgs,
	}
	return command
}

func fromArgs(ctx *cli.Context) {
	fmt.Printf("form File %#v\n", ctx.Args())
}

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
