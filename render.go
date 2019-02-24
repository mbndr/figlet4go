package figlet4go

// RenderOptions are used to set color or maybe future
// options to the AsciiRenderer
type RenderOptions struct {
	// Name of the used font
	FontName string
	// Colors of the font
	FontColor []Color
	// Parser
	Parser Parser
}

// NewRenderOptions creates new RenderOptions
// Sets the default font name
func NewRenderOptions() *RenderOptions {
	p, _ := GetParser("terminal")
	return &RenderOptions{
		FontName: defaultFont,
		Parser:   *p,
	}
}

// AsciiRender is the wrapper to render a string
type AsciiRender struct {
	// FontManager to store all the fonts
	fontMgr *fontManager
}

// NewAsciiRender creates a new AsciiRender
func NewAsciiRender() *AsciiRender {
	return &AsciiRender{
		fontMgr: newFontManager(),
	}
}

// LoadFont loads all *.flf font files recursively in a path
func (ar *AsciiRender) LoadFont(fontPath string) error {
	return ar.fontMgr.loadFontList(fontPath)
}

// LoadBinDataFont loads provided font binary
func (ar *AsciiRender) LoadBindataFont(fontBinary []byte, fontName string) error {
	return ar.fontMgr.loadBindataFont(fontBinary, fontName)
}

// Render renders a string with the default options
// Calls the RenderOpts method with a new RenderOptions object
func (ar *AsciiRender) Render(str string) (string, error) {
	return ar.RenderOpts(str, NewRenderOptions())
}

// RenderOpts renders a string with special RenderOptions
// Can be called from the user (if options wished) or the above Render method
// Contains the whole rendering logic
func (ar *AsciiRender) RenderOpts(str string, opt *RenderOptions) (string, error) {
	// Should the text be colored
	colored := len(opt.FontColor) > 0

	// Load the font
	font := ar.fontMgr.getFont(opt.FontName)

	// Slice holding the chars
	chars := []*asciiChar{}

	// Index of the current color
	curColorIndex := 0

	// Foreach char create the ascii char
	for _, char := range str {
		// AsciiChar
		asciiChar, err := newAsciiChar(font, char)
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

	result += opt.Parser.Prefix

	// Foreach line of the font height
	for curLine := 0; curLine < font.height; curLine++ {
		// Add the current line of the char to the result
		for i := range chars {
			result += chars[i].GetLine(curLine, opt.Parser)
		}
		// A new line at the end
		result += opt.Parser.NewLine
	}

	result += opt.Parser.Suffix

	return result, nil
}
