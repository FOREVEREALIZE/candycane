package candycane

type Color struct {
	red   int
	green int
	blue  int
}

func (c Color) RGB() (int, int, int) {
	return c.red, c.green, c.blue
}

func (c Color) Red() int {
	return c.red
}

func (c Color) Green() int {
	return c.green
}

func (c Color) Blue() int {
	return c.blue
}

func (c Color) Hex() string {
	return intToAnyHex(c.red, 2) + intToAnyHex(c.green, 2) + intToAnyHex(c.blue, 2)
}

func (c Color) HexHash() string {
	return "#" + c.Hex()
}
