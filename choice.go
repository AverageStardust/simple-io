package input

import (
	"fmt"
	"iter"
	"slices"
)

type Choice struct {
	options []string
	width   int
	result  int
}

const NO_RESULT = -1

// Creates a new choice with options.
func NewChoice(options ...string) *Choice {
	width := 0
	for _, option := range options {
		width = max(width, len(option))
	}

	return &Choice{
		options: options,
		width:   width,
		result:  NO_RESULT,
	}
}

// Appends a single new option to a choice.
func (choice *Choice) Add(option string) *Choice {
	choice.options = append(choice.options, option)
	choice.width = max(choice.width, len(option))

	return choice
}

// Appends a slice of new options to a choice.
func (choice *Choice) AddSlice(options []string) *Choice {
	choice.options = slices.Concat(choice.options, options)

	for _, option := range options {
		choice.width = max(choice.width, len(option))
	}

	return choice
}

// Get the string of the selected option, asking the user if needed.
// Returns an empty string if failed to select anything.
func (choice *Choice) ResultString() string {
	index := choice.ResultIndex()

	if index == NO_RESULT {
		return ""
	}

	return choice.options[index]
}

// Gets the index of the selected option, asking the user if needed.
// Returns NO_RESULT if failed to select anything.
func (choice *Choice) ResultIndex() int {
	if choice.result == NO_RESULT {
		choice.Ask()
	}

	return choice.result
}

// Forgets any option already selected.
func (choice *Choice) Forget() *Choice {
	choice.result = NO_RESULT
	return choice
}

// Ask the user, replacing any older selection.
// May fail to select anything, leaving the result as NO_RESULT.
func (choice *Choice) Ask() *Choice {
	choice.Forget()

	if len(choice.options) == 1 {
		choice.result = 0
	}

	if len(choice.options) <= 1 {
		return choice
	}

	print("\033[?25l")       // hide cursor
	defer print("\033[?25h") // show cursor afterward

	index := 0
	if choice.result != NO_RESULT {
		index = choice.result
	}

	choice.render(index, false, false)

	next, stop := iter.Pull(RawTerminalKeys())
	defer stop()
renderLoop:
	for {
		key, ok := next()
		if !ok {
			break
		}

		switch key {
		case KEY_LEFT:
			index -= 1
			if index < 0 {
				index = len(choice.options) - 1
			}

		case KEY_RIGHT:
			index += 1
			if index >= len(choice.options) {
				index = 0
			}

		case KEY_ENTER:
			choice.result = index
			break renderLoop

		case KEY_ESCAPE:
			break renderLoop
		}

		choice.render(index, true, false)
	}

	choice.render(index, true, true)
	print("\n")

	return choice
}

// Displays a choice in the terminal.
func (choice *Choice) render(index int, redraw bool, done bool) {
	if redraw {
		print("\033[2K\033[0G")
	}

	if !done {
		fmt.Printf("Select: <- %-*s ->", choice.width, choice.options[index])
	} else {
		fmt.Printf("Select: %s\n", choice.options[index])
		print("\033[0G")
	}
}
