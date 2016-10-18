package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
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

func downloadUrl(url string, path string) error {

    fmt.Printf("path :%s\n", path)
    fmt.Printf("url :%s\n", url)

	// Create the file
	out, err := os.Create(path)
	if err != nil {
	    fmt.Printf("error :%v\n", err)
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
	    fmt.Printf("error :%v\n", err)
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil

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
	if len(os.Args) != 3 {
		usage()
		return
	}

	urls, err := getUrlsFromFile(*inputFile)
	if err != nil {
		fmt.Sprintf("%v\n", err)
		return
	}

	for _, url := range urls {
		fmt.Printf("%s\n", url)
	}
	fmt.Printf("%s", "\n")

	err = downloadUrl("https://ondemand.npr.org/anon.npr-mp3/npr/atc/2016/10/20161016_atc_evangelical_leader_not_for_trump.mp3", "/tmp/2016/10/20161016_atc_evangelical_leader_not_for_trump.mp3")

}

// https://github.com/thbar/golang-playground/blob/master/download-files.go
// http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
// https://gist.github.com/mitsuse/a45c99c7e405ed60e5ce
// https://jawher.me/2015/01/13/parsing-command-line-arguments-shameless-plug-mowcli/
// https://github.com/tealeg/csv2xlsx/blob/master/main.go
