package figlet4go

import (
	"errors"
	"github.com/fatih/color"
)

// Represents a single ascii character
type asciiChar struct {
	// Slice with the lines of the Char
	Lines []string
	// Color of the char
	Color color.Attribute
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
	if char.Color != 0 {
		colorFunc := color.New(char.Color).SprintFunc()
		return colorFunc(char.Lines[index])
	}
	return char.Lines[index]
}
