package figlet4go

// Explanation of the .flf file header
// THE HEADER LINE
//
// The header line gives information about the FIGfont.  Here is an example
// showing the names of all parameters:
//
//           flf2a$ 6 5 20 15 3 0 143 229    NOTE: The first five characters in
//             |  | | | |  |  | |  |   |     the entire file must be "flf2a".
//            /  /  | | |  |  | |  |   \
//   Signature  /  /  | |  |  | |   \   Codetag_Count
//     Hardblank  /  /  |  |  |  \   Full_Layout*
//          Height  /   |  |   \  Print_Direction
//          Baseline   /    \   Comment_Lines
//           Max_Length      Old_Layout*
//
//   * The two layout parameters are closely related and fairly complex.
//       (See "INTERPRETATION OF LAYOUT PARAMETERS".)
//

import (
	"strings"
)

// Represents a single font
type font struct {
	// Hardblank symbol
	hardblank string
	// Height of one char
	height int
	// A string for each line of the char
	fontSlice []string
}

// Get a slice of strings containing the chars lines
func (f *font) getCharSlice(char rune) []string {

	height := f.height
	beginRow := (int(char) - 32) * height

	lines := make([]string, height)

	// Get the char lines of the char
	for i := 0; i < height; i++ {
		row := f.fontSlice[beginRow+i]
		row = strings.Replace(row, "@", "", -1)
		row = strings.Replace(row, f.hardblank, " ", -1)
		lines[i] = row
	}

	return lines
}
