package keycodes

import (
	"slices"
)

// combines all keycode groups
var allKeycodes Keycodes

// map of map of raw codes to strings
var rawcodeGroups map[KeycodeGroup]map[uint16]Keycode

// map of map of strings to raw codes
var displayGroups map[KeycodeGroup]map[string]Keycode

// map of map of strings to raw codes
var jsCodeGroups map[KeycodeGroup]map[string]Keycode

func init() {
	allKeycodes = slices.Concat(
		characters,
		numbers,
		numpad,
		modifiers,
		arrows,
		special,
		fkeys,
	)

	// build map of rawcode groups
	rawcodeGroups = map[KeycodeGroup]map[uint16]Keycode{
		KeycodeGroupAll:          allKeycodes.mapByRawcode(),
		KeycodeGroupCharacters:   characters.mapByRawcode(),
		KeycodeGroupNumbers:      numbers.mapByRawcode(),
		KeycodeGroupNumpad:       numpad.mapByRawcode(),
		KeycodeGroupModifiers:    modifiers.mapByRawcode(),
		KeycodeGroupArrows:       arrows.mapByRawcode(),
		KeycodeGroupSpecial:      special.mapByRawcode(),
		KeycodeGroupFunctionKeys: fkeys.mapByRawcode(),
	}

	// build map of string groups
	displayGroups = map[KeycodeGroup]map[string]Keycode{
		KeycodeGroupAll:          allKeycodes.mapByDisplay(),
		KeycodeGroupCharacters:   characters.mapByDisplay(),
		KeycodeGroupNumbers:      numbers.mapByDisplay(),
		KeycodeGroupNumpad:       numpad.mapByDisplay(),
		KeycodeGroupModifiers:    modifiers.mapByDisplay(),
		KeycodeGroupArrows:       arrows.mapByDisplay(),
		KeycodeGroupSpecial:      special.mapByDisplay(),
		KeycodeGroupFunctionKeys: fkeys.mapByDisplay(),
	}

	// build map of JSCode groups
	jsCodeGroups = map[KeycodeGroup]map[string]Keycode{
		KeycodeGroupAll:          allKeycodes.mapByJSCode(),
		KeycodeGroupCharacters:   characters.mapByJSCode(),
		KeycodeGroupNumbers:      numbers.mapByJSCode(),
		KeycodeGroupNumpad:       numpad.mapByJSCode(),
		KeycodeGroupModifiers:    modifiers.mapByJSCode(),
		KeycodeGroupArrows:       arrows.mapByJSCode(),
		KeycodeGroupSpecial:      special.mapByJSCode(),
		KeycodeGroupFunctionKeys: fkeys.mapByJSCode(),
	}
}

type KeycodeGroup int

const (
	KeycodeGroupAll = iota
	KeycodeGroupCharacters
	KeycodeGroupNumbers
	KeycodeGroupNumpad
	KeycodeGroupModifiers
	KeycodeGroupArrows
	KeycodeGroupSpecial
	KeycodeGroupFunctionKeys
)

var KeycodeGroups = []struct {
	Value  KeycodeGroup
	TSName string
}{
	{KeycodeGroupAll, "ALL"},
	{KeycodeGroupCharacters, "CHARACTERS"},
	{KeycodeGroupNumbers, "NUMBERS"},
	{KeycodeGroupNumpad, "NUMPAD"},
	{KeycodeGroupModifiers, "MODIFIERS"},
	{KeycodeGroupArrows, "ARROWS"},
	{KeycodeGroupSpecial, "SPECIAL"},
	{KeycodeGroupFunctionKeys, "FUNCTIONKEYS"},
}

type Keycode struct {
	Rawcode uint16
	Display string
	JSCode  string
}

// func (keycode Keycode) String() string {
// 	return keycode.Display
// }

func (keycode Keycode) Bind() Keycode {
	return keycode
}

type Keycodes []Keycode

func (keycodes Keycodes) mapByRawcode() map[uint16]Keycode {
	rawcodeMap := make(map[uint16]Keycode)
	for _, keycode := range keycodes {
		rawcodeMap[keycode.Rawcode] = keycode
	}

	return rawcodeMap
}

func (keycodes Keycodes) mapByDisplay() map[string]Keycode {
	displayMap := make(map[string]Keycode)
	for _, keycode := range keycodes {
		displayMap[keycode.Display] = keycode
	}

	return displayMap
}

func (keycodes Keycodes) mapByJSCode() map[string]Keycode {
	jsCodeMap := make(map[string]Keycode)
	for _, keycode := range keycodes {
		jsCodeMap[keycode.JSCode] = keycode
	}

	return jsCodeMap
}
