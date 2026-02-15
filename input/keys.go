package input

import (
	"fmt"
	"strings"
)

type Key uint64

const (
	KEY_UNKNOWN     Key = 0
	KEY_END_OF_TEXT Key = 3
	KEY_TAB         Key = 9
	KEY_ENTER       Key = 13
	KEY_ESCAPE      Key = 27
)

const (
	KEY_SPACE Key = iota + 32
	KEY_EXCLAMATION
	KEY_QUOTATION_MARK
	KEY_HASH
	KEY_DOLLAR
	KEY_PERCENT
	KEY_AMPERSAND
	KEY_APOSTROPHE
	KEY_LEFT_PARENTHESIS
	KEY_RIGHT_PARENTHESIS
	KEY_ASTERISK
	KEY_PLUS_SIGN
	KEY_COMMA
	KEY_HYPHEN
	KEY_PERIOD
	KEY_SLASH

	KEY_ZERO
	KEY_ONE
	KEY_TWO
	KEY_THREE
	KEY_FOUR
	KEY_FIVE
	KEY_SIX
	KEY_SEVEN
	KEY_EIGHT
	KEY_NINE

	KEY_COLON
	KEY_SEMICOLON
	KEY_LESS_THAN
	KEY_EQUAL
	KEY_GREATER_THAN
	KEY_QUESTION_MARK
	KEY_AT_SIGN

	KEY_SHIFT_A
	KEY_SHIFT_B
	KEY_SHIFT_C
	KEY_SHIFT_D
	KEY_SHIFT_E
	KEY_SHIFT_F
	KEY_SHIFT_G
	KEY_SHIFT_H
	KEY_SHIFT_I
	KEY_SHIFT_J
	KEY_SHIFT_K
	KEY_SHIFT_L
	KEY_SHIFT_M
	KEY_SHIFT_N
	KEY_SHIFT_O
	KEY_SHIFT_P
	KEY_SHIFT_Q
	KEY_SHIFT_R
	KEY_SHIFT_S
	KEY_SHIFT_T
	KEY_SHIFT_U
	KEY_SHIFT_V
	KEY_SHIFT_W
	KEY_SHIFT_X
	KEY_SHIFT_Y
	KEY_SHIFT_Z

	KEY_LEFT_SQUARE_BRACKET
	KEY_BACK_SLASH
	KEY_RIGHT_SQUARE_BRACKET
	KEY_EXPONENT
	KEY_UNDERSCORE
	KEY_GRAVE

	KEY_A
	KEY_B
	KEY_C
	KEY_D
	KEY_E
	KEY_F
	KEY_G
	KEY_H
	KEY_I
	KEY_J
	KEY_K
	KEY_L
	KEY_M
	KEY_N
	KEY_O
	KEY_P
	KEY_Q
	KEY_R
	KEY_S
	KEY_T
	KEY_U
	KEY_V
	KEY_W
	KEY_X
	KEY_Y
	KEY_Z

	KEY_LEFT_CURLY_BRACE
	KEY_PIPE
	KEY_RIGHT_CURLY_BRASE
	KEY_TILDE
	KEY_BACKSPACE
)

const (
	KEY_ALT_A Key = iota*256 + 24859
	KEY_ALT_B
	KEY_ALT_C
	KEY_ALT_D
	KEY_ALT_E
	KEY_ALT_F
	KEY_ALT_G
	KEY_ALT_H
	KEY_ALT_I
	KEY_ALT_J
	KEY_ALT_K
	KEY_ALT_L
	KEY_ALT_M
	KEY_ALT_N
	KEY_ALT_O
	KEY_ALT_P
	KEY_ALT_Q
	KEY_ALT_R
	KEY_ALT_S
	KEY_ALT_T
	KEY_ALT_U
	KEY_ALT_V
	KEY_ALT_W
	KEY_ALT_X
	KEY_ALT_Y
	KEY_ALT_Z
)

const (
	KEY_UP        Key = 4283163
	KEY_DOWN      Key = 4348699
	KEY_RIGHT     Key = 4414235
	KEY_LEFT      Key = 4479771
	KEY_END       Key = 4610843
	KEY_HOME      Key = 4741915
	KEY_INSERT    Key = 2117229339
	KEY_DELETE    Key = 2117294875
	KEY_PAGE_UP   Key = 2117425947
	KEY_PAGE_DOWN Key = 2117491483
)

func (key Key) Name() string {
	if key >= KEY_EXCLAMATION && key <= KEY_TILDE { // ascii text

		if key >= KEY_SHIFT_A && key <= KEY_SHIFT_Z {
			letter := rune(key)
			return fmt.Sprintf("Shift + %c", letter)
		}

		return strings.ToUpper(string(rune(key)))
	}

	switch key {
	case KEY_UNKNOWN:
		return "Unknown"
	case KEY_END_OF_TEXT:
		return "End of Text"
	case KEY_TAB:
		return "Tab"
	case KEY_ENTER:
		return "Enter"
	case KEY_ESCAPE:
		return "Escape"
	case KEY_UP:
		return "Arrow Up"
	case KEY_DOWN:
		return "Arrow Down"
	case KEY_RIGHT:
		return "Arrow Left"
	case KEY_LEFT:
		return "Arrow Right"
	case KEY_END:
		return "End"
	case KEY_HOME:
		return "Home"
	case KEY_INSERT:
		return "Insert"
	case KEY_DELETE:
		return "Delete"
	case KEY_PAGE_UP:
		return "Page Up"
	case KEY_PAGE_DOWN:
		return "Page Down"
	}

	if key >= KEY_ALT_A && key <= KEY_ALT_Z && (key-KEY_ALT_A)%256 == 0 {
		withoutAlt := (key-KEY_ALT_A)/256 + KEY_A
		letter := strings.ToUpper(string(rune(withoutAlt)))
		return fmt.Sprintf("Alt + %s", letter)
	}

	return "Unnamed"
}

func (key Key) known() bool {
	if key >= KEY_SPACE && key <= KEY_BACKSPACE { // ascii text
		return true
	}

	switch key {
	case KEY_END_OF_TEXT, KEY_TAB, KEY_ENTER, KEY_ESCAPE: // ascii control keys
		return true
	case KEY_UP, KEY_DOWN, KEY_RIGHT, KEY_LEFT: // arrow keys
		return true
	case KEY_END, KEY_HOME, KEY_INSERT, KEY_DELETE, KEY_PAGE_UP, KEY_PAGE_DOWN: // control keys
		return true
	}

	// alt letters
	if key >= KEY_ALT_A && key <= KEY_ALT_Z && (key-KEY_ALT_A)%256 == 0 {
		return true
	}

	return false
}
