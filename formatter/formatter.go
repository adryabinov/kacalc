package formatter

import (
	"kacalc/converter"
	"strconv"
)

func FormatResult(i int, format string) string {
	if format == "rom" {
		result := converter.IntToRom(i)
		return result
	}
	if format == "dec" {
		result := strconv.Itoa(i)
		return result
	}
	return ""
}
