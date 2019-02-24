package figlet4go

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Default font if no other valid given
const defaultFont string = "standard"

// Extension of a font file
const extension string = "flf"

// Builtin fonts to load
var defaultFonts []string = []string{
	"standard",
	"larry3d",
}

// Holds the available fonts
type fontManager struct {
	// The already read fonts
	fontLib map[string]*font
	// The in given pathes found fonts
	fontList map[string]string
}

// Create a new fontmanager
// Initializes the fontManager,
// loads the builtin fonts and returns it
func newFontManager() *fontManager {
	fm := &fontManager{
		fontLib:  make(map[string]*font),
		fontList: make(map[string]string),
	}
	fm.loadBuildInFont()
	return fm
}

// Get a font by name
// Default font if no other font could be loaded
func (fm *fontManager) getFont(fontName string) *font {
	// Get the font from the fontLib
	_, ok := fm.fontLib[fontName]
	// Font not found
	if !ok {
		// Try to load it from loaded fontList
		err := fm.loadDiskFont(fontName)
		// Font also not found here, use the default font
		if err != nil {
			fontName = defaultFont
		}
	}

	return fm.fontLib[fontName]
}

// Loads all .flf files recursively in the fontPath path
// Saves the found font files in a map with the name as the key
// and the path as the value. Doesn't load them at this point
// for performance. Called in the AsciiRenderer
func (fm *fontManager) loadFontList(fontPath string) error {
	// Walk through the path
	return filepath.Walk(fontPath, func(path string, info os.FileInfo, err error) error {
		// Return an error if occurred
		if err != nil {
			return err
		}
		// If the current item is a directory or has not the correct suffix
		if info.IsDir() || !strings.HasSuffix(info.Name(), "."+extension) {
			return nil
		}
		// Extract the font name
		fontName := strings.TrimSuffix(info.Name(), "."+extension)
		// Save the font to the list
		fm.fontList[fontName] = path

		return nil
	})
}

// Load a font from disk
// The font must be registered in the fontList
func (fm *fontManager) loadDiskFont(fontName string) error {
	// Get the fontpath
	path, ok := fm.fontList[fontName]
	// Font is not registered
	if !ok {
		return errors.New("Font Not Found: " + fontName)
	}

	// Read file contents
	fontStr, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Parse the file contents
	font, err := parseFontContent(string(fontStr))
	if err != nil {
		return err
	}

	// Register the font object in the fontLib
	fm.fontLib[fontName] = font

	return nil
}

// Load the builtin fonts from the bindata.go file
// Load all fonts specified on top (defaultFonts)
func (fm *fontManager) loadBuildInFont() error {

	// Load each default font
	for _, name := range defaultFonts {
		// Get Contents
		fontStr, err := Asset(name + "." + extension)
		if err != nil {
			return err
		}
		// Load the font
		err = fm.loadBindataFont(fontStr, name)
		if err != nil {
			return err
		}
	}

	return nil
}

// Load a bindata font
func (fm *fontManager) loadBindataFont(fontBinary []byte, fontName string) error {

	// Get the font
	font, err := parseFontContent(string(fontBinary))
	if err != nil {
		return err
	}
	// Register the font object in the fontLib
	fm.fontLib[fontName] = font

	return nil
}

// Parse a font from a content string
// Used to load fonts from disk and the builtin fonts
func parseFontContent(cont string) (*font, error) {
	// Get all lines
	lines := strings.Split(cont, "\n")

	if len(lines) < 1 {
		return nil, errors.New("Font content error")
	}

	// Get the header metadata
	header := strings.Split(lines[0], " ")

	// Line end of the comment
	commentEndLine, _ := strconv.Atoi(header[5])

	// Char height
	height, _ := strconv.Atoi(header[1])

	// Initialize the font
	font := &font{
		hardblank: header[0][len(header[0])-1:],
		height:    height,
		fontSlice: lines[commentEndLine+1:],
	}

	return font, nil
}
