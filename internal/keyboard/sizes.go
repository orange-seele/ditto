package keyboard

type Finger int

const (
	Pinky Finger = iota
	Ring
	Middle
	Index
	Thumb
	Any
)

type key struct {
	label     string
	width     int
	finger    Finger
	gap       bool
	rightless bool
	leftless  bool
	divLabel  string
	evCode    uint16
}

/*
* Here is a general reference for all supported standards and how they are laid out
* https://en.wikipedia.org/wiki/Keyboard_layout#/media/File:Physical_keyboard_layouts_comparison_ANSI_ISO_KS_ABNT_JIS.png
 */

/*
* Reference for ANSI standard
* https://en.wikipedia.org/wiki/British_and_American_keyboards#/media/File:KB_United_States-NoAltGr.svg
 */
var sizesANSI = map[int][][]key{
	60:  size60ANSI,
	65:  size65ANSI,
	75:  size75ANSI,
	80:  size80ANSI,
	96:  size96ANSI,
	100: size100ANSI,
}

/*
* Reference for ISO standard
* https://en.wikipedia.org/wiki/British_and_American_keyboards#/media/File:KB_United_Kingdom.svg
 */
var sizesISO = map[int][][]key{
	60:  size60ISO,
	65:  size65ISO,
	75:  size75ISO,
	80:  size80ISO,
	96:  size96ISO,
	100: size100ISO,
}

/*
* Reference for ABNT standard
* https://en.wikipedia.org/wiki/List_of_QWERTY_keyboard_language_variants#/media/File:KB_Portuguese_Brazil.svg
 */
var sizesABNT = map[int][][]key{
	60:  size60ABNT,
	65:  size65ABNT,
	75:  size75ABNT,
	80:  size80ABNT,
	96:  size96ABNT,
	100: size100ABNT,
}

/*
* Reference for JIS standard
* https://en.wikipedia.org/wiki/Japanese_input_method#/media/File:KB_Japanese.svg
 */
var sizesJIS = map[int][][]key{
	60:  size60JIS,
	65:  size65JIS,
	75:  size75JIS,
	80:  size80JIS,
	96:  size96JIS,
	100: size100JIS,
}

var size60ANSI = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
	},
	{
		{label: "Tab↹", width: u2, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE, keyR,
		keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
	},
	{
		{label: "Caps", width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: "Enter", width: u2_50, finger: keyEnter.finger, evCode: keyEnter.evCode},
	},
	{
		{label: keyLeftShift.label, width: u3_50, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_50, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
	},
}

var size65ANSI = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		{label: "Tab↹", width: u2, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE, keyR,
		keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: "Enter", width: keyEnter.width, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		{label: keyLeftShift.label, width: u2_75, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u6, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyFn, keyRightCtrl, keyLeft, keyDown, keyRight,
	},
}

var size75ANSI = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyPrintScreen.label, width: keyPrintScreen.width, finger: Any, evCode: keyPrintScreen.evCode},
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace,
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyBackslash,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u2_75, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyLeftShift, keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp,
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
	},
}

var size80ANSI = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u3_50, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
	},
	{
		{label: keyLeftShift.label, width: u3_75, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
	},
}

var size96ANSI = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace, keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyBackslash, keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u2_75, finger: keyEnter.finger, evCode: keyEnter.evCode},
		keyPad4,
		keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		keyLeftShift, keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp, keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
		{label: keyPad0.label, width: u1, finger: Middle, evCode: keyPad0.evCode},
		keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size100ANSI = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u3_50, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
		keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		{label: keyLeftShift.label, width: u3_75, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad0, keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size60ISO = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
	},
	{
		{label: keyLeftShiftISO.label, width: u2, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_50, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
	},
}

var size65ISO = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyLeftShiftISO,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u6, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyFn, keyRightCtrl, keyLeft, keyDown, keyRight,
	},
}

var size75ISO = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyPrintScreen.label, width: keyPrintScreen.width, finger: Any, evCode: keyPrintScreen.evCode},
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace,
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyEnterISO,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound, keyEnterISOBlank,
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyLeftShiftISO,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShiftISO, keyUp,
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
	},
}

var size80ISO = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe, keyPound,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
	},
	{
		keyLeftShift,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
	},
}

