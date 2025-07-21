package internal

import "github.com/fatih/color"

var (
	white   = color.New(color.FgWhite).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	black   = color.New(color.FgBlack).SprintFunc()

	hwhite   = color.New(color.FgHiWhite).SprintFunc()
	hyellow  = color.New(color.FgHiYellow).SprintFunc()
	hgreen   = color.New(color.FgHiGreen).SprintFunc()
	hblue    = color.New(color.FgHiBlue).SprintFunc()
	hmagenta = color.New(color.FgHiMagenta).SprintFunc()
	hblack   = color.New(color.FgHiBlack).SprintFunc()
)

var colors = map[string]func(...interface{}) string{
	"white":    white,
	"yellow":   yellow,
	"green":    green,
	"blue":     blue,
	"magenta":  magenta,
	"black":    black,
	"hwhite":   hwhite,
	"hyellow":  hyellow,
	"hgreen":   hgreen,
	"hblue":    hblue,
	"hmagenta": hmagenta,
	"hblack":   hblack,
}
