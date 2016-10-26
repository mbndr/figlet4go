// Print a string with another font

package main

import (
	"fmt"
	"github.com/probandula/figlet4go"
)

const str string = "otherfont"

func main() {
	ascii := figlet4go.NewAsciiRender()

	// Add the font name to the options
	options := figlet4go.NewRenderOptions()
	// This font is included so you don't have to specify a path with 'ascii.LoadFont(/path/to/fonts/)'
	options.FontName = "larry3d"

	// Use the RenderOpts instead of the Render method
	renderStr, err := ascii.RenderOpts(str, options)
	if err != nil {
		panic(err)
	}

	fmt.Print(renderStr)
}
