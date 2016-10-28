package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/probandula/figlet4go"
	"log"
	"strings"
)

var (
	str      *string = flag.String("str", "", "String to be converted with FIGlet")
	font     *string = flag.String("font", "", "Font name to use")
	fontpath *string = flag.String("fontpath", "", "Font path to load fonts from")
	colors   *string = flag.String("colors", "", "Character colors separated by ';'\n\tPossible colors: black, red, green, yellow, blue, magenta, cyan, white")
)

func main() {
	// Parse the flags
	flag.Parse()

	// Validate and log the error
	err := validate()
	if err != nil {
		log.Fatal(err)
	}

	// Create objects
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()

	// Load fonts
	if *fontpath != "" {
		ascii.LoadFont(*fontpath)
	}

	// Set the font
	options.FontName = *font

	// Set colors
	if *colors != "" {
		options.FontColor = getColorSlice(*colors)
	}

	// Render the string
	renderStr, err := ascii.RenderOpts(*str, options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(renderStr)
}

// Get a slice with colors to give to the RenderOptions
// Splits the given string with the separator ";"
func getColorSlice(colorStr string) []color.Attribute {

	givenColors := strings.Split(colorStr, ";")

	colors := make([]color.Attribute, len(givenColors))

	for i, c := range givenColors {
		switch c {
		case "black":
			colors[i] = color.FgBlack
		case "red":
			colors[i] = color.FgRed
		case "green":
			colors[i] = color.FgGreen
		case "yellow":
			colors[i] = color.FgYellow
		case "blue":
			colors[i] = color.FgBlue
		case "magenta":
			colors[i] = color.FgMagenta
		case "cyan":
			colors[i] = color.FgCyan
		case "white":
			colors[i] = color.FgWhite
		}
	}

	return colors
}

// Validate if all required options are given
// flag.Parse() must be called before this
func validate() error {
	if *str == "" {
		return errors.New("No string given")
	}
	return nil
}
