package html

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"strings"

	"toast.cafe/x/gmi"
)

// Contents writes out html for the contents of the document
// WARNING: no HTML escaping is done!
func Contents(p *gmi.Parser, o io.Writer) {
	w := bufio.NewWriter(o)
	var pft bool
	var list bool
	var br int
	for _, v := range p.Lines {
		// terminating lists
		if list && v.Type() != gmi.UnorderedListType {
			fmt.Fprint(w, "</ul>\n")
			list = false
		}
		// <br> handling
		if v.Type() == gmi.TextType && v.String() == "" {
			br++
			continue
		}
		if br > 0 {
			fmt.Fprint(w, strings.Repeat("<br>\n", br-1))
			br = 0
		}
		// normal handling
		switch v.Type() {
		case gmi.TextType:
			fmt.Fprintf(w, "<p>%s</p>\n", v)
		case gmi.LinkType:
			n := v.String()
			if n == "" {
				n = v.Link()
			}
			fmt.Fprintf(w, "<a href='%s'>%s</a>\n", v.Link(), n)
		case gmi.PreformatToggleType:
			if pft {
				fmt.Fprint(w, "</pre>\n")
			} else {
				fmt.Fprint(w, "<pre>")
			}
			pft = !pft
		case gmi.PreformatType:
			fmt.Fprintln(w, v)
		case gmi.HeadingType:
			fmt.Fprintf(w, "<a href='%s'><h%d>%s</h%[2]d></a>\n", url.QueryEscape(v.String()), v.Level(), v)
		case gmi.UnorderedListType:
			if !list {
				fmt.Fprint(w, "<ul>\n")
				list = true
			}
			fmt.Fprintf(w, "<li>%s</li>\n", v)
		case gmi.QuoteType:
			fmt.Fprintf(w, "<blockquote>%s</blockquote>\n", v)
		}
	}
	w.Flush()
}

// ContentsReader will use the reader to parse the document and then invoke Contents on the resulting AST
func ContentsReader(r io.Reader, o io.Writer) error {
	p := gmi.NewParser(r)
	e := p.Parse()
	if e != nil {
		return e
	}
	Contents(p, o)
	return nil
}

// TOC writes out html for a navigation component of the document
// WARNING: no HTML escaping is done!
// TODO: actually write this
func TOC(p *gmi.Parser, o io.Writer) {}

// TOCReader will use the reader to parse the document and then invoke TOC on the resulting AST
func TOCReader(r io.Reader, o io.Writer) error {
	p := gmi.NewParser(r)
	e := p.Parse()
	if e != nil {
		return e
	}
	TOC(p, o)
	return nil
}
