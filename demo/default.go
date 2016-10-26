// Print a simple string with the 'standard' figlet font

package main

import (
	"fmt"
	"github.com/probandula/figlet4go"
)

// String to be printed
const str string = "Default"

func main() {
	// Create the renderer	
	ascii := figlet4go.NewAsciiRender()

	// Render the string
	renderStr, err := ascii.Render(str)
	if err != nil {
		panic(err)
	}

	// Print the string
	fmt.Print(renderStr)
}
