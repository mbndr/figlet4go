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
package figlet4go


// Represents a single font
type font struct {
	// Hardblank symbol
	hardblank string
	// Height of one char
	height    int
	//
	fontSlice []string
}

func (f *font) getCharSlice() []string {
	// TODO here will be the logic of NewAsciiChar
	return []string{} 
}