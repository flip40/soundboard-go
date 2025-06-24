package keycodes

func RawcodeToString(rawcode uint16) string {
	return rawcodeGroups[KeycodeGroupAll][rawcode].Display
}

func StringToRawcode(str string) uint16 {
	return displayGroups[KeycodeGroupAll][str].Rawcode
}

func JSCodeToRawcode(jsCode string) uint16 {
	return jsCodeGroups[KeycodeGroupAll][jsCode].Rawcode
}

// FOR BINDING TO WAILS
type KeycodeHelper struct{}

func (helper KeycodeHelper) GetRawcodeGroups() map[KeycodeGroup]map[uint16]Keycode {
	return rawcodeGroups
}

func (helper KeycodeHelper) GetDisplayGroups() map[KeycodeGroup]map[string]Keycode {
	return displayGroups
}

func (helper KeycodeHelper) GetJSCodeGroups() map[KeycodeGroup]map[string]Keycode {
	return jsCodeGroups
}
