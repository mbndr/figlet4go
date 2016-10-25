package figlet4go

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const defaultFontName string = "standard"

// Represents a single font
type font struct {
	hardblank string
	height    int
	fontSlice []string
}

// Holds the fonts
type fontManager struct {
	// Font libraries
	fontLib map[string]*font
	// Font name to path
	fontList map[string]string
}

// Create new fontmanager
func newFontManager() *fontManager {
	this := &fontManager{}

	this.fontLib = make(map[string]*font)
	this.fontList = make(map[string]string)
	this.loadBuildInFont()

	return this
}

// Load all font *.flf files in the fontPath recursivly
func (this *fontManager) loadFont(fontPath string) error {

	return filepath.Walk(fontPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".flf") {
			return nil
		}

		fontName := strings.TrimSuffix(info.Name(), ".flf")
		this.fontList[fontName] = path
		return nil
	})
}

// Load the default font
func (this *fontManager) loadBuildInFont() error {
	fontStr, err := Asset("standard.flf")
	if err != nil {
		panic(err)
	}

	font, err := this.parseFontContent(string(fontStr))
	if err != nil {
		return err
	}
	this.fontLib[defaultFontName] = font
	return nil
}

// Load a font from disk
func (this *fontManager) loadDiskFont(fontName string) error {

	fontFilePath, ok := this.fontList[fontName]
	if !ok {
		return errors.New("FontName Not Found.")
	}

	// read full file content
	fileBuf, err := ioutil.ReadFile(fontFilePath)
	if err != nil {
		return err
	}

	font, err := this.parseFontContent(string(fileBuf))
	if err != nil {
		return err
	}

	this.fontLib[fontName] = font
	return nil
}

// Parse a font from a content string
func (this *fontManager) parseFontContent(cont string) (*font, error) {
	lines := strings.Split(cont, "\n")
	if len(lines) < 1 {
		return nil, errors.New("font content error")
	}

	// flf2a$ 7 5 16 -1 12
	// Fender by Scooter 8/94 (jkratten@law.georgetown.edu)
	//
	// Explanation of first line:
	// flf2 - "magic number" for file identification
	// a    - should always be `a', for now
	// $    - the "hardblank" -- prints as a blank, but can't be smushed
	// 7    - height of a character
	// 5    - height of a character, not including descenders
	// 10   - max line length (excluding comment lines) + a fudge factor
	// -1   - default smushmode for this font (like "-m 15" on command line)
	// 12   - number of comment lines

	header := strings.Split(lines[0], " ")

	font := &font{}
	font.hardblank = header[0][len(header[0])-1:]
	font.height, _ = strconv.Atoi(header[1])

	commentEndLine, _ := strconv.Atoi(header[5])
	font.fontSlice = lines[commentEndLine+1:]

	return font, nil
}

// Get a font by name
func (this *fontManager) getFont(fontName string) (*font, error) {
	font, ok := this.fontLib[fontName]
	if !ok {
		err := this.loadDiskFont(fontName)
		if err != nil {
			font, _ := this.fontLib[defaultFontName]
			return font, nil
		}
	}
	font, _ = this.fontLib[fontName]
	return font, nil
}
