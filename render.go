package figlet4go

import (
	"github.com/fatih/color"
)

// RenderOptions are used to set color or maybe future
// options to the AsciiRenderer
type RenderOptions struct {
	// Name of the used font
	FontName string
	// Colors of the font
	FontColor []color.Attribute
}

// Create new RenderOptions
// Sets the default font name
func NewRenderOptions() *RenderOptions {
	return &RenderOptions{
		FontName: defaultFontName,
	}
}

// AsciiRender is the wrapper to render a string
type AsciiRender struct {
	// FontManager to store all the fonts
	fontMgr *fontManager
}

// Create a new AsciiRender
func NewAsciiRender() *AsciiRender {
	return &AsciiRender{
		fontMgr: newFontManager(),
	}
}

// Loading all *.flf font files recursively in a path
func (ar *AsciiRender) LoadFont(fontPath string) error {
	return ar.fontMgr.loadFont(fontPath)
}

// Render a string with the default options
// Calls the RenderOpts method with a new RenderOptions object
func (ar *AsciiRender) Render(str string) (string, error) {
	return ar.RenderOpts(str, NewRenderOptions())
}

// Render a string with special RenderOptions
// Can be called from the user (if options wished) or the above Render method
// Contains the whole rendering logic
func (ar *AsciiRender) RenderOpts(str string, opt *RenderOptions) (string, error) {
	// Should the text be colored
	colored := len(opt.FontColor) > 0

	// Load the font
	font, err := ar.fontMgr.getFont(opt.FontName)
	if err != nil {
		return "", err
	}

	// Slice holding the chars
	chars := []*AsciiChar{}

	// Index of the current color
	curColorIndex := 0

	// Foreach char create the ascii char
	for _, char := range str {
		// AsciiChar
		asciiChar, err := NewAsciiChar(font, char)
		if err != nil {
			return "", err
		}

		// Set color if given
		if colored {
			// Start colors from beginning if length is reached
			if curColorIndex == len(opt.FontColor) {
				curColorIndex = 0
			}
			// Assign color and increment the index
			asciiChar.Color = opt.FontColor[curColorIndex]
			curColorIndex++
		}

		// Append the char to the chars slice
		chars = append(chars, asciiChar)
	}

	// Result which will be returned
	result := ""

	// Foreach line of the font height
	for curLine := 0; curLine < font.height; curLine++ {
		// Add the current line of the char to the result
		for i, _ := range chars {
			result += chars[i].GetLine(curLine)
		}
		// A new line at the end
		result += "\n"
	}

	return result, nil
}
