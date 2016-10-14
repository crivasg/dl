package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var app *cli.App

func fromFileCommand() cli.Command {
	command := cli.Command{
		Name:      "formFile",
		ShortName: "ff",
		Usage:     "Obtain URLs from filename",
		Action:    fromFile,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "filename",
				Value: "",
				Usage: "Filename containing the URLs to download",
			},
		},
	}
	return command
}

func fromFile(ctx *cli.Context) {
	filename := ctx.String("filename")
	fmt.Printf("fromFile %s\n", filename)
}

func fromArgsCommand() cli.Command {
	command := cli.Command{
		Name:      "formArgs",
		ShortName: "fa",
		Usage:     "Obtain URLs from command line arguments",
		Action:    fromArgs,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "arguments",
				Value: "",
				Usage: "URLs to download",
			},
		},
	}
	return command
}

func fromArgs(ctx *cli.Context) {
	arguments := ctx.String("arguments")
	fmt.Printf("fromArgs %s\n", arguments)
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
