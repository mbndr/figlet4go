# FIGlet for Go

**Currently in Development.  
Sould work but will be improved (add demos, better font adding, maybe better performance, more default fonts)**

A port of [figlet](http://www.figlet.org/) to golang and fork of [getwe/figlet4go](https://github.com/getwe/figlet4go).

![screenshot](./screenshot/figlet4go.png)

## Installation

```
go get -u github.com/probandula/figlet4go
```

## Usage
```go
// Create the renderer
ascii := figlet4go.NewAsciiRender()

// Optional: Add color to the letters (https://github.com/fatih/color needed)

// Render and print the string
renderStr, _ := ascii.Render("Hello World")
fmt.Print(renderStr)
```

## Default font
The default font is built into the `bindata.go` file with the tool [go-bindata](https://github.com/jteeuwen/go-bindata).  
The bash script for building the default font is stored in `tools/` (`go-bindata` must be installed).

## Use the demo
There are [demo](https://github.com/probandula/figlet4go/blob/master/demo) programs for trying out the library.  
To run them, `cd` into the `demo/` directory and run `go run [filename]` on any program you want to run.