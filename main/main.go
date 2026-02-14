package main

import input "github.com/AverageStardust/simple-input"

func main() {
	for key := range input.RawTerminalKeys() {
		println(key)
	}
}
