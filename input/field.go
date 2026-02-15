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

// Propmts the user to enter a string.
func String(question string) (answer string, err error) {
	fmt.Printf("%s: ", question)

	answer, err = reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if err != nil {
		return "", err
	}

	print("\n")
	return answer, nil
}

// Propmts the user to enter an integer.
func Integer(question string) (answer int, err error) {
	fmt.Printf("%s: ", question)

	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	_, err = fmt.Sscanf(text, "%d", &answer)
	if err != nil {
		return 0, err
	}

	print("\n")
	return
}

// Propmts the user to enter a boolean.
func Confirm(question string) (confirmed bool) {
	fmt.Printf("%s (y/n): ", question)

	// scan a string to make sure we consume all input
	var answer string
	fmt.Scanf("%s", &answer)

	confirmed = len(answer) > 0 && (answer[0] == 'Y' || answer[0] == 'y')

	output.MoveUpToBeginning(1)
	output.EraseLine()
	if confirmed {
		fmt.Printf("%s: Yes\n", question)
	} else {
		fmt.Printf("%s: No\n", question)
	}

	print("\n")
	return
}

// Propmts the user to enter a password.
// The password is not display password in terminal.
// The password is hashed before being returned.
func Password() (passHash string, err error) {
	print("Password: ")

	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", nil
	}

	println(strings.Repeat("*", len(passwordBytes)))

	print("\n")

	hashBytes := sha256.Sum256(passwordBytes)
	return base64.RawStdEncoding.EncodeToString(hashBytes[:]), nil
}
