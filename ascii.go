package main

import (
	"strings"
)

func buildKeyboard(size int) string {
	rows, ok := keyboardSizes[size]
	if !ok {
		return ""
	}
	var lines []string
	for i, row := range rows {
		if i == 0 {
			lines = append(lines, buildTopLine(row))
		}
		lines = append(lines, buildMidLine(row))
		if i < len(rows)-1 {
			lines = append(lines, buildDivLine(row))
		} else {
			lines = append(lines, buildBotLine(row))
		}
	}
	return strings.Join(lines, "\n")
}

func buildTopLine(keys []Key) string {
	var b strings.Builder
	b.WriteByte(',')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.Width))
		b.WriteByte(',')
	}
	return b.String()
}

func buildMidLine(keys []Key) string {
	var b strings.Builder
	b.WriteByte('|')
	for _, k := range keys {
		label := k.Label
		if k.DivLabel != "" {
			label = ""
		}
		b.WriteString(fingerStyle[k.Finger].Render(centerLabel(label, k.Width)))
		if k.Rightless {
			b.WriteByte(' ')
			continue
		}
		b.WriteByte('|')
	}
	return b.String()
}

func buildDivLine(keys []Key) string {
	var b strings.Builder
	b.WriteByte('|')
	for _, k := range keys {
		if k.Gap {
			if k.DivLabel != "" {
				b.WriteString(fingerStyle[k.Finger].Render(centerLabel(k.DivLabel, k.Width)))
			} else {
				b.WriteString(strings.Repeat(" ", k.Width))
			}
			if k.Rightless {
				b.WriteByte(',')
			} else {
				b.WriteByte('\'')
			}
			continue
		}
		b.WriteString(strings.Repeat("-", k.Width))
		if k.Leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}

func buildBotLine(keys []Key) string {
	var b strings.Builder
	b.WriteByte('\'')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.Width))
		if k.Leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}
