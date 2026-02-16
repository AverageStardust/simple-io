package input

import (
	"encoding/binary"
	"iter"
	"os"

	"github.com/AverageStardust/simple-io/output"
	"golang.org/x/term"
)

var cookedTerminalState *term.State = nil

// Enters raw terminal mode and then gets key press as an iterator.
// Whenever pulling from the iterator the program will pause until the next key press.
// The returned iterator must be stopped to exit raw mode.
// Will immediately terminate the program with code 0 if the user enters ctrl+c.
func RawKeys() (keys iter.Seq[Key]) {
	return func(yield func(bytes Key) bool) {
		err := RawMode()
		defer CookedMode()

		if err != nil {
			return
		}

		for {
			key, err := rawKeyAssuingMode()

			if err != nil {
				return
			}

			if key != KEY_UNKNOWN {
				if !yield(key) {
					return
				}
			}
		}
	}
}

func RawKey() Key {
	RawMode()
	defer CookedMode()

	key, err := rawKeyAssuingMode()

	if err != nil {
		return KEY_UNKNOWN
	} else {
		return key
	}
}

// Get a raw key input, assuming we are already in raw mode
func rawKeyAssuingMode() (Key, error) {
	// get raw terminal key
	var buffer [16]byte
	n, err := os.Stdin.Read(buffer[:])
	if err != nil {
		return 0, err
	}

	var key = KEY_UNKNOWN
	if n <= 8 {
		key = Key(binary.LittleEndian.Uint64(buffer[:8]))

		if !key.known() {
			key = KEY_UNKNOWN
		}
	}

	if key == KEY_END_OF_TEXT {
		// return to raw terminal mode since we are about to terminate
		CookedMode()

		output.ShowCursor()
		print("Force Quit\n")
		os.Exit(0)
	}

	return key, nil
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
		cookedTerminalState = nil
	}
}
