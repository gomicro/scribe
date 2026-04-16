// Package color provides ANSI escape code helpers for terminal color output.
// Every helper function applies bold in addition to the named color — the
// escape sequence used is [1;XXm (bold + color) rather than [XXm (color only).
// The exported constants (FgBlack, FgRed, etc.) are plain color codes without
// bold and can be used directly if non-bold output is needed.
package color

import "fmt"

const (
	escape = "\x1b"
)

const (
	Reset = iota
)

const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	FgHiBlack = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func unset() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

func BlackFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgBlack, s, unset())
}

func RedFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgRed, s, unset())
}

func GreenFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgGreen, s, unset())
}

func YellowFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgYellow, s, unset())
}

func BlueFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgBlue, s, unset())
}

func MagentaFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgMagenta, s, unset())
}

func CyanFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgCyan, s, unset())
}

func WhiteFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgWhite, s, unset())
}

func HiBlackFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiBlack, s, unset())
}

func HiRedFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiRed, s, unset())
}

func HiGreenFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiGreen, s, unset())
}

func HiYellowFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiYellow, s, unset())
}

func HiBlueFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiBlue, s, unset())
}

func HiMagentaFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiMagenta, s, unset())
}

func HiCyanFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiCyan, s, unset())
}

func HiWhiteFg(s string) string {
	return fmt.Sprintf("%s[1;%dm%s%s", escape, FgHiWhite, s, unset())
}
