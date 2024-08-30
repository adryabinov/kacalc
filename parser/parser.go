package parser

import (
	"bufio"
	"kacalc/converter"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLn() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	return input
}

func ParseInput() (int, int, string, string) {
	var firstArg, secondArg int
	var format1, format2 string
	var err error
	input := readLn()
	splited := strings.Split(input, " ")
	if len(splited) != 3 {
		log.Fatal("to many arguments")
	}
	operator := splited[1]
	switch operator {
	case "+":
	case "-":
	case "*":
	case "/":
	default:
		log.Fatal("invalid operator")
	}

	firstArg, err = strconv.Atoi(splited[0])
	if (err == nil) && (0 < firstArg) && (firstArg <= 10) {
		format1 = "dec"
	}
	if err != nil {
		firstArg = converter.RomToInt(splited[0])
		if firstArg != 0 {
			format1 = "rom"
		} else {
			log.Fatal("invalid first arg")
		}
	}
	secondArg, err = strconv.Atoi(splited[2])
	if (err == nil) && (0 < secondArg) && (secondArg <= 10) {
		format2 = "dec"
	}
	if err != nil {
		secondArg = converter.RomToInt(splited[2])
		if secondArg != 0 {
			format2 = "rom"
		} else {
			log.Fatal("invalid second arg")
		}
	}

	if format1 == format2 && format1 != "" {
		return firstArg, secondArg, operator, format1
	} else {
		log.Fatal("args must be one type")
	}
	return firstArg, secondArg, operator, format1
}
