package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var input = flag.String("input", "", "input json file")
var output = flag.String("output", "/dev/stdout", "output file")

func main() {
	flag.Parse()
	f, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(f)
	var doc map[string]interface{}
	err = decoder.Decode(&doc)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	data = append(data, '\n')
	out, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	_, err = out.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
