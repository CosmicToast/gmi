package gmi

import "testing"

func parseSingleLine(s string) Line {
	p := NewParser(nil)
	return p.parseLine(s)
}

func equalLines(l1, l2 Line) bool {
	return l1.Type() == l2.Type() &&
		l1.Level() == l2.Level() &&
		l1.Link() == l2.Link() &&
		l1.Prefix() == l2.Prefix() &&
		l1.String() == l2.String()
}

func TestEmpty(t *testing.T) {
	var expect = []struct {
		in  string
		out Line
	}{
		{"", TextLine("")},
		{"=>", LinkLine{"", ""}},
		{"=> ", LinkLine{"", ""}},
		{"#", HeadingLine{"", 1}},
		{"##", HeadingLine{"", 2}},
		{"###", HeadingLine{"", 3}},
		{"# ", HeadingLine{"", 1}},
		{"## ", HeadingLine{"", 2}},
		{"### ", HeadingLine{"", 3}},
		{"*", TextLine("*")},
		{"* ", UnorderedListLine("")},
		{">", QuoteLine("")},
		{"> ", QuoteLine("")},
	}

	for _, v := range expect {
		out := parseSingleLine(v.in)
		if !equalLines(out, v.out) {
			t.Errorf("unexpected line for %s", v.in)
		}
	}
}
