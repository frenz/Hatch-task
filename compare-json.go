package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func encodeToString(input interface{}) string {
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%v", input)))
	return fmt.Sprintf("%x", hex.EncodeToString(hash.Sum(nil)))
}

func compareHashMaps(mapSource, mapTarget map[string]bool) bool {
	for k := range mapSource {
		if mapTarget[k] != true {
			return false
		}
		delete(mapTarget, k)
		delete(mapSource, k)
	}
	return (len(mapSource) + len(mapTarget)) == 0
}
func streamToMapHash(filename string) (res map[string]bool, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return res, err
	}

	defer f.Close()

	fileStream := bufio.NewReader(f)
	dec := json.NewDecoder(fileStream)
	res = make(map[string]bool)

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		return res, err
	}

	if fmt.Sprintf("%v", t) == "{" {
		f.Close()
		f, err := os.Open(filename)
		if err != nil {
			return res, err
		}

		defer f.Close()
		fileStream := bufio.NewReader(f)
		dec := json.NewDecoder(fileStream)
		res = make(map[string]bool)
		var m interface{}
		err = dec.Decode(&m)
		if err != nil {
			return res, err
		}

		res[encodeToString(m)] = true
	} else {
		// while the array contains values
		for dec.More() {
			var m interface{}
			// decode an array value (Message)
			err = dec.Decode(&m)
			if err != nil {
				return res, err
			}

			res[encodeToString(m)] = true
		}
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}

func main() {
	fmt.Printf("Starting compare-json.go...\n")
	fileSource := flag.String("src-json", "./data/input1.json", "First json file to be compaired")
	fileTarget := flag.String("tgt-json", "./data/input2.json", "Second json file to be compaired")
	flag.Parse()
	fmt.Printf("File src name: %s\n", *fileSource)
	fmt.Printf("File tgt name: %s\n", *fileTarget)
	chSource := make(chan map[string]bool)
	chTarget := make(chan map[string]bool)
	go func() {
		m, err := streamToMapHash(*fileSource)
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(0)
		}
		chSource <- m
	}()

	go func() {
		m, err := streamToMapHash(*fileTarget)
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(0)
		}
		chTarget <- m
	}()
	if compareHashMaps(<-chSource, <-chTarget) {
		fmt.Printf("Two files contains some data!!!\n")
	} else {
		fmt.Printf("Two files contains different data!!!\n")
	}
}
