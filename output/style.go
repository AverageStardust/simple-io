package output

import (
	"fmt"
	"strconv"
	"strings"
)

type simpleColor int
type trueColor [3]byte
type textWeight int

const (
	defaultColor simpleColor = 0
	isTrueColor  simpleColor = 1

	// Foreground
	blackFg simpleColor = iota + 28 // 30..37
	redFg
	greenFg
	yellowFg
	blueFg
	magentaFg
	cyanFg
	whiteFg

	// Background
	blackBg simpleColor = iota + 30 // 40..47
	redBg
	greenBg
	yellowBg
	blueBg
	magentaBg
	cyanBg
	whiteBg
)

const (
	regularWeight textWeight = iota
	boldWeight
	dimWeight
)

type Style struct {
	fgColor       simpleColor
	fgTrueColor   trueColor
	bgColor       simpleColor
	bgTrueColor   trueColor
	weight        textWeight
	italic        bool
	underline     bool
	strikethrough bool
}

func (style *Style) Black() *Style {
	style.fgColor = blackFg
	return style
}

func (style *Style) Red() *Style {
	style.fgColor = redFg
	return style
}

func (style *Style) Green() *Style {
	style.fgColor = greenFg
	return style
}

func (style *Style) Yellow() *Style {
	style.fgColor = yellowFg
	return style
}

func (style *Style) Blue() *Style {
	style.fgColor = blueFg
	return style
}

func (style *Style) Magenta() *Style {
	style.fgColor = magentaFg
	return style
}

func (style *Style) Cyan() *Style {
	style.fgColor = cyanFg
	return style
}

func (style *Style) White() *Style {
	style.fgColor = whiteFg
	return style
}

func (style *Style) TrueColor(red, green, blue byte) *Style {
	style.fgColor = isTrueColor
	style.fgTrueColor = [3]byte{red, green, blue}
	return style
}

func (style *Style) DefaultColor() *Style {
	style.fgColor = defaultColor
	return style
}

func (style *Style) BlackBg() *Style {
	style.bgColor = blackBg
	return style
}

func (style *Style) RedBg() *Style {
	style.bgColor = redBg
	return style
}

func (style *Style) GreenBg() *Style {
	style.bgColor = greenBg
	return style
}

func (style *Style) YellowBg() *Style {
	style.bgColor = yellowBg
	return style
}

func (style *Style) BlueBg() *Style {
	style.bgColor = blueBg
	return style
}

func (style *Style) MagentaBg() *Style {
	style.bgColor = magentaBg
	return style
}

func (style *Style) CyanBg() *Style {
	style.bgColor = cyanBg
	return style
}

func (style *Style) WhiteBg() *Style {
	style.bgColor = whiteBg
	return style
}

func (style *Style) TrueColorBg(red, green, blue byte) *Style {
	style.bgColor = isTrueColor
	style.bgTrueColor = [3]byte{red, green, blue}
	return style
}

func (style *Style) DefaultColorBg() *Style {
	style.bgColor = defaultColor
	return style
}

// Makes text have a regular weight, not bold or dim
func (style *Style) Regular() *Style {
	style.weight = regularWeight
	return style
}

// Makes text bold and not dim
func (style *Style) Bold() *Style {
	style.weight = boldWeight
	return style
}

// Makes text dim and not bold
func (style *Style) Dim() *Style {
	style.weight = dimWeight
	return style
}

// Makes text italicized
func (style *Style) Italic() *Style {
	style.italic = true
	return style
}

// Makes text not italicized
func (style *Style) NoItalic() *Style {
	style.italic = false
	return style
}

// Makes text underlined
func (style *Style) Underline() *Style {
	style.underline = true
	return style
}

// Makes text not underlined
func (style *Style) NoUnderline() *Style {
	style.underline = false
	return style
}

// Makes text crossed out
func (style *Style) Strikethrough() *Style {
	style.strikethrough = true
	return style
}

// Makes text not crossed out
func (style *Style) NoStrikethrough() *Style {
	style.strikethrough = false
	return style
}

func (style *Style) Print(args ...any) *Style {
	style.printRaw(fmt.Sprint(args...))
	return style
}

func (style *Style) Println(args ...any) *Style {
	style.printRaw(fmt.Sprint(args...))
	print("\n")
	return style
}

func (style *Style) Printf(format string, args ...any) *Style {
	style.printRaw(fmt.Sprintf(format, args...))
	return style
}

func (style *Style) printRaw(text string) {
	var codes []string

	if style.fgColor >= blackFg {
		codes = append(codes, strconv.Itoa(int(style.fgColor)))
	} else if style.fgColor == isTrueColor {
		color := style.fgTrueColor
		codes = append(codes, "38")
		codes = append(codes, "2")
		codes = append(codes, strconv.Itoa(int(color[0])))
		codes = append(codes, strconv.Itoa(int(color[1])))
		codes = append(codes, strconv.Itoa(int(color[2])))
	}

	if style.bgColor >= blackFg {
		codes = append(codes, strconv.Itoa(int(style.bgColor)))
	} else if style.bgColor == isTrueColor {
		color := style.bgTrueColor
		codes = append(codes, "48")
		codes = append(codes, "2")
		codes = append(codes, strconv.Itoa(int(color[0])))
		codes = append(codes, strconv.Itoa(int(color[1])))
		codes = append(codes, strconv.Itoa(int(color[2])))
	}

	if style.weight != regularWeight {
		if style.weight == boldWeight {
			codes = append(codes, "1")
		} else { // style.weight == dimWeight
			codes = append(codes, "2")
		}
	}

	if style.italic {
		codes = append(codes, "3")
	}

	if style.underline {
		codes = append(codes, "4")
	}

	if style.strikethrough {
		codes = append(codes, "9")
	}

	if len(codes) == 0 {
		print(text)
	} else {
		fmt.Printf("\x1b[%sm%s\x1b[0m", strings.Join(codes, ";"), text)
	}
}
