package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var outputPath = flag.String("o", "/tmp/", "Output folder")
var inputFile = flag.String("f", "", "Path the the files containing the urls to download")

func usage() {
	fmt.Printf(`%s: -f=<URLs Input File> -o=<Output Folder>
`, os.Args[0])
}

func getUrlsFromFile(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr_line := strings.Trim(scanner.Text(), "\t ")
		match, _ := regexp.MatchString("^htt(p|ps)://", curr_line)
		if match == true {
			lines = append(lines, curr_line)
		}
	}
	return lines, scanner.Err()

}

func main() {

	flag.Parse()
	if len(os.Args) != 2 {
		usage()
		return
	}

	url, err := getUrlsFromFile(*inputFile)
	if err != nil {
		fmt.Sprintf("%v\n", err)
		return
	}

}

// https://github.com/thbar/golang-playground/blob/master/download-files.go
// http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
// https://gist.github.com/mitsuse/a45c99c7e405ed60e5ce
// https://jawher.me/2015/01/13/parsing-command-line-arguments-shameless-plug-mowcli/
// https://github.com/tealeg/csv2xlsx/blob/master/main.go
