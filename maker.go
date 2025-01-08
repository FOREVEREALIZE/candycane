package candycane

import "strings"

func FromHex(code string) Color {
	cleanCode := strings.TrimPrefix(code, "#")
	if len(cleanCode) != 6 {
		panic("Invalid hex code")
	}

	return Color{
		red:   doubleHexToInt(cleanCode[0:2]),
		green: doubleHexToInt(cleanCode[2:4]),
		blue:  doubleHexToInt(cleanCode[4:6]),
	}
}
