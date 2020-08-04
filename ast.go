package gmi

import "strings"

var (
	_ Line = (*TextLine)(nil)
	_ Line = (*LinkLine)(nil)
	_ Line = (*PreformatToggleLine)(nil)
	_ Line = (*PreformatLine)(nil)
	_ Line = (*HeadingLine)(nil)
	_ Line = (*UnorderedListLine)(nil)
	_ Line = (*QuoteLine)(nil)
)

// Line represents a Line in text/gemini in a logical way
type Line interface {
	CoreType() LineType // CoreType returns the core type of the line (returns Text Type for advanced line types)
	Level() int         // Level returns the level of the heading, only if the type is HeadingType, or 0 otherwise
	Link() string       // Link returns the target of the link, only if the type is LinkType, or the empty string otherwise
	Prefix() string     // Prefix returns the physical prefix of the line, or the empty string otherwise
	String() string     // String() should return the logical contents of the line, not the identifier/prefix
	Type() LineType     // Type returns the true type of the line
}

// ---- types

// TextLine represents a text text/gemini line
type TextLine string

// LinkLine represents a link text/gemini line
type LinkLine struct {
	link string
	name string
}

// PreformatToggleLine represents a "```" toggle in text/gemini
type PreformatToggleLine string

// PreformatLine represents a preformatted text/gemini line
type PreformatLine string

// HeadingLine represents a heading text/gemini line
type HeadingLine struct {
	contents string
	level    int
}

// UnorderedListLine represents an unordered list entry in text/gemini
type UnorderedListLine string

// QuoteLine represents a quoted text/gemini line
type QuoteLine string

// ---- text line

// CoreType returns TextType, which is a valid core type
func (r TextLine) CoreType() LineType { return TextType }

// Level returns zero, since TextLine is not a Heading type
func (r TextLine) Level() int { return 0 }

// Link returns an empty string, since TextLine is not a Link type
func (r TextLine) Link() string { return "" }

// Prefix returns an empty string, since TextLine does not have an identifier prefix
func (r TextLine) Prefix() string { return "" }

// String returns the contents of the text line
func (r TextLine) String() string { return string(r) }

// Type returns TextType
func (r TextLine) Type() LineType { return TextType }

// ---- link line

// CoreType returns LinkType, which is a valid core type
func (r LinkLine) CoreType() LineType { return LinkType }

// Level returns zero, since LinkLine is not a Heading type
func (r LinkLine) Level() int { return 0 }

// Link returns the target link
func (r LinkLine) Link() string { return r.link }

// Prefix returns "=>", which is the identifier for link lines
func (r LinkLine) Prefix() string { return "=>" }

// String returns the user-friendly link name, if there is one
func (r LinkLine) String() string { return r.name }

// Type returns LinkType
func (r LinkLine) Type() LineType { return LinkType }

// ---- preformat toggle line

// CoreType returns PreformatToggleType, which is a valid core type
func (r PreformatToggleLine) CoreType() LineType { return PreformatToggleType }

// Level returns zero, since PreformatToggleLine is not a Heading type
func (r PreformatToggleLine) Level() int { return 0 }

// Link returns the empty string, since PreformatToggleType is not a Link type
func (r PreformatToggleLine) Link() string { return "" }

// Prefix returns "```", which is the identifier for preformat toggle lines
func (r PreformatToggleLine) Prefix() string { return "```" }

// String returns the text following the prefix, but only for the on-toggle one (as per the spec)
func (r PreformatToggleLine) String() string { return string(r) }

// Type returns PreformatToggleType
func (r PreformatToggleLine) Type() LineType { return PreformatToggleType }

// ---- preformat line

// CoreType returns PreformatType, which is a valid core type
func (r PreformatLine) CoreType() LineType { return PreformatType }

// Level returns zero, since PreformatLine is not a Heading type
func (r PreformatLine) Level() int { return 0 }

// Link returns the empty string, since PreformatLine is not a Link type
func (r PreformatLine) Link() string { return "" }

// Prefix returns the empty string, since PreformatLine does not have an identifier prefix
func (r PreformatLine) Prefix() string { return "" }

// String returns the verbatim contents of the preformatted line
func (r PreformatLine) String() string { return string(r) }

// Type returns PreformatType
func (r PreformatLine) Type() LineType { return PreformatType }

// ---- heading line

// CoreType returns TextType, which is an approximation for the Heading type
func (r HeadingLine) CoreType() LineType { return TextType }

// Level returns the level of the heading
func (r HeadingLine) Level() int { return r.level }

// Link returns the empty string, since HeadingLine is not a Link type
func (r HeadingLine) Link() string { return "" }

// Prefix returns the number of "#"s corresponding to the level of the heading
func (r HeadingLine) Prefix() string { return strings.Repeat("#", r.level) }

// String returns the heading itself
func (r HeadingLine) String() string { return r.contents }

// Type returns HeadingType
func (r HeadingLine) Type() LineType { return HeadingType }

// ---- unordered list line

// CoreType returns TextType, which is an approximation for the Unordered List type
func (r UnorderedListLine) CoreType() LineType { return TextType }

// Level returns zero, since UnorderedListLine is not a Heading type
func (r UnorderedListLine) Level() int { return 0 }

// Link returns the empty string, since UnorderedListLine is not a Link type
func (r UnorderedListLine) Link() string { return "" }

// Prefix returns "* ", which is the identifier for the unordered list lines
// The consumer is expected to ignore this and use bullet symbols for styling, but it is still present for CoreType consumers.
func (r UnorderedListLine) Prefix() string { return "* " }

// String returns the contents of the list entry
func (r UnorderedListLine) String() string { return string(r) }

// Type returns UnorderedListType
func (r UnorderedListLine) Type() LineType { return UnorderedListType }

// ---- quote line

// CoreType returns TextType, which is an approximation for the Quote type
func (r QuoteLine) CoreType() LineType { return TextType }

// Level returns zero, since QuoteType is not a Heading type
func (r QuoteLine) Level() int { return 0 }

// Link returns the empty string, since QuoteType is not a Link type
func (r QuoteLine) Link() string { return "" }

// Prefix returns ">", which is the identifier for the quote lines
func (r QuoteLine) Prefix() string { return ">" }

// String returns the quote contents
// Note that any whitespace immediately after the prefix is discarded, despite not being mentioned in the spec.
func (r QuoteLine) String() string { return string(r) }

// Type returns QuoteType
func (r QuoteLine) Type() LineType { return QuoteType }
