package main

import (
	"flag"
	"io"
	"log"
	"os"

	"toast.cafe/x/gmi/dprint"
	"toast.cafe/x/gmi/html"
)

var s struct {
	input  string
	mode   string
	output string

	in  io.Reader
	out io.Writer
}

func main() {
	flag.StringVar(&s.mode, "m", "pretty", "output mode (pretty|debug)")
	flag.StringVar(&s.output, "o", "", "output file (default: stdout)")
	flag.Parse()
	s.input = flag.Arg(0)

	var err error
	if s.input == "" {
		s.in = os.Stdin
	} else {
		s.in, err = os.Open(s.input)
		if err != nil {
			log.Fatalf("could not open file: %s", s.input)
		}
	}
	if s.output == "" {
		s.out = os.Stdout
	} else {
		s.out, err = os.Open(s.output)
		if err != nil {
			log.Fatalf("could not open file: %s", s.output)
		}
	}

	switch s.mode {
	case "pretty":
		log.Fatalf("mode %s is not yet implemented", s.mode)
	case "debug":
		dprint.PrintReader(s.in, s.out)
	default:
		log.Fatalf("Unknown mode: %s", s.mode)
	}
}
