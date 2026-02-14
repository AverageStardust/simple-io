package main

import (
	"encoding/binary"
	"iter"
	"os"

	"golang.org/x/term"
)

// Enters raw terminal mode and then gets key press as an iterator.
// Whenever pulling from the iterator the program will pause until the next key press.
// The returned iterator must be stopped to exit raw mode.
// Will immediately terminate the program with code 0 if the user enters ctrl+c.
func RawTerminalKeys() (keys iter.Seq[uint32]) {
	// enable raw terminal mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil
	}

	return func(yield func(bytes uint32) bool) {
		// later disable raw terminal mode
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		for {
			// get raw terminal key
			var buffer [8]byte
			n, err := os.Stdin.Read(buffer[:])
			if err != nil {
				return
			}

			// can't parse fully, ignore
			if n > 4 {
				continue
			}

			key := binary.LittleEndian.Uint32(buffer[:])

			if key == KEY_END_OF_TEXT {
				// disable raw terminal mode since we are about to immediately terminate
				term.Restore(int(os.Stdin.Fd()), oldState)

				ShowCursor()
				print("Force Quit\n")
				os.Exit(0)
			}

			if !knownKey(key) {
				key = KEY_UNKNOWN
			}

			if !yield(key) {
				return
			}
		}
	}
}
