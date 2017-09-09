package main

import (
	"flag"
	"fmt"
	"github.com/mbndr/figlet4go"
	"log"
	"os"
	"strings"
)

var (
	str      *string = flag.String("str", "", "String to be converted with FIGlet")
	font     *string = flag.String("font", "", "Font name to use")
	fontpath *string = flag.String("fontpath", "", "Font path to load fonts from")
	colors   *string = flag.String("colors", "", "Character colors separated by ';'\n\tPossible colors: black, red, green, yellow, blue, magenta, cyan, white, or any hexcode (f.e. '885DBA')")
	parser   *string = flag.String("parser", "terminal", "Parser to use\tPossible parsers: terminal, html")
	file     *string = flag.String("file", "", "File to write to")
)

func main() {
	// Parse the flags
	flag.Parse()

	// Validate and log the error
	validate()

	// Create objects
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()

	// Load fonts
	if *fontpath != "" {
		ascii.LoadFont(*fontpath)
	}

	// Set the font
	options.FontName = *font

	// Set the parser
	p, err := figlet4go.GetParser(*parser)
	if err != nil {
		p, _ = figlet4go.GetParser("terminal")
	}
	options.Parser = *p

	// Set colors
	if *colors != "" {
		options.FontColor = getColorSlice(*colors)
	}

	// Render the string
	renderStr, err := ascii.RenderOpts(*str, options)
	if err != nil {
		log.Fatal(err)
	}

	// Write to file if given
	if *file != "" {
		// Create file
		f, err := os.Create(*file)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		// Write to file
		b, err := f.WriteString(renderStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Wrote %d bytes to %s\n", b, *file)
		return
	}

	// Default is printing
	fmt.Print(renderStr)
}

// Get a slice with colors to give to the RenderOptions
// Splits the given string with the separator ";"
func getColorSlice(colorStr string) []figlet4go.Color {

	givenColors := strings.Split(colorStr, ";")

	colors := make([]figlet4go.Color, len(givenColors))

	for i, c := range givenColors {
		switch c {
		case "black":
			colors[i] = figlet4go.ColorBlack
		case "red":
			colors[i] = figlet4go.ColorRed
		case "green":
			colors[i] = figlet4go.ColorGreen
		case "yellow":
			colors[i] = figlet4go.ColorYellow
		case "blue":
			colors[i] = figlet4go.ColorBlue
		case "magenta":
			colors[i] = figlet4go.ColorMagenta
		case "cyan":
			colors[i] = figlet4go.ColorCyan
		case "white":
			colors[i] = figlet4go.ColorWhite
		default:
			// Try to parse the TrueColor from the string
			color, err := figlet4go.NewTrueColorFromHexString(c)
			if err != nil {
				log.Fatal(err)
			}
			colors[i] = color
		}
	}

	return colors
}

// Validate if all required options are given
// flag.Parse() must be called before this
func validate() {
	if *str == "" {
		flag.Usage()
		os.Exit(1)
	}
}
