package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

type File struct {
	URL  string
	Path string
}

type Download struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Name      string `xml:"name,attr"`
	URL       string `xml:"url,attr"`
	Subfolder string `xml:"subfolder,attr"`
	Filename  string `xml:"filename,attr"`
	Hash      string `xml:"hash,attr"`
}

var outputPath = flag.String("output_folder", ``, "Output folder")
var inputFile = flag.String("input_file", ``, "Path the the files containing the urls to download")
var xmlFile = flag.String("xml", ``, "Path to the xml file containing the files to download")

func downloadUrl(url string, path string) error {

	fmt.Printf("path :%s\n", path)
	fmt.Printf("url :%s\n", url)

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dl: %v\n", err)
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

func basenameURL(i string) string {

	u, err := url.Parse(i)
	if err != nil {
		return ""
	}

	slice1 := strings.Split(u.Path, "/")
	filename := slice1[len(slice1)-1]

	return filename

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

	// if non of the file inputs are provided, exit.
	if len(strings.Trim(*xmlFile, " ")) != 0 && len(strings.Trim(*inputFile, " ")) != 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if the outputPath is not provided, set to the TEMP environment variable
	output_dir := *outputPath
	if len(strings.Trim(output_dir, " ")) == 0 {
		if runtime.GOOS == "windows" {
			output_dir = os.Getenv("TEMP")
		} else {
			output_dir = "/tmp"
		}
	}
	fmt.Fprintf(os.Stdout, "Output Dir: %s\n", output_dir)

	if len(strings.Trim(*xmlFile, " ")) != 0 {

		data, err := ioutil.ReadFile(*xmlFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			os.Exit(1)
		}

		var download Download
		err = xml.Unmarshal(data, &download)

		items := download.Items
		for _, item := range items {
			fmt.Fprintf(os.Stdout, "dl: %s\n", item.URL)
		}
	}

	return

	if len(os.Args) != 3 {
		flag.PrintDefaults()
		return
	}

	urls, err := getUrlsFromFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dl: %v\n", err)
		return
	}

	for _, url := range urls {
		fmt.Printf("%s\n", url)
		basename := basenameURL(url)
		filePath := filepath.Join(output_dir, basename)
		err = downloadUrl(url, filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dl: %v\n", err)
		}
	}
	fmt.Printf("%s", "\n")
}

// https://github.com/thbar/golang-playground/blob/master/download-files.go
// http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
// https://gist.github.com/mitsuse/a45c99c7e405ed60e5ce
// https://jawher.me/2015/01/13/parsing-command-line-arguments-shameless-plug-mowcli/
// https://github.com/tealeg/csv2xlsx/blob/master/main.go
