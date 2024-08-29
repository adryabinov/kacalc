package formatter

import (
	"kacalc/romanconverter"
	"strconv"
)

func FormatResult(i int, format string) string {
	if format == "rom" {
		result := romanconverter.IntToRom(i)
		return result
	}
	if format == "dec" {
		result := strconv.Itoa(i)
		return result
	}
	return ""
}
