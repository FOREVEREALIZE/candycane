package candycane

// Color represents a color. Direct access to color values should be done via methods.
type Color struct {
	red   int
	green int
	blue  int
}

// RGB returns three ints representing the red, green and blue channels of the color
func (c Color) RGB() (int, int, int) {
	return c.red, c.green, c.blue
}

// Red returns an int representing the red channel of the color
func (c Color) Red() int {
	return c.red
}

// Green returns an int representing the green channel of the color
func (c Color) Green() int {
	return c.green
}

// Blue returns an int representing the blue channel of the color
func (c Color) Blue() int {
	return c.blue
}

// Hex returns a string representing the hex color code of the color in the format "RRGGBB"
func (c Color) Hex() string {
	return intToAnyHex(c.red, 2) + intToAnyHex(c.green, 2) + intToAnyHex(c.blue, 2)
}

// HexHash returns a string representing the hex color code of the color in the format "#RRGGBB"
func (c Color) HexHash() string {
	return "#" + c.Hex()
}
