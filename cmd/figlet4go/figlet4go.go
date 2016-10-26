package main

import (
	"fmt"
	"flag"
	"log"
	"errors"
	_"github.com/fatih/color"
	"github.com/probandula/figlet4go"
)

var (
	str *string = flag.String("str", "", "String to be converted with FIGlet")
	font *string = flag.String("font", "", "Font name to use")
	fontpath *string = flag.String("fontpath", "", "Font path to load fonts from")
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


	// Render the string
	renderStr, err := ascii.RenderOpts(*str, options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(renderStr)
}

// Validate if all required options are given
// flag.Parse() must be called before this
func validate() error {
	if *str == "" {
		return errors.New("No string given")
	}
	return nil
}