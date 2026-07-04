//go:build !linux

package input

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func StartInput(p *tea.Program) error {
	go ListenHook(p)
	return nil
}

func PrintStartError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}
