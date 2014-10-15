package figlet

import ()

type AsciiRender struct {
	fontMgr *FontManager
}

func NewAsciiRender() *AsciiRender {
	this := &AsciiRender{}

	this.fontMgr = NewFontManager()
	return this
}

func (this *AsciiRender) LoadFont(fontPath string) error {
	return this.fontMgr.LoadFont(fontPath)
}

func (this *AsciiRender) Render(asciiStr string, options interface{}) error {
	return nil
}
