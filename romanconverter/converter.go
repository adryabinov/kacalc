package romanconverter

func RomToInt(rom string) int {
	mapRomToInt := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10}
	return mapRomToInt[rom]
}

func IntToRom(i int) string {
	result := ""
	for i >= 100 {
		result += "C"
		i -= 100
	}
	for i >= 90 {
		result += "XC"
		i -= 90
	}
	for i >= 50 {
		result += "L"
		i -= 50
	}
	for i >= 40 {
		result += "XL"
		i -= 40
	}
	for i >= 10 {
		result += "X"
		i -= 10
	}
	for i >= 9 {
		result += "IX"
		i -= 9
	}
	for i >= 5 {
		result += "V"
		i -= 5
	}
	for i >= 4 {
		result += "IV"
		i -= 4
	}
	for i >= 1 {
		result += "I"
		i -= 1
	}
	return result
}
