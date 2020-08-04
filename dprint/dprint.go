package dprint

import (
	"bufio"
	"fmt"
	"io"

	"toast.cafe/x/gmi"
)

// Print will write an AST debug format to the given output writer
func Print(p *gmi.Parser, o io.Writer) {
	w := bufio.NewWriter(o)
	var pft bool
	for _, v := range p.Lines {
		switch v.Type() {
		case gmi.TextType:
			fmt.Fprintf(w, "TEXT: %s\n", v)
		case gmi.LinkType:
			fmt.Fprintf(w, "LINK: %s (%s)\n", v.Link(), v)
		case gmi.PreformatToggleType:
			fmt.Fprintln(w, "PFTT; PFT:")
			pft = !pft
		case gmi.PreformatType:
			if !pft { // should never happen, debug case
				fmt.Fprintf(w, "PFT:  ")
			} else {
				fmt.Fprintf(w, "\t")
			}
			fmt.Fprintf(w, "%s\n", v)
		case gmi.HeadingType:
			fmt.Fprintf(w, "H%d:   %s\n", v.Level(), v)
		case gmi.UnorderedListType:
			fmt.Fprintf(w, "LIST: %s\n", v)
		case gmi.QuoteType:
			fmt.Fprintf(w, "QUOT: %s\n", v)
		}
	}
	w.Flush()
}

// PrintReader will use the reader to parse the document and then invoke Print on the resulting AST
func PrintReader(r io.Reader, o io.Writer) error {
	p := gmi.NewParser(r)
	e := p.Parse()
	if e != nil {
		return e
	}
	Print(p, o)
	return nil
}
