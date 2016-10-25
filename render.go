package figlet4go

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"strings"
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
func (ar *AsciiRender) Render(str string) (string, error) {
	return ar.render(str, NewRenderOptions())
}

// Render a string with special RenderOptions
func (ar *AsciiRender) RenderOpts(str string, opts *RenderOptions) (string, error) {
	return ar.render(str, opts)
}



func (this *AsciiRender) convertChar(font *font, char rune) ([]string, error) {

	if char < 0 || char > 127 {
		return nil, errors.New("Not Ascii")
	}

	height := font.height
	begintRow := (int(char) - 32) * height

	word := make([]string, height, height)

	for i := 0; i < height; i++ {
		row := font.fontSlice[begintRow+i]
		row = strings.Replace(row, "@", "", -1)
		row = strings.Replace(row, font.hardblank, " ", -1)
		word[i] = row
	}

	return word, nil
}

func (this *AsciiRender) render(asciiStr string, opt *RenderOptions) (string, error) {

	font, _ := this.fontMgr.getFont(opt.FontName)

	wordlist := make([][]string, 0)
	for _, char := range asciiStr {
		word, err := this.convertChar(font, char)
		if err != nil {
			return "", err
		}
		wordlist = append(wordlist, word)
	}

	result := ""

	wordColorFunc := make([]func(a ...interface{}) string, len(wordlist))
	for i, _ := range wordColorFunc {
		if i < len(opt.FontColor) {
			wordColorFunc[i] = color.New(opt.FontColor[i]).SprintFunc()
		} else {
			wordColorFunc[i] = fmt.Sprint
		}
	}

	for i := 0; i < font.height; i++ {
		for j := 0; j < len(wordlist); j++ {
			result += wordColorFunc[j]((wordlist[j][i]))
		}
		result += "\n"
	}
	return result, nil
}
