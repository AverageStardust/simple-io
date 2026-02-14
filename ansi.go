package main

import "fmt"

func MoveToHome() {
	print("\x1B[H")
}

func MoveTo(row int, col int) {
	fmt.Printf("\x1B[%d;%dH", col, row)
}

func MoveUp(rows int) {
	fmt.Printf("\x1B[%dA", rows)
}

func MoveDown(rows int) {
	fmt.Printf("\x1B[%dB", rows)
}

func MoveRight(cols int) {
	fmt.Printf("\x1B[%dC", cols)
}

func MoveLeft(cols int) {
	fmt.Printf("\x1B[%dD", cols)
}

func MoveDownToBeginning(rows int) {
	fmt.Printf("\x1B[%dE", rows)
}

func MoveUpToBeginning(rows int) {
	fmt.Printf("\x1B[%dF", rows)
}

func MoveToColumn(col int) {
	fmt.Printf("\x1B[%dG", col)
}

func SavePosition() {
	print("\x1B7")
}

func RestorePosition() {
	print("\x1B8")
}

func EraseToScreenEnd() {
	print("\x1B[0J")
}

func EraseToScreenStart() {
	print("\x1B[1J")
}

func EraseScreen() {
	print("\x1B[2J")
}

func EraseToLineEnd() {
	print("\x1B[0K")
}

func EraseToLineStart() {
	print("\x1B[1K")
}

func EraseLine() {
	print("\x1B[2K")
}

func HideCursor() {
	print("\x1B[?25l")
}

func ShowCursor() {
	print("\x1B[?25h")
}
