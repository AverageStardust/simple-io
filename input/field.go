package input

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/AverageStardust/simple-io/output"
	"golang.org/x/term"
)

var reader = bufio.NewReader(os.Stdin)

// Prompts the user to enter a string.
func String(question string) (answer string, err error) {
	output.LockScreen()
	fmt.Printf("%s: ", question)
	output.UnlockScreen()

	answer, err = reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if err != nil {
		return "", err
	}

	output.LockScreen()
	print("\n")
	output.UnlockScreen()

	callOnEntered()
	return answer, nil
}

// Prompts the user to enter an integer.
func Integer(question string) (answer int, err error) {
	output.LockScreen()
	fmt.Printf("%s: ", question)
	output.UnlockScreen()

	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	_, err = fmt.Sscanf(text, "%d", &answer)
	if err != nil {
		return 0, err
	}

	output.LockScreen()
	print("\n")
	output.UnlockScreen()

	callOnEntered()
	return
}

// Prompts the user to enter a boolean.
func Confirm(question string) (confirmed bool) {
	output.LockScreen()
	fmt.Printf("%s (y/n): ", question)
	output.UnlockScreen()

	// scan a string to make sure we consume all input
	var answer string
	fmt.Scanf("%s", &answer)

	confirmed = len(answer) > 0 && (answer[0] == 'Y' || answer[0] == 'y')

	output.LockScreen()
	output.MoveUpToBeginning(1)
	output.EraseLine()
	if confirmed {
		fmt.Printf("%s: Yes\n", question)
	} else {
		fmt.Printf("%s: No\n", question)
	}
	print("\n")
	output.UnlockScreen()

	callOnEntered()
	return
}

// Prompts the user to enter a password.
// The password is not display password in terminal.
// The password is hashed before being returned.
func Password() (passHash string, err error) {
	output.LockScreen()
	print("Password: ")
	output.UnlockScreen()

	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", nil
	}

	output.LockScreen()
	println(strings.Repeat("*", len(passwordBytes)))
	print("\n")
	output.UnlockScreen()

	hashBytes := sha256.Sum256(passwordBytes)
	return base64.RawStdEncoding.EncodeToString(hashBytes[:]), nil
}

func KeyBind() (key Key) {
	output.LockScreen()
	print("Key Bind: ")
	output.UnlockScreen()

	key = KEY_UNKNOWN
	for key == KEY_UNKNOWN {
		key = RawKey()
	}

	output.LockScreen()
	println(key.Name())
	output.UnlockScreen()

	callOnEntered()
	return key
}
