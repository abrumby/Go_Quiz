package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvfileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	_ = csvfileName
	file, err := os.Open(*csvfileName)
	if err != nil {
		fmt.Printf("Failed to open the csv file: %s", *csvfileName)
		os.Exit(1)
	}
	_ = file
}
