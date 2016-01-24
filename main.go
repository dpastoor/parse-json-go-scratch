package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
