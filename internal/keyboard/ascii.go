// Package keyboard renders the ASCII keyboard visualization and manages
// layout remapping (qwerty, dvorak, colemak, etc.) and custom layout files.
package keyboard

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	ansi "github.com/charmbracelet/x/ansi"
	evdev "github.com/gvalkov/golang-evdev"
)

func Render(layout string, size int, standard string, pressedKeys map[uint16]bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	sd, ok := resolveStandard(standard)
	if !ok {
		return ""
	}
	rows, ok := sd.sizes[size]
	if !ok {
		return ""
	}

	layoutMap := layouts[layout]
	remapped := make([][]key, len(rows))
	for i, row := range rows {
		remapped[i] = applyLayout(row, layoutMap)
	}

	shiftHeld := pressedKeys[evdev.KEY_LEFTSHIFT] || pressedKeys[evdev.KEY_RIGHTSHIFT]
	altGrHeld := pressedKeys[evdev.KEY_RIGHTALT]
	kanaHeld := pressedKeys[evdev.KEY_KATAKANAHIRAGANA]

	shiftMap := sd.shiftMap
	if lm, ok := shiftMaps[layout]; ok {
		shiftMap = lm
	}
	altGrMap := sd.altGrMap
	if am, ok := altGrMaps[layout]; ok {
		altGrMap = am
	}

	pressed := resolvePressed(rows, remapped, pressedKeys)
	if kanaHeld {
		applyKana(remapped, shiftHeld)
	}
	applyModifiers(remapped, shiftHeld, shiftMap, altGrHeld, altGrMap)
	return renderKeys(remapped, pressed, fingerStyle, fingerActive)
}

func applyLayout(keys []key, layoutMap map[string]string) []key {
	result := make([]key, len(keys))
	for i, k := range keys {
		if layoutMap != nil {
			if newLabel, ok := layoutMap[k.label]; ok {
				k.label = newLabel
			}
		}
		result[i] = k
	}
	return result
}

func resolveStandard(standard string) (standardData, bool) {
	sd, ok := standards[standard]
	return sd, ok
}

func resolvePressed(rows, remapped [][]key, pressedKeys map[uint16]bool) [][]bool {
	evCodeToOrigLabel := make(map[uint16]string)
	for _, row := range rows {
		for _, k := range row {
			evCodeToOrigLabel[k.evCode] = k.label
		}
	}

	labelCount := make(map[string]int)
	for _, row := range remapped {
		for _, k := range row {
			labelCount[k.label]++
		}
	}

	evCodeCount := make(map[uint16]int)
	for _, row := range rows {
		for _, k := range row {
			evCodeCount[k.evCode]++
		}
	}

	pressedByEvCode := make(map[uint16]bool)
	pressedByLabel := make(map[string]bool)
	for code, down := range pressedKeys {
		if !down {
			continue
		}
		if evCodeCount[code] > 1 {
			pressedByEvCode[code] = true
			continue
		}
		label, ok := evCodeToOrigLabel[code]
		if !ok {
			pressedByEvCode[code] = true
			continue
		}
		switch count := labelCount[label]; {
		case count > 1:
			pressedByEvCode[code] = true
		case count == 1:
			pressedByLabel[label] = true
		default:
			pressedByEvCode[code] = true
		}
	}

	pressed := make([][]bool, len(remapped))
	for i, keys := range remapped {
		pressed[i] = make([]bool, len(keys))
		for j, k := range keys {
			if pressedByEvCode[k.evCode] || pressedByLabel[k.label] {
				pressed[i][j] = true
			}
		}
	}
	return pressed
}

func applyModifiers(keys [][]key, shiftHeld bool, shiftMap map[string]string, altGrHeld bool, altGrMap map[string]string) {
	for _, row := range keys {
		if altGrHeld && altGrMap != nil {
			for j := range row {
				if newLabel, ok := altGrMap[row[j].label]; ok {
					row[j].label = newLabel
				}
			}
			if shiftHeld {
				for j := range row {
					row[j].label = strings.ToUpper(row[j].label)
				}
			}
		} else if shiftHeld && shiftMap != nil {
			for j := range row {
				if newLabel, ok := shiftMap[row[j].label]; ok {
					row[j].label = newLabel
				}
			}
		}
	}
}

func applyKana(keys [][]key, shiftHeld bool) {
	for _, row := range keys {
		for j := range row {
			label := row[j].label
			if shiftHeld {
				if v, ok := kanaSmallMap[label]; ok {
					row[j].label = v
					continue
				}
			}
			if v, ok := kanaLayout[label]; ok {
				row[j].label = v
			}
		}
	}
}

func renderKeys(keys [][]key, pressed [][]bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	var lines []string
	for i, kr := range keys {
		if i == 0 {
			lines = append(lines, topLine(kr))
		}
		lines = append(lines, midLine(kr, pressed[i], fingerStyle, fingerActive))
		if i < len(keys)-1 {
			lines = append(lines, divLine(kr, fingerStyle))
		} else {
			lines = append(lines, botLine(kr))
		}
	}
	return strings.Join(lines, "\n")
}

func topLine(keys []key) string {
	var b strings.Builder
	b.WriteByte(',')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.width))
		b.WriteByte(',')
	}
	return b.String()
}

func midLine(keys []key, pressed []bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	var b strings.Builder
	for i, k := range keys {
		label := k.label
		if k.divLabel != "" {
			label = ""
		}

		isPressed := i < len(pressed) && pressed[i]

		if i == 0 {
			if isPressed {
				b.WriteString(fingerActive[k.finger].Render("|"))
			} else {
				b.WriteByte('|')
			}
		}

		if isPressed {
			b.WriteString(fingerActive[k.finger].Render(centerLabel(label, k.width)))
		} else {
			b.WriteString(fingerStyle[k.finger].Render(centerLabel(label, k.width)))
		}

		if k.rightless {
			b.WriteByte(' ')
		} else {
			nextPressed := i+1 < len(pressed) && pressed[i+1]
			if isPressed || nextPressed {
				f := k.finger
				if nextPressed && !isPressed {
					f = keys[i+1].finger
				}
				b.WriteString(fingerActive[f].Render("|"))
			} else {
				b.WriteByte('|')
			}
		}
	}
	return b.String()
}

func divLine(keys []key, fingerStyle map[Finger]lipgloss.Style) string {
	var b strings.Builder
	b.WriteByte('|')
	for _, k := range keys {
		if k.gap {
			if k.divLabel != "" {
				b.WriteString(fingerStyle[k.finger].Render(centerLabel(k.divLabel, k.width)))
			} else {
				b.WriteString(strings.Repeat(" ", k.width))
			}
			if k.rightless {
				b.WriteByte(',')
			} else {
				b.WriteByte('\'')
			}
			continue
		}
		b.WriteString(strings.Repeat("-", k.width))
		if k.leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}

func botLine(keys []key) string {
	var b strings.Builder
	b.WriteByte('\'')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.width))
		if k.leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}

func centerLabel(s string, width int) string {
	vw := ansi.StringWidth(s)
	if vw >= width {
		return s
	}
	total := width - vw
	left := total / 2
	right := total - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
