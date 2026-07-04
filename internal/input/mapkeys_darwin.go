//go:build darwin

package input

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

func init() {
	keyMapper = mapMacHidToEvdev
}

// NOTE:
// Still has to be tested on an actual device running macOS.
// Unfortunately, I don't have one...

// macOS HID keycodes (Apple HIToolbox kVK_* constants) → evdev keycodes.
const (
	macA          = 0x00
	macS          = 0x01
	macD          = 0x02
	macF          = 0x03
	macH          = 0x04
	macG          = 0x05
	macZ          = 0x06
	macX          = 0x07
	macC          = 0x08
	macV          = 0x09
	macISOExtra   = 0x0A
	macB          = 0x0B
	macQ          = 0x0C
	macW          = 0x0D
	macE          = 0x0E
	macR          = 0x0F
	macY          = 0x10
	macT          = 0x11
	mac1          = 0x12
	mac2          = 0x13
	mac3          = 0x14
	mac4          = 0x15
	mac6          = 0x16
	mac5          = 0x17
	macEqual      = 0x18
	mac9          = 0x19
	mac7          = 0x1A
	macMinus      = 0x1B
	mac8          = 0x1C
	mac0          = 0x1D
	macRBracket   = 0x1E
	macO          = 0x1F
	macU          = 0x20
	macLBracket   = 0x21
	macI          = 0x22
	macP          = 0x23
	macReturn     = 0x24
	macL          = 0x25
	macJ          = 0x26
	macK          = 0x27
	macSemicolon  = 0x28
	macBackslash  = 0x29
	macComma      = 0x2A
	macSlash      = 0x2B
	macN          = 0x2C
	macM          = 0x2D
	macPeriod     = 0x2E
	macTab        = 0x30
	macSpace      = 0x31
	macGrave      = 0x32
	macBackspace  = 0x33
	macEsc        = 0x35
	macLCmd       = 0x37
	macLShift     = 0x38
	macCapsLock   = 0x39
	macLOption    = 0x3A
	macLCtrl      = 0x3B
	macRShift     = 0x3C
	macROption    = 0x3D
	macRCtrl      = 0x3E
	macFn         = 0x3F
	macF17        = 0x40
	macNPDecimal  = 0x41
	macNPMultiply = 0x43
	macNPPlus     = 0x45
	macNPClear    = 0x47
	macNPDivide   = 0x4B
	macNPEnter    = 0x4C
	macNPMinus    = 0x4E
	macF18        = 0x4F
	macF19        = 0x50
	macNPEquals   = 0x51
	macNP0        = 0x52
	macNP1        = 0x53
	macNP2        = 0x54
	macNP3        = 0x55
	macNP4        = 0x56
	macNP5        = 0x57
	macNP6        = 0x58
	macNP7        = 0x59
	macF20        = 0x5A
	macNP8        = 0x5B
	macNP9        = 0x5C
	macF5         = 0x60
	macF6         = 0x61
	macF7         = 0x62
	macF3         = 0x63
	macF8         = 0x64
	macF9         = 0x65
	macF11        = 0x67
	macF13        = 0x69
	macF16        = 0x6A
	macF14        = 0x6B
	macF10        = 0x6D
	macF12        = 0x6F
	macF15        = 0x71
	macHelp       = 0x72
	macHome       = 0x73
	macPageUp     = 0x74
	macDel        = 0x75
	macF4         = 0x76
	macEnd        = 0x77
	macF2         = 0x78
	macPageDown   = 0x79
	macF1         = 0x7A
	macLeft       = 0x7B
	macRight      = 0x7C
	macDown       = 0x7D
	macUp         = 0x7E
)

