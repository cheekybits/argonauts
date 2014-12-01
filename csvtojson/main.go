package main

import (
	"log"
	"os"

	"github.com/cheekybits/argonauts/csvtojson/go"
)

func main() {
	converter := csvtojson.New(os.Stdin, os.Stdout)
	if err := converter.Go(); err != nil {
		log.Fatalln(err)
	}
}
