package keycodes

import "slices"

var keycodeGroups = map[KeycodeGroup]Keycodes{
	KeycodeGroupAll: slices.Concat(
		characters,
		numbers,
		numpad,
		modifiers,
		arrows,
		special,
		fkeys,
	),
	KeycodeGroupCharacters:   characters,
	KeycodeGroupNumbers:      numbers,
	KeycodeGroupNumpad:       numpad,
	KeycodeGroupModifiers:    modifiers,
	KeycodeGroupArrows:       arrows,
	KeycodeGroupSpecial:      special,
	KeycodeGroupFunctionKeys: fkeys,
}

var characters = Keycodes{
	{65, "A", "KeyA"},
	{66, "B", "KeyB"},
	{67, "C", "KeyC"},
	{68, "D", "KeyD"},
	{69, "E", "KeyE"},
	{70, "F", "KeyF"},
	{71, "G", "KeyG"},
	{72, "H", "KeyH"},
	{73, "I", "KeyI"},
	{74, "J", "KeyJ"},
	{75, "K", "KeyK"},
	{76, "L", "KeyL"},
	{77, "M", "KeyM"},
	{78, "N", "KeyN"},
	{79, "O", "KeyO"},
	{80, "P", "KeyP"},
	{81, "Q", "KeyQ"},
	{82, "R", "KeyR"},
	{83, "S", "KeyS"},
	{84, "T", "KeyT"},
	{85, "U", "KeyU"},
	{86, "V", "KeyV"},
	{87, "W", "KeyW"},
	{88, "X", "KeyX"},
	{89, "Y", "KeyY"},
	{90, "Z", "KeyZ"},
}

var numbers = Keycodes{
	{48, "0", "Digit0"},
	{49, "1", "Digit1"},
	{50, "2", "Digit2"},
	{51, "3", "Digit3"},
	{52, "4", "Digit4"},
	{53, "5", "Digit5"},
	{54, "6", "Digit6"},
	{55, "7", "Digit7"},
	{56, "8", "Digit8"},
	{57, "9", "Digit9"},
}

var numpad = Keycodes{
	{96, "num0", "Numpad0"},
	{97, "num1", "Numpad1"},
	{98, "num2", "Numpad2"},
	{99, "num3", "Numpad3"},
	{100, "num4", "Numpad4"},
	{101, "num5", "Numpad5"},
	{102, "num6", "Numpad6"},
	{103, "num7", "Numpad7"},
	{104, "num8", "Numpad8"},
	{105, "num9", "Numpad9"},
	{106, "nummult", "NumpadMultiply"},
	{107, "numadd", "NumpadAdd"},
	{109, "numsub", "NumpadSubtract"},
	{110, "numdec", "NumpadDecimal"},
	{111, "numdiv", "NumpadDivide"},
}

var modifiers = Keycodes{
	{162, "Ctrl", "ControlLeft"},
	{163, "RightCtrl", "ControlRight"},
	{160, "Shift", "ShiftLeft"},
	{161, "RightShift", "ShiftRight"},
	{164, "Alt", "AltLeft"},
	{165, "RightAlt", "AltRight"},
}

var arrows = Keycodes{
	{37, "Left", "ArrowLeft"},
	{38, "Up", "ArrowUp"},
	{39, "Right", "ArrowRight"},
	{40, "Down", "ArrowDown"},
}

var special = Keycodes{
	{8, "Backspace", "Backspace"},
	{9, "Tab", "Tab"},
	{20, "CapsLock", "CapsLock"},
	{27, "Escape", "Escape"},
	{32, "Space", "Space"},
	{91, "Windows", "Windows"},

	{186, ";", "Semicolon"},
	{187, "=", "Equal"},
	{188, ",", "Comma"},
	{189, "-", "Minus"},
	{190, ".", "Period"},
	{191, "/", "Slash"},
	{192, "`", "Backquote"},

	{219, "[", "BracketLeft"},
	{220, "\\", "Backslash"},
	{221, "]", "BracketRight"},
	{222, "'", "Quote"},
}

var fkeys = Keycodes{
	{112, "F1", "F1"},
	{113, "F2", "F2"},
	{114, "F3", "F3"},
	{115, "F4", "F4"},
	{116, "F5", "F5"},
	{117, "F6", "F6"},
	{118, "F7", "F7"},
	{119, "F8", "F8"},
	{120, "F9", "F9"},
	{121, "F10", "F10"},
	{122, "F11", "F11"},
	{123, "F12", "F12"},
}
