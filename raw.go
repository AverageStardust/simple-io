package input

import (
	"encoding/binary"
	"iter"
	"os"

	"golang.org/x/term"
)

const (
	KEY_END_OF_TEXT = 3
	KEY_TAB         = 9
	KEY_ENTER       = 13
	KEY_ESCAPE      = 27
	KEY_BACKSPACE   = 127
	KEY_RIGHT       = 4414235
	KEY_LEFT        = 4479771
	KEY_DELETE      = 2117294875
)

// Enters raw terminal mode and gets raw key inputs as an iterator
// Iterator must be finished once it is started, or else the terminal will remain in raw mode
// Will immediately terminate the program with code 0 if the user enters ctrl+c
func getRawTerminalKeys() (keys iter.Seq[uint32]) {
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
			var buffer [4]byte
			_, err := os.Stdin.Read(buffer[:])
			if err != nil {
				return
			}
			key := binary.LittleEndian.Uint32(buffer[:])

			if !yield(key) {
				return
			}

			if key == KEY_END_OF_TEXT {
				// disable raw terminal mode since we are about to immediately terminate
				term.Restore(int(os.Stdin.Fd()), oldState)
				print("\nForce Quit\n")
				os.Exit(0)
			}
		}
	}
}
