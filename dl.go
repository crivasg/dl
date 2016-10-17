package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var app *cli.App

func fromFileCommand() cli.Command {
	command := cli.Command{
		Name:      "fromFile",
		ShortName: "ff",
		Usage:     "Obtain URLs from filename",
		Action:    fromFile,
	}
	return command
}

func fromFile(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		fmt.Printf("Incorrect usage\n")
		return
	}
	fmt.Printf("from File %#v\n", ctx.Args())
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
	fmt.Printf("from Arguments %#v\n", ctx.Args())
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

// https://github.com/thbar/golang-playground/blob/master/download-files.go
// http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
// https://gist.github.com/mitsuse/a45c99c7e405ed60e5ce
// https://jawher.me/2015/01/13/parsing-command-line-arguments-shameless-plug-mowcli/