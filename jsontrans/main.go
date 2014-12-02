package main

import (
	"log"
	"os"

	"github.com/cheekybits/argonauts/jsontrans/go"
)

func main() {
	if err := jsontrans.New(os.Stdin, os.Stdout, os.Args).Go(); err != nil {
		log.Fatalln(err)
	}
}
