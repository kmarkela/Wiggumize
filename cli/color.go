package cli

import (
	"github.com/fatih/color"
)

// const (
// 	Reset        color.Attribute = iota
// 	Bold
// 	Faint
// 	Italic
// 	Underline
// 	BlinkSlow
// 	BlinkRapid
// 	ReverseVideo
// 	Concealed
// 	CrossedOut
// )

type ColorPrinter struct {
	Red     *color.Color
	Green   *color.Color
	Yellow  *color.Color
	Blue    *color.Color
	Magenta *color.Color
	Cyan    *color.Color
	White   *color.Color
}

func NewColorPrinter() *ColorPrinter {
	return &ColorPrinter{
		Red:     color.New(color.FgRed),
		Green:   color.New(color.FgGreen),
		Yellow:  color.New(color.FgYellow),
		Blue:    color.New(color.FgBlue),
		Magenta: color.New(color.FgMagenta),
		Cyan:    color.New(color.FgCyan),
		White:   color.New(color.FgWhite),
	}
}

func (cp ColorPrinter) AddAttributeString(cc *color.Color, att string) {
	switch att {
	case "Reset":
		cc.Add(color.Reset)
	case "Bold":
		cc.Add(color.Bold)
	case "Faint":
		cc.Add(color.Faint)
	case "Italic":
		cc.Add(color.Italic)
	case "Underline":
		cc.Add(color.Underline)
	case "BlinkSlow":
		cc.Add(color.BlinkSlow)
	case "BlinkRapid":
		cc.Add(color.BlinkRapid)
	case "ReverseVideo":
		cc.Add(color.ReverseVideo)
	case "Concealed":
		cc.Add(color.Concealed)
	case "CrossedOut":
		cc.Add(color.CrossedOut)
	}
}
