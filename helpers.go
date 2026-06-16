package main

import (
	"os"

	"github.com/charmbracelet/x/term"
)

func getTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(uintptr(os.Stdout.Fd()))
	return width, height, err
}
