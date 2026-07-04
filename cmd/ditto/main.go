package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/input"
	"github.com/arvingarciabtw/ditto/internal/tui"
)

func main() {
	cfg := config.LoadWithFlags()

	p := tea.NewProgram(tui.InitModel(cfg))
	if err := input.StartInput(p); err != nil {
		input.PrintStartError(err)
		os.Exit(1)
	}
	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
