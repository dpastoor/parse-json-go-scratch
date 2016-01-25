package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/jeffail/gabs"
)

func TestFunc(t *testing.T) {

	filename := "2015-06-01-1.json.gz"
	// file, err := os.Open(filename)
	// //
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// defer file.Close()
	// i := 0
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	if len(line) > 1 {
	//
	// 	}
	// 	i++
	// 	// jsonParsed, err := gabs.ParseJSON([]byte(scanner.Text()))
	// 	// if err != nil {
	// 	// 	fmt.Println("ERROR")
	// 	// }
	// 	// i++
	// 	//
	// 	// value, ok = jsonParsed.Path("type").Data().(string)
	// 	// if ok {
	// 	// 	types = append(types, value)
	// 	// }
	// }
	// fmt.Println(i)
	// // fmt.Println(len(types))
	ReadLine(filename)
}
func ReadLine(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	gz, err := gzip.NewReader(f)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer gz.Close()

	r := bufio.NewReaderSize(gz, 100*1024)
	numLines := 0
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
				return -1
			}
			value, _ = jsonParsed.Path("type").Data().(string)
			types = append(types, value)
		} else {
			fmt.Println("buffer size to small")
			overflow = append(overflow, line...)
		}
		line, isPrefix, err = r.ReadLine()
	}
	if err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("number of lines: ", numLines)
	return len(types)
}
