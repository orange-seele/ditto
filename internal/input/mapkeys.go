//go:build !linux && !darwin

package input

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

func init() {
	keyMapper = mapGohookToEvdev
}

// gohook Windows virtual keycodes used by robotn/gohook.
// Most letter/number keys pass through 1:1 with evdev codes.
// These are the ones that need explicit remapping.
const (
	hookF11         = 69
	hookF12         = 70
	hookNumpadSlash = 3637
	hookNumpadEnter = 3612
	hookRAlt        = 3640
	hookRCtrl       = 3641
	hookPause       = 3673
	hookScrollLock  = 3674
	hookCmd         = 3675
	hookRCmd        = 3676
	hookMenu        = 3677
	hookHome        = 3679
	hookUp          = 57416
	hookDown        = 57424
	hookLeft        = 57419
	hookRight       = 57421
	hookPageUp      = 57425
	hookPageDown    = 57426
	hookInsert      = 57427
	hookDelete      = 57423
	hookHome2       = 57435
	hookEnd         = 57436
	hookF11_2       = 57431
	hookF12_2       = 57433
	hookKPDot       = 57434
	hookKPComma     = 57372
	hookPause2      = 57428
	hookScrollLock2 = 57429
	hookCapsLock    = 57399
	hookKPDot2      = 57401
)

// Internal evdev-like keycodes not yet defined in the base package.
const (
	evPause = 119
)

func mapGohookToEvdev(code uint16) uint16 {
	switch code {
	case hookF11:
		return base.KEY_F11
	case hookF12:
		return base.KEY_F12
	case hookNumpadSlash:
		return base.KEY_KPSLASH
	case hookNumpadEnter:
		return base.KEY_KPENTER
	case hookRAlt:
		return base.KEY_RIGHTALT
	case hookRCtrl:
		return base.KEY_RIGHTCTRL
	case hookPause:
		return evPause
	case hookScrollLock:
		return base.KEY_SCROLLLOCK
	case hookCmd:
		return base.KEY_LEFTMETA
	case hookRCmd:
		return base.KEY_RIGHTMETA
	case hookMenu:
		return base.KEY_HANJA
	case hookHome:
		return base.KEY_HOME
	case hookUp:
		return base.KEY_UP
	case hookDown:
		return base.KEY_DOWN
	case hookLeft:
		return base.KEY_LEFT
	case hookRight:
		return base.KEY_RIGHT
	case hookPageUp:
		return base.KEY_PAGEUP
	case hookPageDown:
		return base.KEY_PAGEDOWN
	case hookInsert:
		return base.KEY_INSERT
	case hookDelete:
		return base.KEY_DELETE
	case hookHome2:
		return base.KEY_HOME
	case hookEnd:
		return base.KEY_END
	case hookF11_2:
		return base.KEY_F11
	case hookF12_2:
		return base.KEY_F12
	case hookKPDot:
		return base.KEY_KPDOT
	case hookKPComma:
		return base.KEY_YEN
	case hookPause2:
		return evPause
	case hookScrollLock2:
		return base.KEY_SCROLLLOCK
	case hookCapsLock:
		return base.KEY_CAPSLOCK
	case hookKPDot2:
		return base.KEY_SYSRQ
	default:
		return code
	}
}
