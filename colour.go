package tty

import (
	"fmt"
)

// ECode is an ANSI escape code
type ECode int

// Ansi escape code constants. See
// http://ascii-table.com/ansi-escape-sequences.php

// General text attributes
const (
	OFF  ECode = iota
	BOLD       // 1
	_
	_
	UNDERLINE // 4
	BLINK     // 5
	_
	REVERSE   // 7
	CONCEALED // 8
)

// Foreground text attributes
const (
	BLACK   ECode = iota + 30
	RED           // 31
	GREEN         // 32
	YELLOW        // 33
	BLUE          // 34
	MAGENTA       // 35
	CYAN          // 36
	WHITE         // 37
)

// Background text attributes
const (
	BG_GREY    ECode = iota + 40
	BG_RED           // 41
	BG_GREEN         // 42
	BG_YELLOW        // 43
	BG_BLUE          // 44
	BG_MAGENTA       // 45
	BG_CYAN          // 46
	BG_WHITE         // 47
)

// AnsiEscape accepts ANSI escape codes and strings to form escape sequences.
// For example, to create a string with a colorized prefix,
//
//      AnsiEscape(BOLD, GREEN, "[DEBUG] ", OFF, "Here is the debug output")
//
// and a nicely escaped string for terminal output will be returned.
func AnsiEscape(c ...interface{}) (out string) {
	for _, val := range c {
		switch t := val.(type) {
		case ECode:
			out += fmt.Sprintf("\x1b[%dm", val)
		case string:
			out += fmt.Sprintf("%s", val)
		default:
			fmt.Printf("unexpected type: %T\n", t)
		}
	}
	if c[len(c)-1] != OFF {
		out += "\x1b[0m"
	}
	return
}

// AnsiEscapeS is like AnsiEscape, but takes only one escape code, and wrap the given
// text using that code and OFF
func AnsiEscapeS(code ECode, msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", code, msg)
}
