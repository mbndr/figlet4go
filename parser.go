package figlet4go

// Parser stores some output specific stuff
type Parser struct {
	// Name used for switching in colors.go
	Name     string
	// Parser prefix
	Prefix   string
	// Parser suffix
	Suffix   string
	// Newline representation
	NewLine  string
	// Things to be replaced (f.e. \n to <br>)
	Replaces map[string]string
}

var (
	// Default terminal parser
	ParserTerminal Parser = Parser{"terminal", "", "", "\n", nil}
	// Parser for HTML code
	ParserHTML Parser = Parser{"html", "<code>", "</code>", "<br>", map[string]string{" ": "&nbsp;"}}
)
