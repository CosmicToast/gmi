package gmi

import (
	"fmt"
	"strings"
)

const maxHeading = 3

// Heading represents a Heading in a TOC context, alongside its associated 3 levels within the overall structure
type Heading struct {
	HL [maxHeading]int // current section number, extendable in case we get >3 headings
	// for example, {1, 3, 0} is # followed by 3 ##s
	H *HeadingLine
}

// are all entries at index and forward 0?
func followingZero(i [maxHeading]int, index int) bool {
	for l := len(i); index < l; index++ {
		if i[index] != 0 {
			return false
		}
		index++
	}
	return true
}

// Numbered returns a pretty-printed string representing the numbering component of the heading in the document
func (r Heading) Numbered() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%d.", r.HL[0])                // always print the first one
	for index := 1; index <= maxHeading; index++ { // start at the second one
		if followingZero(r.HL, index) {
			return b.String()
		}
		fmt.Fprintf(&b, "%d.", r.HL[index])
	}
	return b.String()
}

func (r Heading) String() string {
	return fmt.Sprintf("%s %s", r.Numbered(), r.H.Data())
}

// TOC returns the calculated Table of Contents
// The returned structure is the list of just the headings of the parsed output.
// Force forces the recalculation of the TOC if it was already calculated, in case you changed the underlying structure.
func (p *Parser) TOC(force bool) []*Heading {
	if p.toc == nil || force {
		p.calcTOC()
	}
	return p.toc
}

func (p *Parser) calcTOC() {
	p.toc = nil // in case of force
	var hl [maxHeading]int
	for _, v := range p.Lines {
		l := v.Level()
		if l < 1 || l > maxHeading { // skip invalid
			continue
		}

		// 1 -> 0++, 1 = 0, 2 = 0
		// 2 -> 1++, 2 = 0
		// 3 -> 2++
		hl[l-1]++
		for l < maxHeading {
			hl[l] = 0
			l++
		}

		h := Heading{hl, v.(*HeadingLine)}
		p.toc = append(p.toc, &h)
	}
}

// Title will calculate the title from the TOC (without forcing recalculation) and return one is one is found
// Title is considered to be whatever the first '#' heading is
func (p *Parser) Title() string {
	toc := p.TOC(false) // don't force calculation
	for _, v := range toc {
		if v.HL[0] == 1 { // we can't skip 1 since we calculate it
			return v.H.Data()
		}
	}
	return ""
}
