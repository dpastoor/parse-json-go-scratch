package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jeffail/gabs"
)

func main() {

	filename := "2015-06-01-1.json.gz"

	readLine(filename)

}

func readLine(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	gz, err := gzip.NewReader(f)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer gz.Close()

	r := bufio.NewReaderSize(gz, 100*1024)
	numLines := 0
	i := 0
	var overflow []byte
	var value string
	var types []string
	line, isPrefix, err := r.ReadLine()
	for err == nil {
		if !isPrefix {
			if len(overflow) > 0 {
				line = append(overflow, line...)
				overflow = nil
			}
			jsonParsed, err := gabs.ParseJSON(line)
			if err != nil {
				fmt.Println("error in JSON parsing at line: ", numLines)
				return 0
			}
			value, _ = jsonParsed.Path("type").Data().(string)
			types = append(types, value)
			numLines++
		} else {
			fmt.Println("buffer size to small")
			overflow = append(overflow, line...)
			i++
		}
		line, isPrefix, err = r.ReadLine()
	}
	if err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("number of lines: ", numLines)
	return len(types)
}
