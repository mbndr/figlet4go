package figlet4go

import (
	"errors"
	"github.com/fatih/color"
	"strings"
)

type AsciiChar struct {
	// Slice with the lines of the Char
	Lines []string
	// Color of the char
	Color color.Attribute
}

func NewAsciiChar(font *font, char rune) (*AsciiChar, error) {
	// If not ascii, throw an error
	if char < 0 || char > 127 {
		return nil, errors.New("Not Ascii")
	}

	height := font.height
	beginRow := (int(char) - 32) * height

	lines := make([]string, height)

	for i := 0; i < height; i++ {
		row := font.fontSlice[beginRow+i]
		row = strings.Replace(row, "@", "", -1)
		row = strings.Replace(row, font.hardblank, " ", -1)
		lines[i] = row
	}

	return &AsciiChar{Lines: lines}, nil
}

// Return a line of the char as string with color if set
func (char *AsciiChar) GetLine(index int) string {
	if char.Color != 0 {
		colorFunc := color.New(char.Color).SprintFunc()
		return colorFunc(char.Lines[index])
	}
	return char.Lines[index]
}
