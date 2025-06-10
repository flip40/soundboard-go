package keycodes

func RawcodeToString(rawcode uint16) string {
	return rawcodes[rawcode]
}

func StringToRawcode(str string) uint16 {
	return rawcodeStrings[str]
}

// FOR BINDING TO WAILS
type KeycodeHelper struct{}

func (helper KeycodeHelper) RawcodeToString(rawcode uint16) string {
	return RawcodeToString(rawcode)
}

func (helper KeycodeHelper) StringToRawcode(str string) uint16 {
	return StringToRawcode(str)
}

func (helper KeycodeHelper) GetKeycodes() map[uint16]string {
	return rawcodes
}
