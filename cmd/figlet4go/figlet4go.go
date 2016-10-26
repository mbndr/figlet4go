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
	colors   *string = flag.String("colors", "", "Character colors seperated by ';'\n\tPossible colors: black, red, green, yellow, blue, magenta, cyan, white")
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

	fmt.Println(renderStr)
}

// Get a slice with colors to give to the RenderOptions
// Splits the given string with the seperator ";"
func getColorSlice(colorStr string) []color.Attribute {

	givenColors := strings.Split(colorStr, ";")

	colors := make([]color.Attribute, len(givenColors))

	for i, c := range givenColors {
		if c == "black" {
			colors[i] = color.FgBlack
			continue
		}
		if c == "red" {
			colors[i] = color.FgRed
			continue
		}
		if c == "green" {
			colors[i] = color.FgGreen
			continue
		}
		if c == "yellow" {
			colors[i] = color.FgYellow
			continue
		}
		if c == "blue" {
			colors[i] = color.FgBlue
			continue
		}
		if c == "magenta" {
			colors[i] = color.FgMagenta
			continue
		}
		if c == "cyan" {
			colors[i] = color.FgCyan
			continue
		}
		if c == "white" {
			colors[i] = color.FgWhite
			continue
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
