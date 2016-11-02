package figlet4go

import (
	"fmt"
	"errors"
	"encoding/hex"
)

// Escape char
const escape string = "\x1b"

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

// Every color has a pre- and a suffix
type Color interface {
	getPrefix() string
	getSuffix() string
}

// Ansi color
type AnsiColor struct {
	code int
}

// Truecolor with rgb Attributes
type TrueColor struct {
	r int
	g int
	b int
}

// Prefix for ansi color
func (tc TrueColor) getPrefix() string {
	return fmt.Sprintf("%v[38;2;%d;%d;%dm", escape, tc.r, tc.g, tc.b)
}

// Suffix for ansi color
func (tc TrueColor) getSuffix() string {
	return fmt.Sprintf("%v[0m", escape)
}

func GetTrueColorFromHexString(c string) (*TrueColor, error) {
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
func (ac AnsiColor) getPrefix() string {
	return fmt.Sprintf("%v[0;%dm", escape, ac.code)
}

// Suffix for ansi color
func (ac AnsiColor) getSuffix() string {
	return fmt.Sprintf("%v[0m", escape)
}
