package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Printf("Starting compare-json.go...\n")
	fileSource := flag.String("src-json", "./data/input1.json", "First json file to be compaired")
	fileTarget := flag.String("target-json", "./data/input2.json", "Second json file to be compaired")
	flag.Parse()
	fmt.Printf("File src name: %s\n", *fileSource)
	fmt.Printf("File src name: %s\n", *fileTarget)
}
