package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func encodeToString(input interface{}) string {
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%v", input)))
	return fmt.Sprintf("%x", hex.EncodeToString(hash.Sum(nil)))
}

func fileJSONToByte(name string) []byte {
	jsonData, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print(err)
	}
	return jsonData
}
func byteToInterface(jsonData []byte) interface{} {
	var result interface{}
	// unmarshall it
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("error:", err)
	}
	return result
}

func byteToMapInterface(jsonData []byte) map[string]interface{} {
	return byteToInterface(jsonData).(map[string]interface{})
}

func byteToArrayInterface(data []byte) []interface{} {
	return byteToInterface(data).([]interface{})
}

func readBytes(data []byte) (result map[string]bool) {
	result = make(map[string]bool)
	for _, b := range data {
		if b == '[' {
			for _, item := range byteToArrayInterface(data) {
				key := encodeToString(item)
				result[key] = true
			}
		}
		if b == '{' {
			key := encodeToString(byteToMapInterface(data))
			result[key] = true
		}
		return result
	}
	return result
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
		chSource <- readBytes(fileJSONToByte(*fileSource))
	}()

	go func() {
		chTarget <- readBytes(fileJSONToByte(*fileTarget))
	}()
	if compareHashMaps(<-chSource, <-chTarget) {
		fmt.Printf("Two files contains some data!!!\n")
	} else {
		fmt.Printf("Two files contains different data!!!\n")
	}
}
