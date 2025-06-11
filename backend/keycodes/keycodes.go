package keycodes

// map of map of raw codes to strings
var rawcodeGroups map[KeycodeGroup]map[uint16]Keycode

// map of map of strings to raw codes
var displayGroups map[KeycodeGroup]map[string]Keycode

// map of map of strings to raw codes
var jsCodeGroups map[KeycodeGroup]map[string]Keycode

func init() {
	// build map of rawcode groups
	rawcodeGroups = mapKeycodeGroups(mapByRawcode)
	displayGroups = mapKeycodeGroups(mapByDisplay)
	jsCodeGroups = mapKeycodeGroups(mapByJSCode)
}

func mapKeycodeGroups[T comparable](mappingFunc func(Keycodes) map[T]Keycode) map[KeycodeGroup]map[T]Keycode {
	keycodeGroupsMap := make(map[KeycodeGroup]map[T]Keycode)
	for group, keycodes := range keycodeGroups {
		keycodeGroupsMap[group] = mappingFunc(keycodes)
	}

	return keycodeGroupsMap
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

func mapByRawcode(keycodes Keycodes) map[uint16]Keycode {
	rawcodeMap := make(map[uint16]Keycode)
	for _, keycode := range keycodes {
		rawcodeMap[keycode.Rawcode] = keycode
	}

	return rawcodeMap
}

func mapByDisplay(keycodes Keycodes) map[string]Keycode {
	displayMap := make(map[string]Keycode)
	for _, keycode := range keycodes {
		displayMap[keycode.Display] = keycode
	}

	return displayMap
}

func mapByJSCode(keycodes Keycodes) map[string]Keycode {
	jsCodeMap := make(map[string]Keycode)
	for _, keycode := range keycodes {
		jsCodeMap[keycode.JSCode] = keycode
	}

	return jsCodeMap
}
