package figlet4go

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// Escape char
const escape string = "\x1b"

// Terminal AnsiColors
var (
	ColorBlack   AnsiColor = AnsiColor{30}
	ColorRed     AnsiColor = AnsiColor{31}
	ColorGreen   AnsiColor = AnsiColor{32}
	ColorYellow  AnsiColor = AnsiColor{33}
	ColorBlue    AnsiColor = AnsiColor{34}
	ColorMagenta AnsiColor = AnsiColor{35}
	ColorCyan    AnsiColor = AnsiColor{36}
	ColorWhite   AnsiColor = AnsiColor{37}
)

// TrueColor lookalikes for displaying AnsiColor f.e. with the HTML parser
// Colors based on http://clrs.cc/
// "TrueColorForAnsiColor"
var tcfac map[AnsiColor]TrueColor = map[AnsiColor]TrueColor{
	ColorBlack:   {0, 0, 0},
	ColorRed:     {255, 65, 54},
	ColorGreen:   {149, 189, 64},
	ColorYellow:  {255, 220, 0},
	ColorBlue:    {0, 116, 217},
	ColorMagenta: {177, 13, 201},
	ColorCyan:    {105, 206, 245},
	ColorWhite:   {255, 255, 255},
}

// Color has a pre- and a suffix
type Color interface {
	getPrefix(p Parser) string
	getSuffix(p Parser) string
}

// AnsiColor representation
type AnsiColor struct {
	code int
}

// TrueColor with rgb Attributes
type TrueColor struct {
	r int
	g int
	b int
}

// Prefix for ansi color
func (tc TrueColor) getPrefix(p Parser) string {
	switch p.Name {

	case "terminal":
		return fmt.Sprintf("%v[38;2;%d;%d;%dm", escape, tc.r, tc.g, tc.b)

	case "html":
		return fmt.Sprintf("<span style='color: rgb(%d,%d,%d);'>", tc.r, tc.g, tc.b)
	}

	return ""
}

// Suffix for ansi color
func (tc TrueColor) getSuffix(p Parser) string {
	switch p.Name {

	case "terminal":
		return fmt.Sprintf("%v[0m", escape)

	case "html":
		return "</span>"
	}

	return ""
}

// NewTrueColorFromHexString returns a Truecolor object based on a hexadezimal string
func NewTrueColorFromHexString(c string) (*TrueColor, error) {
	rgb, err := hex.DecodeString(c)
	if err != nil {
		return nil, errors.New("Invalid color given (" + c + ")")
	}

	return &TrueColor{
		int(rgb[0]),
		int(rgb[1]),
		int(rgb[2]),
	}, nil
}

// Prefix for ansi color
func (ac AnsiColor) getPrefix(p Parser) string {
	switch p.Name {

	case "terminal":
		return fmt.Sprintf("%v[0;%dm", escape, ac.code)

	case "html":
		// Get the TrueColor for the AnsiColor
		tc := tcfac[ac]
		return tc.getPrefix(p)
	}

	return ""

}

// Suffix for ansi color
func (ac AnsiColor) getSuffix(p Parser) string {
	switch p.Name {

	case "terminal":
		return fmt.Sprintf("%v[0m", escape)

	case "html":
		// Get the TrueColor for the AnsiColor
		tc := tcfac[ac]
		return tc.getSuffix(p)
	}

	return ""
}
