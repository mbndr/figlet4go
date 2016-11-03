package figlet4go

import (
	"errors"
	"strings"
)

// Represents a single ascii character
type asciiChar struct {
	// Slice with the lines of the Char
	Lines []string
	// Color of the char
	Color Color
}

// Creates a new ascii character
func newAsciiChar(font *font, char rune) (*asciiChar, error) {
	// If not ascii, throw an error
	if char < 0 || char > 127 {
		return nil, errors.New("Not Ascii")
	}

	// Get the font's representation of the char
	lines := font.getCharSlice(char)

	return &asciiChar{Lines: lines}, nil
}

// Return a line of the char as string with color if set
func (char *asciiChar) GetLine(index int, p Parser) string {
	prefix := ""
	suffix := ""

	line := handleReplaces(char.Lines[index], p)

	if char.Color != nil {
		prefix = char.Color.getPrefix(p)
		suffix = char.Color.getSuffix(p)
	}

	return prefix + line + suffix
}

// Replace all parser specific things
func handleReplaces(str string, p Parser) string {
	if p.Replaces == nil {
		return str
	}
	// Replace for each entry
	for old, new := range p.Replaces {
		str = strings.Replace(str, old, new, -1)
	}
	return str
}
