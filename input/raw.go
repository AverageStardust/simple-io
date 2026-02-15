package input

import (
	"encoding/binary"
	"iter"
	"os"

	"github.com/AverageStardust/simple-input/output"
	"golang.org/x/term"
)

var cookedTerminalState *term.State = nil

// Enters raw terminal mode and then gets key press as an iterator.
// Whenever pulling from the iterator the program will pause until the next key press.
// The returned iterator must be stopped to exit raw mode.
// Will immediately terminate the program with code 0 if the user enters ctrl+c.
func RawTerminalKeys() (keys iter.Seq[uint32]) {
	err := RawMode()
	if err != nil {
		return nil
	}

	return func(yield func(bytes uint32) bool) {
		defer CookedMode()

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
				// return to raw terminal mode since we are about to terminate
				CookedMode()

				output.ShowCursor()
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

// Enter raw terminal mode, meaning each key press are received one by one.
// Does nothing if it's already in raw mode.
func RawMode() (err error) {
	if cookedTerminalState == nil {
		cookedTerminalState, err = term.MakeRaw(int(os.Stdin.Fd()))
		return err
	} else {
		return nil
	}
}

// Enter cooked terminal mode, meaning key presses are only received after pressing enter.
// Does nothing if it's already in cooked mode.
func CookedMode() {
	if cookedTerminalState != nil {
		term.Restore(int(os.Stdin.Fd()), cookedTerminalState)
	}
}
