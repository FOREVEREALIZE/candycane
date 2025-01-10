package candycane

import (
	"fmt"
	"slices"
	"strings"
)

func singleHexToInt(h string) int {
	cleanCode := strings.ToUpper(h)

	if !slices.Contains(strings.Split("0123456789ABCDEF", ""), cleanCode) || len(cleanCode) > 1 {
		panic("Invalid hex code")
	}

	return map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}[cleanCode]
}

func intToAnyHex(n int, p int) string {
	return fmt.Sprintf(fmt.Sprintf("%%0%dX", p), n)
}

func doubleHexToInt(h string) int {
	for _, s := range strings.Split(h, "") {
		if !slices.Contains(strings.Split("0123456789ABCDEF", ""), s) {
			panic("Invalid hex code")
		}
	}

	if len(h) > 2 {
		panic("Invalid hex code")
	}

	return (singleHexToInt(h[0:1]) << 4) | singleHexToInt(h[1:2])
}
