package gmi

// LineType represents a text/gemini line type
type LineType int

// Line Types
const (
	TextType LineType = iota
	LinkType
	PreformatToggleType
	PreformatType
	HeadingType
	UnorderedListType
	QuoteType
)