func mapMacHidToEvdev(code uint16) uint16 {
	switch code {
	case macA:
		return base.KEY_A
	case macS:
		return base.KEY_S
	case macD:
		return base.KEY_D
	case macF:
		return base.KEY_F
	case macH:
		return base.KEY_H
	case macG:
		return base.KEY_G
	case macZ:
		return base.KEY_Z
	case macX:
		return base.KEY_X
	case macC:
		return base.KEY_C
	case macV:
		return base.KEY_V
	case macB:
		return base.KEY_B
	case macQ:
		return base.KEY_Q
	case macW:
		return base.KEY_W
	case macE:
		return base.KEY_E
	case macR:
		return base.KEY_R
	case macY:
		return base.KEY_Y
	case macT:
		return base.KEY_T
	case mac1:
		return base.KEY_1
	case mac2:
		return base.KEY_2
	case mac3:
		return base.KEY_3
	case mac4:
		return base.KEY_4
	case mac6:
		return base.KEY_6
	case mac5:
		return base.KEY_5
	case macEqual:
		return base.KEY_EQUAL
	case mac9:
		return base.KEY_9
	case mac7:
		return base.KEY_7
	case macMinus:
		return base.KEY_MINUS
	case mac8:
		return base.KEY_8
	case mac0:
		return base.KEY_0
	case macRBracket:
		return base.KEY_RIGHTBRACE
	case macO:
		return base.KEY_O
	case macU:
		return base.KEY_U
	case macLBracket:
		return base.KEY_LEFTBRACE
	case macI:
		return base.KEY_I
	case macP:
		return base.KEY_P
	case macReturn:
		return base.KEY_ENTER
	case macL:
		return base.KEY_L
	case macJ:
		return base.KEY_J
	case macK:
		return base.KEY_K
	case macSemicolon:
		return base.KEY_SEMICOLON
	case macBackslash:
		return base.KEY_BACKSLASH
	case macComma:
		return base.KEY_COMMA
	case macSlash:
		return base.KEY_SLASH
	case macN:
		return base.KEY_N
	case macM:
		return base.KEY_M
	case macPeriod:
		return base.KEY_DOT
	case macTab:
		return base.KEY_TAB
	case macSpace:
		return base.KEY_SPACE
	case macGrave:
		return base.KEY_GRAVE
	case macBackspace:
		return base.KEY_BACKSPACE
	case macEsc:
		return base.KEY_ESC
	case macLCmd:
		return base.KEY_LEFTMETA
	case macLShift:
		return base.KEY_LEFTSHIFT
	case macCapsLock:
		return base.KEY_CAPSLOCK
	case macLOption:
		return base.KEY_LEFTALT
	case macLCtrl:
		return base.KEY_LEFTCTRL
	case macRShift:
		return base.KEY_RIGHTSHIFT
	case macROption:
		return base.KEY_RIGHTALT
	case macRCtrl:
		return base.KEY_RIGHTCTRL
	case macFn:
		return base.KEY_FN
	case macISOExtra:
		return base.KEY_102ND
	case macNPDecimal:
		return base.KEY_KPDOT
	case macNPMultiply:
		return base.KEY_KPASTERISK
	case macNPPlus:
		return base.KEY_KPPLUS
	case macNPMinus:
		return base.KEY_KPMINUS
	case macNPDivide:
		return base.KEY_KPSLASH
	case macNPEnter:
		return base.KEY_KPENTER
	case macNP0:
		return base.KEY_KP0
	case macNP1:
		return base.KEY_KP1
	case macNP2:
		return base.KEY_KP2
	case macNP3:
		return base.KEY_KP3
	case macNP4:
		return base.KEY_KP4
	case macNP5:
		return base.KEY_KP5
	case macNP6:
		return base.KEY_KP6
	case macNP7:
		return base.KEY_KP7
	case macNP8:
		return base.KEY_KP8
	case macNP9:
		return base.KEY_KP9
	case macF1:
		return base.KEY_F1
	case macF2:
		return base.KEY_F2
	case macF3:
		return base.KEY_F3
	case macF4:
		return base.KEY_F4
	case macF5:
		return base.KEY_F5
	case macF6:
		return base.KEY_F6
	case macF7:
		return base.KEY_F7
	case macF8:
		return base.KEY_F8
	case macF9:
		return base.KEY_F9
	case macF10:
		return base.KEY_F10
	case macF11:
		return base.KEY_F11
	case macF12:
		return base.KEY_F12
	case macUp:
		return base.KEY_UP
	case macDown:
		return base.KEY_DOWN
	case macLeft:
		return base.KEY_LEFT
	case macRight:
		return base.KEY_RIGHT
	case macPageUp:
		return base.KEY_PAGEUP
	case macPageDown:
		return base.KEY_PAGEDOWN
	case macHome:
		return base.KEY_HOME
	case macEnd:
		return base.KEY_END
	case macDel:
		return base.KEY_DELETE
	default:
		return code
	}
}
