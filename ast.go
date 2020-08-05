package gmi

import (
	"fmt"
	"strings"
)

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
	Type() LineType     // Type returns the true type of the line
	Level() int         // Level returns the level of the heading, if it's a HeadingType, or 0 otherwise

	Data() string   // Data returns the content data of the line
	Meta() string   // Meta returns the secondary data of the line (e.g the link for link lines), if any
	Prefix() string // Prefix returns the documented prefix to the line, without whitespace

	String() string // String (Stringer) implements a post-formatted representation of the original line, including prefix
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

// revive:disable:exported
// implementations are already documented in the interface

// ---- text line

func (r TextLine) CoreType() LineType { return TextType }
func (r TextLine) Type() LineType     { return TextType }
func (r TextLine) Level() int         { return 0 }
func (r TextLine) Data() string       { return string(r) }
func (r TextLine) Meta() string       { return "" }
func (r TextLine) Prefix() string     { return "" }
func (r TextLine) String() string     { return string(r) }

// ---- link line

func (r LinkLine) CoreType() LineType { return LinkType }
func (r LinkLine) Type() LineType     { return LinkType }
func (r LinkLine) Level() int         { return 0 }
func (r LinkLine) Data() string       { return r.link }
func (r LinkLine) Meta() string       { return r.name }
func (r LinkLine) Prefix() string     { return "=>" }
func (r LinkLine) String() string {
	if len(r.name) > 0 {
		return fmt.Sprintf("%s %s %s", r.Prefix(), r.link, r.name)
	}
	return fmt.Sprintf("%s %s", r.Prefix(), r.link)
}

// ---- preformat toggle line

func (r PreformatToggleLine) CoreType() LineType { return PreformatToggleType }
func (r PreformatToggleLine) Type() LineType     { return PreformatToggleType }
func (r PreformatToggleLine) Level() int         { return 0 }
func (r PreformatToggleLine) Data() string       { return string(r) }
func (r PreformatToggleLine) Meta() string       { return "" }
func (r PreformatToggleLine) Prefix() string     { return "```" }
func (r PreformatToggleLine) String() string     { return fmt.Sprintf("%s%s", r.Prefix(), r.Meta()) }

// ---- preformat line

func (r PreformatLine) CoreType() LineType { return PreformatType }
func (r PreformatLine) Type() LineType     { return PreformatType }
func (r PreformatLine) Level() int         { return 0 }
func (r PreformatLine) Data() string       { return string(r) }
func (r PreformatLine) Meta() string       { return "" }
func (r PreformatLine) Prefix() string     { return "" }
func (r PreformatLine) String() string     { return string(r) }

// ---- heading line

func (r HeadingLine) CoreType() LineType { return TextType }
func (r HeadingLine) Type() LineType     { return HeadingType }
func (r HeadingLine) Level() int         { return r.level }
func (r HeadingLine) Data() string       { return r.contents }
func (r HeadingLine) Meta() string       { return "" }
func (r HeadingLine) Prefix() string     { return strings.Repeat("#", r.level) }
func (r HeadingLine) String() string     { return fmt.Sprintf("%s %s", r.Prefix(), r.Data()) }

// ---- unordered list line

func (r UnorderedListLine) CoreType() LineType { return TextType }
func (r UnorderedListLine) Type() LineType     { return UnorderedListType }
func (r UnorderedListLine) Level() int         { return 0 }
func (r UnorderedListLine) Data() string       { return string(r) }
func (r UnorderedListLine) Meta() string       { return "" }
func (r UnorderedListLine) Prefix() string     { return "*" } // mostly present for CoreType consumers, use bullet points instead
func (r UnorderedListLine) String() string     { return fmt.Sprintf("%s %s", r.Prefix(), r.Data()) }

// ---- quote line

func (r QuoteLine) CoreType() LineType { return TextType }
func (r QuoteLine) Type() LineType     { return QuoteType }
func (r QuoteLine) Level() int         { return 0 }
func (r QuoteLine) Data() string       { return string(r) }
func (r QuoteLine) Meta() string       { return "" }
func (r QuoteLine) Prefix() string     { return ">" }
func (r QuoteLine) String() string     { return fmt.Sprintf("%s %s", r.Prefix(), r.Data()) }
