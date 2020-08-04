package main

import (
	"log"
	"os"

	"toast.cafe/x/gmi/dprint"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	dprint.PrintReader(file, os.Stdout)
}
