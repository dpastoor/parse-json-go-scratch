package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFunc(t *testing.T) {

	filename := "2015-01-01-2.json"
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
func ReadLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 100*1024)
	numLines := 0
	i := 0
	var overflow []byte
	line, isPrefix, err := r.ReadLine()
	for err == nil {
		if !isPrefix {
			if len(overflow) > 0 {
				line = append(overflow, line...)
				overflow = nil
				fmt.Println("number of buffer overflows: ", string(line))
				ioutil.WriteFile("temp.json", line, 0644)
			}
			fmt.Println("clean line")
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
	fmt.Println("number of buffer overflows: ", i)
	fmt.Println("number of lines: ", numLines)

}
