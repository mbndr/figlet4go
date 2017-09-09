# FIGlet for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/mbndr/figlet4go)](https://goreportcard.com/report/github.com/mbndr/figlet4go)

`figlet4go` is a go library which is a port of [FIGlet](http://www.figlet.org/) to Golang.  
With `figlet4go` it's easy to create **ascii text banners** in the command-line or with the given api.

![screenshot](./screenshot/figlet4go.png)


This Repository used to be a fork of [getwe/figlet4go](https://github.com/getwe/figlet4go), but I changed so much that it's not compatible anymore

## Installation

```
$ go get -u github.com/mbndr/figlet4go/...
```

## Usage

### Command-line
You can use the `figlet4go` command in the command-line.  
For example (generates the banner on top):
```bash
$ figlet4go -str "figlet4go" -font "larry3d" -colors "green;FF9900;cyan"
```
For a usage instruction read the commands usage with `figlet4go -h`.

### Basic
You have to create a renderer (`ascii`) and let it render the desired string through the `Render` method. After that you can simply print the returned string.
```go
import "github.com/mbndr/figlet4go"

// ...

ascii := figlet4go.NewAsciiRender()

// The underscore would be an error
renderStr, _ := ascii.Render("Hello World")
fmt.Print(renderStr)
```

### Colored
The colors given in the `[]figlet4go.Color` slice are repeating if the string is longer than the slice. You have to call the `RenderOpts` instead of the `Render` method to give the Renderer the Options.  
If you use a `TrueColor` color, you have to ensure that your [terminal supports](https://gist.github.com/XVilka/8346728/) it.  
If you use a `AnsiColor` with an `TrueColor` only parser (f.e. `ParserHTML`), `TrueColor` objects are automatically generated.
```go
import "github.com/mbndr/figlet4go"

// ...

ascii := figlet4go.NewAsciiRender()

// Adding the colors to RenderOptions
options := figlet4go.NewRenderOptions()
options.FontColor = []figlet4go.Color{
	// Colors can be given by default ansi color codes...
	figlet4go.ColorGreen,
	figlet4go.ColorYellow,
	figlet4go.ColorCyan,
	// ...or by an hex string...
	figlet4go.NewTrueColorFromHexString("885DBA"),
	// ...or by an TrueColor object with rgb values
	figlet4go.TrueColor{136, 93, 186},
}

renderStr, _ := ascii.RenderOpts("Hello Colors", options)
fmt.Print(renderStr)
```

### Other font
If you want to use another font, you have to specify the name of the font as in this example.  
Is the font you want to use not [included](#builtin) you have to load the font manually with the `LoadFont` method. This method will walk the path recursively and load all `.flf` files.
```go
import "github.com/mbndr/figlet4go"

// ...

ascii := figlet4go.NewAsciiRender()

options := figlet4go.NewRenderOptions()
options.FontName = "larry3d"

// If 'larry3d' wouldn't be included you would have to load your .flf files like that:
ascii.LoadFont("/path/to/fonts/")

renderStr, _ := ascii.RenderOpts("Hello Fonts", options)
fmt.Print(renderStr)
```

### Other parser
A Parser can be set through the `GetParser` function with a valid key
```go
import "github.com/mbndr/figlet4go"

// ...

ascii := figlet4go.NewAsciiRender()

options := figlet4go.NewRenderOptions()
p, _ := figlet4go.GetParser("html")
options.Parser = *p

renderStr, _ := ascii.RenderOpts("Hello Fonts", options)
fmt.Print(renderStr)
```

## Parsers
There a currently these Parsers available:

| Parser | What does it do?                                                     |
| --------- | ------                                                     |
| ParserTerminal  | Parses the result directly |
| ParserHTML   | Parses a pasteable `<code>` html block  |

## Fonts

### Builtin
The builtin fonts are built into the `bindata.go` file with the tool [go-bindata](https://github.com/jteeuwen/go-bindata).  
The bash script for building the default font is stored in `tools/` (`go-bindata` must be installed).

The default font is `standard`. These are the builtin fonts:

| Font name | Source                                                     |
| --------- | ------                                                     |
| standard  | http://www.figlet.org/fontdb_example.cgi?font=standard.flf |
| larry3d   | http://www.figlet.org/fontdb_example.cgi?font=larry3d.flf  |

### Other fonts
Other fonts can mainly be found on [figlet](http://www.figlet.org). You have to load them as in [this example](#other-font).

## Todo
- [ ] Tests
- [ ] automatically the perfect char margin
- [ ] Linebreak possible?
- [ ] Pointer-Value standarization
- [ ] Parser as interface
- [x] Cli client
- [x] Colors in the cli client
- [x] No dependencies (fatih/color)
- [x] Truecolor support
- [x] More parsers (HTML)
- [x] Better parsers (maybe stored in a map)
- [x] Writer choosing for writing to file