var size96ISO = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace, keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyEnterISO, keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound, keyEnterISOBlank, keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		keyLeftShiftISO,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShiftISO, keyUp,
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
		{label: keyPad0.label, width: u1, finger: Middle, evCode: keyPad0.evCode},
		keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size100ISO = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe, keyPound,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
		keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		keyLeftShift,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad0, keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size60ABNT = [][]key{
	{
		keyApostrophe, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyAcute, keyLeftBrace,
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keyCedilla, keyTilde, keyRightBrace,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
	},
	{
		{label: "Shf", width: u1_75, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySemicolon, keySlash,
		{label: keyRightShiftISO.label, width: u2, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_50, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
	},
}

var size65ABNT = [][]key{
	{
		keyApostrophe, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyAcute, keyLeftBrace,
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keyCedilla, keyTilde, keyRightBrace,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		{label: keyLeftShiftISO.label, width: u1_50, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySemicolon, keySlash,
		{label: "Shf", width: u1, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		keyUp,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u6, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyFn, keyRightCtrl, keyLeft, keyDown, keyRight,
	},
}

var size75ABNT = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyPrintScreen.label, width: keyPrintScreen.width, finger: Any, evCode: keyPrintScreen.evCode},
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyApostrophe, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace,
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyAcute, keyLeftBrace,
		keyEnterISO,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keyCedilla, keyTilde, keyRightBrace,
		keyEnterISOBlank,
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		{label: "Shf", width: u1, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySemicolon, keySlash,
		{label: "Shf", width: u1, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		keyUp,
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
	},
}

var size80ABNT = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
	},
	{
		keyApostrophe, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyAcute, keyLeftBrace, keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keyCedilla, keyTilde, keyRightBrace,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
	},
	{
		{label: keyLeftShiftISO.label, width: u2, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySemicolon, keySlash,
		{label: keyRightShiftISO.label, width: u2, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u0_75, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
	},
}

var size96ABNT = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyApostrophe, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace, keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyAcute, keyLeftBrace,
		keyEnterISO, keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keyCedilla, keyTilde, keyRightBrace,
		keyEnterISOBlank, keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		{label: "Shf", width: u1, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySemicolon, keySlash,
		{label: "Shf", width: u1, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		keyUp,
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
		{label: keyPad0.label, width: u1, finger: Middle, evCode: keyPad0.evCode},
		keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size100ABNT = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyApostrophe, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyAcute, keyLeftBrace, keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keyCedilla, keyTilde, keyRightBrace,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
		keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		{label: keyLeftShiftISO.label, width: u2, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySemicolon, keySlash,
		{label: keyRightShiftISO.label, width: u2, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u0_75, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad0, keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size60JIS = [][]key{
	{
		{label: "全半", width: u1_50, finger: Pinky, evCode: keyApostrophe.evCode},
		key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus,
		{label: "^", width: u1, finger: Pinky, evCode: keyBlank.evCode},
		keyYen,
		{label: "<", width: u1, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP,
		{label: "@", width: u1, finger: keyLeftBrace.finger, evCode: keyLeftBrace.evCode},
		{label: "[", width: u1, finger: keyRightBrace.finger, evCode: keyRightBrace.evCode},
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyColon,
		keyRightBrace,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
	},
	{
		{label: "Shift", width: u2_50, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		{label: "Shift", width: u2_75, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keyMuhenkan,
		{label: keySpace.label, width: u3_50, finger: keySpace.finger, evCode: keySpace.evCode},
		keyHenkan, keyKana, keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
	},
}

var size65JIS = [][]key{
	{
		{label: "全半", width: u1_50, finger: Pinky, evCode: keyApostrophe.evCode},
		key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus,
		{label: "^", width: u1, finger: Pinky, evCode: keyBlank.evCode},
		keyYen,
		{label: "<", width: u1, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP,
		{label: "@", width: u1, finger: keyLeftBrace.finger, evCode: keyLeftBrace.evCode},
		{label: "[", width: u1, finger: keyRightBrace.finger, evCode: keyRightBrace.evCode},
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyColon,
		keyRightBrace,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		{label: "Shf", width: u1_75, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		{label: "Shft", width: u2, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		keyUp,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta,
		{label: keyLeftAlt.label, width: u1, finger: keyLeftAlt.finger, evCode: keyLeftAlt.evCode},
		keyMuhenkan,
		{label: keySpace.label, width: u3_50, finger: keySpace.finger, evCode: keySpace.evCode},
		keyHenkan, keyKana,
		{label: keyRightAlt.label, width: u1, finger: keyRightAlt.finger, evCode: keyRightAlt.evCode},
		keyFn, keyRightCtrl, keyLeft, keyDown, keyRight,
	},
}

var size75JIS = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyPrintScreen.label, width: keyPrintScreen.width, finger: Any, evCode: keyPrintScreen.evCode},
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		{label: "全", width: u1, finger: Pinky, evCode: keyApostrophe.evCode},
		key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus,
		{label: "^", width: u1, finger: Pinky, evCode: keyBlank.evCode},
		keyYen,
		{label: "<", width: u1, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP,
		{label: "@", width: u1, finger: keyLeftBrace.finger, evCode: keyLeftBrace.evCode},
		{label: "[", width: u1, finger: keyRightBrace.finger, evCode: keyRightBrace.evCode},
		keyEnterISO,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyColon,
		keyRightBrace,
		keyEnterISOBlank,
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		{label: "Shf", width: u1_75, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		{label: "Shf", width: u1_75, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		keyUp,
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta,
		{label: keyLeftAlt.label, width: u1, finger: keyLeftAlt.finger, evCode: keyLeftAlt.evCode},
		keyMuhenkan,
		{label: keySpace.label, width: u3, finger: keySpace.finger, evCode: keySpace.evCode},
		keyHenkan, keyKana,
		{label: keyRightAlt.label, width: u1, finger: keyRightAlt.finger, evCode: keyRightAlt.evCode},
		keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
	},
}

var size80JIS = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
	},
	{
		{label: "全半", width: u1_50, finger: Pinky, evCode: keyApostrophe.evCode},
		key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus,
		{label: "^", width: u1, finger: Pinky, evCode: keyBlank.evCode},
		keyYen,
		{label: "<-", width: u1_50, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP,
		{label: "@", width: u1, finger: keyLeftBrace.finger, evCode: keyLeftBrace.evCode},
		{label: "[", width: u1, finger: keyRightBrace.finger, evCode: keyRightBrace.evCode},
		keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyColon,
		keyRightBrace,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
	},
	{
		{label: "Shift", width: u2_50, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		{label: "Shift", width: u3, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u0_75, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		keyMuhenkan,
		{label: keySpace.label, width: u3_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyHenkan, keyKana,
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
	},
}

var size96JIS = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		{label: "全", width: u1, finger: Pinky, evCode: keyApostrophe.evCode},
		key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus,
		{label: "^", width: u1, finger: Pinky, evCode: keyBlank.evCode},
		keyYen,
		{label: "<", width: u1, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP,
		{label: "@", width: u1, finger: keyLeftBrace.finger, evCode: keyLeftBrace.evCode},
		{label: "[", width: u1, finger: keyRightBrace.finger, evCode: keyRightBrace.evCode},
		keyEnterISO, keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyColon,
		keyRightBrace,
		keyEnterISOBlank, keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		{label: "Shf", width: u1_75, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		{label: "Shf", width: u1_75, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		keyUp,
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta,
		{label: keyLeftAlt.label, width: u1, finger: keyLeftAlt.finger, evCode: keyLeftAlt.evCode},
		keyMuhenkan,
		{label: keySpace.label, width: u3, finger: keySpace.finger, evCode: keySpace.evCode},
		keyHenkan, keyKana,
		{label: keyRightAlt.label, width: u1, finger: keyRightAlt.finger, evCode: keyRightAlt.evCode},
		keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
		{label: keyPad0.label, width: u1, finger: Middle, evCode: keyPad0.evCode},
		keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size100JIS = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		{label: "全半", width: u1_50, finger: Pinky, evCode: keyApostrophe.evCode},
		key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus,
		{label: "^", width: u1, finger: Pinky, evCode: keyBlank.evCode},
		keyYen,
		{label: "<-", width: u1_50, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP,
		{label: "@", width: u1, finger: keyLeftBrace.finger, evCode: keyLeftBrace.evCode},
		{label: "[", width: u1, finger: keyRightBrace.finger, evCode: keyRightBrace.evCode},
		keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyColon,
		keyRightBrace,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
		keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		{label: "Shift", width: u2_50, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		{label: "Shift", width: u3, finger: keyRightShiftISO.finger, evCode: keyRightShiftISO.evCode},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u0_75, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		keyMuhenkan,
		{label: keySpace.label, width: u3_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyHenkan, keyKana,
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad0, keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}
