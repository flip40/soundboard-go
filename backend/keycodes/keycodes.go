package keycodes

func init() {
	rawcodeStrings = make(map[string]uint16)
	for k, v := range rawcodes {
		rawcodeStrings[v] = k
	}
}

var rawcodeStrings map[string]uint16

// map of raw codes to strings
var rawcodes = map[uint16]string{
	// characters
	65: "a",
	66: "b",
	67: "c",
	68: "d",
	69: "e",
	70: "f",
	71: "g",
	72: "h",
	73: "i",
	74: "j",
	75: "k",
	76: "l",
	77: "m",
	78: "n",
	79: "o",
	80: "p",
	81: "q",
	82: "r",
	83: "s",
	84: "t",
	85: "u",
	86: "v",
	87: "w",
	88: "x",
	89: "y",
	90: "z",

	// numbers list
	49:  "1",
	50:  "2",
	51:  "3",
	52:  "4",
	53:  "5",
	54:  "6",
	55:  "7",
	56:  "8",
	57:  "9",
	48:  "0",
	189: "-",
	187: "=",

	// numpad
	96:  "0",
	97:  "1",
	98:  "2",
	99:  "3",
	100: "4",
	101: "5",
	102: "6",
	103: "7",
	104: "8",
	105: "9",
	106: "*",
	107: "+",
	109: "-",
	110: ".",
	111: "/",

	// modifiers
	162: "ctrl",
	163: "rctl",
	160: "shift",
	161: "rshift",
	164: "alt",
	165: "ralt",

	// arrows
	37: "left",
	38: "up",
	39: "right",
	40: "down",

	// special
	8:  "backspace",
	9:  "tab",
	20: "caps lock",
	27: "escape",
	32: "space",
	91: "windows",

	186: ";",
	188: ",",
	190: ".",
	191: "/",
	192: "`",

	219: "[",
	220: "\\",
	221: "]",
	222: "'",

	// function keys
	112: "F1",
	113: "F2",
	114: "F3",
	115: "F4",
	116: "F5",
	117: "F6",
	118: "F7",
	119: "F8",
	120: "F9",
	121: "F10",
	122: "F11",
	123: "F12",
}
