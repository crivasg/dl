package main

import (
	"flag"
	"fmt"
	"os"
)

var outputPath = flag.String("o", "", "Output folder")
var inputFile = flag.String("f", "", "Path the the files containing the urls to download")

func usage() {
	fmt.Printf(`%s: -f=<URLs Input File> -o=<Output Folder>
`, os.Args[0])
}

func main() {

	flag.Parse()
	if len(os.Args) != 2 {
		usage()
		return
	}

}

// fromArgs
// fromFile

// https://github.com/thbar/golang-playground/blob/master/download-files.go
// http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
// https://gist.github.com/mitsuse/a45c99c7e405ed60e5ce
// https://jawher.me/2015/01/13/parsing-command-line-arguments-shameless-plug-mowcli/
