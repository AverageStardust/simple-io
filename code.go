package input

import (
	"fmt"
	"iter"
	"slices"
	"strings"
	"time"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

type rawWriter struct {
	lineNumber *int
}

const lineNumberFormat = "%3d  "

// Creates a small code entry with syntax highlighting and line numbers
// Currently works poorly if the screen must scroll to see all of it
// Best for small snippits of code
func Code(lexerName, styleName string, lineLimit, tabSize int) string {
	code := new([]rune)

	quit := make(chan struct{})

	go codeStyler(code, quit, lexerName, styleName)
	codeEditor(code, lineLimit, tabSize)

	quit <- struct{}{}

	return string(*code)
}

// Controls the input and output of the code editor
func codeEditor(code *[]rune, lineLimit, tabSize int) {
	next, stop := iter.Pull(getRawTerminalKeys())
	defer stop()

	lineNumber := 1
	fmt.Printf("\033[H\033[J"+lineNumberFormat, lineNumber)

editLoop:
	for {
		key, ok := next()
		if !ok {
			break
		}

		switch {
		case key >= 32 && key <= 126: // ascii range
			rune := rune(key)
			*code = append(*code, rune)
			print(string(rune))

		case key == KEY_TAB:
			*code = append(*code, slices.Repeat([]rune{' '}, tabSize)...)
			print(strings.Repeat(" ", tabSize))

		case key == KEY_BACKSPACE || key == KEY_DELETE:
			deleteCodeRune(code, &lineNumber, tabSize)

		case key == KEY_ENTER:
			if lineNumber < lineLimit {
				*code = append(*code, '\n')
				lineNumber++
				fmt.Printf("\033[1E"+lineNumberFormat, lineNumber)
			}

		case key == KEY_ESCAPE:
			break editLoop
		}
	}

	print("\033[1E")
}

// Controls what happens when the backspace/delete keys are pressed in the editor
func deleteCodeRune(code *[]rune, lineNumber *int, tabSize int) {
	if len(*code) == 0 {
		return
	}

	switch (*code)[len(*code)-1] {
	case '\n':
		*code = (*code)[:len(*code)-1]
		*lineNumber--
		cols := countCols(*code)
		fmt.Printf("\033[1F\033[%dC\033[J", cols+5)

	case ' ':
		cols := countCols(*code)
		if cols%tabSize == 0 && cols >= tabSize {
			lineAllSpaces := true
			for i := range cols {
				if (*code)[len(*code)-i-1] != ' ' {
					lineAllSpaces = false
				}
			}

			if lineAllSpaces {
				// remove rest of a tab
				*code = (*code)[:len(*code)-tabSize]
				fmt.Printf("\033[%dD\033[K", tabSize)
				break
			}
		}

		fallthrough
	default:
		*code = (*code)[:len(*code)-1]
		print("\033[1D\033[K")
	}
}

// Replaces the code with a syntax highlighted version on a regular basis
func codeStyler(code *[]rune, quit chan struct{}, lexerName, styleName string) {
	lexer := lexers.Get(lexerName)
	style := styles.Get(styleName)

	oldCode := ""
	for {
		select {
		case <-quit:
			return
		default:
			currentCode := *code

			if len(currentCode) > 0 {
				trimmedCode := strings.TrimRight(string(currentCode), " ")

				if oldCode != trimmedCode {
					trailingSpaces := len(currentCode) - len(trimmedCode)

					print("\033[H\033[J")

					iterator, _ := lexer.Tokenise(nil, trimmedCode)
					formatters.TTY16m.Format(newRawWriter(), style, iterator)

					if trailingSpaces > 0 {
						fmt.Printf("\033[%dC", trailingSpaces)
					}

					oldCode = trimmedCode
				}
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Count the number of columns in the last line of code
func countCols(code []rune) int {
	cols := 0

	for i := len(code) - 1; i >= 0; i-- {
		if code[i] == '\n' {
			break
		} else {
			cols++
		}
	}

	return cols
}

// Count the number of rows in the code
func countRows(code []rune) int {
	rows := 1

	for _, rune := range code {
		if rune == '\n' {
			rows++
		}
	}

	return rows
}

// Creates a writer for writing in raw terminal mode
func newRawWriter() rawWriter {
	return rawWriter{new(int)}
}

// Writes to the screen in raw terminal mode
// Handles newlines with ANSI codes and inserts line numbers
func (writer rawWriter) Write(bytes []byte) (n int, err error) {
	lastWrite := 0

	if *writer.lineNumber == 0 {
		*writer.lineNumber++
		fmt.Printf(lineNumberFormat, *writer.lineNumber)
	}

	for i, byte := range bytes {
		if byte == '\n' {
			print(string(bytes[lastWrite:i]))
			*writer.lineNumber++
			fmt.Printf("\033[1E"+lineNumberFormat, *writer.lineNumber)

			lastWrite = i + 1
		}
	}
	print(string(bytes[lastWrite:]))

	return len(bytes), nil
}
