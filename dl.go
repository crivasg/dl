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
	"strings"
)

type FileD struct {
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

	fileinfo, err1 := os.Stat(path)
	if err1 == nil {
		fmt.Fprintf(os.Stdout, "dl: %s %.2fMB\n",
			fileinfo.Name(), float64(fileinfo.Size())/(1024.0*1024.0))
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
	if len(strings.Trim(*xmlFile, " ")) == 0 && len(strings.Trim(*inputFile, " ")) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if the outputPath is not provided, set to the TEMP environment variable
	output_dir := *outputPath
	if len(strings.Trim(output_dir, " ")) == 0 {
		output_dir = os.TempDir()
	}
	fmt.Fprintf(os.Stdout, "Output Dir: %s\n", output_dir)

	var urls []FileD

	if len(strings.Trim(*xmlFile, " ")) != 0 {
		data, err := ioutil.ReadFile(*xmlFile)
		if err == nil {
			var download Download
			err = xml.Unmarshal(data, &download)

			items := download.Items
			for _, item := range items {
				fullpath := filepath.Join(output_dir, item.Subfolder, item.Filename)
				urls = append(urls, FileD{URL: item.URL, Path: fullpath})
			}
		}
	}

	if len(strings.Trim(*inputFile, " ")) != 0 {
		data, err := ioutil.ReadFile(*inputFile)
		if err == nil {
			for _, line := range strings.Split(string(data), "\n") {
				basename := basenameURL(line)
				fullpath := filepath.Join(output_dir, basename)
				urls = append(urls, FileD{URL: line, Path: fullpath})
			}
		}
	}

	for index, url := range urls {
		fmt.Fprintf(os.Stdout, "dl: %-3d %s --> %s\n", index+1, url.URL, url.Path)
		dirname := filepath.Dir(url.Path)

		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			os.MkdirAll(dirname, 0777)
		}

		err := downloadUrl(url.URL, url.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dl: %v\n", err)
		}
	}
}

// test example:
// ./dl -xml=sample/download_utils.xml -input_file=sample/download_radiolab.txt -output_folder=/tmp

// https://github.com/thbar/golang-playground/blob/master/download-files.go
// http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
// https://gist.github.com/mitsuse/a45c99c7e405ed60e5ce
// https://jawher.me/2015/01/13/parsing-command-line-arguments-shameless-plug-mowcli/
// https://github.com/tealeg/csv2xlsx/blob/master/main.go
