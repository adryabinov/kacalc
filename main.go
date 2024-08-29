package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func romToInt(rom string) int {
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

func intToRom(i int) string {
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

func main() {
	var result string
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	valid_dec, _ := regexp.MatchString("([123456789]|10) [+-/*] ([123456789]|10)", input)
	valid_rom, _ := regexp.MatchString("(I|II|III|IV|V|VI|VII|VIII|IX|X) [+-/*] (I|II|III|IV|V|VI|VII|VIII|IX|X)", input)
	switch {
	case valid_dec:
		fmt.Println("dec")
		splited := strings.Split(input, " ")
		operand := splited[1]
		firstArg, err := strconv.Atoi(splited[0])
		if err != nil {
			log.Fatal(err)
		}
		secondArg, err := strconv.Atoi(splited[2])
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case operand == "+":
			result = strconv.Itoa(firstArg + secondArg)
		case operand == "-":
			result = strconv.Itoa(firstArg - secondArg)
		case operand == "*":
			result = strconv.Itoa(firstArg * secondArg)
		case operand == "/":
			result = strconv.Itoa(firstArg / secondArg)
		}
		fmt.Println(result)
	case valid_rom:
		fmt.Println("rom")
		splited := strings.Split(input, " ")
		operand := splited[1]
		firstArg := romToInt(splited[0])
		secondArg := romToInt(splited[2])
		switch {
		case operand == "+":
			result = intToRom(firstArg + secondArg)
		case operand == "-":
			result = intToRom(firstArg - secondArg)
		case operand == "*":
			result = intToRom(firstArg * secondArg)
		case operand == "/":
			result = intToRom(firstArg / secondArg)
		}
		fmt.Println(result)
	default:
		fmt.Println("INPUT ERR")
	}

}
