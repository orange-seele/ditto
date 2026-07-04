package input

type KeyMsg struct {
	Code uint16
	Down bool
}

var keyMapper func(uint16) uint16
