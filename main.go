package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"

	"github.com/jeffail/gabs"
)

func main() {

	filename := "2015-01-01-2.json.gz"
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	gz, err := gzip.NewReader(file)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer gz.Close()

	scanner := bufio.NewScanner(gz)
	var value string
	var ok bool
	var types []string
	for scanner.Scan() {
		jsonParsed, err := gabs.ParseJSON([]byte(scanner.Text()))
		if err != nil {
			fmt.Println("ERROR")
		}

		value, ok = jsonParsed.Path("type").Data().(string)
		types = append(types, value)
		fmt.Println(jsonParsed)
		fmt.Println("")
		fmt.Println(value)
		fmt.Println(ok)
	}
	fmt.Println(len(types))

}
