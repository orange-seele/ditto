//go:build !linux

package input

import (
	tea "charm.land/bubbletea/v2"
	hook "github.com/robotn/gohook"
)

func ListenHook(p *tea.Program) {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		switch ev.Kind {
		case hook.KeyDown, hook.KeyUp:
			p.Send(KeyMsg{
				Code: keyMapper(ev.Keycode),
				Down: ev.Kind == hook.KeyDown,
			})
		}
	}
}
