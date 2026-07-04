//go:build linux

package input

import (
	tea "charm.land/bubbletea/v2"
)

func StartInput(p *tea.Program) error {
	devs, err := Devices()
	if err != nil {
		return err
	}
	for _, dev := range devs {
		go ListenToKeyboard(p, dev)
	}
	return nil
}

func PrintStartError(err error) {
	PrintDeviceError(err)
}
