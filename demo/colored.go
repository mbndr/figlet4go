// Print a colored string with the 'standard' figlet font

package main

import (
	"fmt"
	"github.com/probandula/figlet4go"
	// This package is used to define the colors
	"github.com/fatih/color"
)

const str string = "Colored"

func main() {
	ascii := figlet4go.NewAsciiRender()

	// Add the colors to the options
	options := figlet4go.NewRenderOptions()
	options.FontColor = []color.Attribute{
		color.FgGreen,
		color.FgYellow,
		color.FgCyan,
	}

	// Use the RenderOpts instead of the Render method
	renderStr, err := ascii.RenderOpts(str, options)
	if err != nil {
		panic(err)
	}

	fmt.Print(renderStr)
}
