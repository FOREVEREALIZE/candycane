package candycane

import "strings"

// FromHex takes a hex code as a string and returns a Color that represents it
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

// FromRGB takes RGB values as three ints and returns a Color that represent them
func FromRGB(red int, green int, blue int) Color {
	return Color{
		red:   red,
		green: green,
		blue:  blue,
	}
}
