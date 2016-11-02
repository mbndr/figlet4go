package figlet4go

import (
	"errors"
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
func (char *asciiChar) GetLine(index int) string {
	prefix := ""
	suffix := ""

	if char.Color != nil {
		prefix = char.Color.getPrefix()
		suffix = char.Color.getSuffix()
	}

	return prefix + char.Lines[index] + suffix
}
