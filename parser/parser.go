package parser

import (
	"bufio"
	"errors"
	"kacalc/romanconverter"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLn() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	return input
}

func typeInput(s string) (string, error) {
	valid_dec, _ := regexp.MatchString(`^([1-9]|10)\s[+|-|*|/]\s([1-9]|10)$`, s)
	valid_rom, _ := regexp.MatchString(`^(I|II|III|IV|V|VI|VII|VIII|IX|X)\s[+|-|*|/]\s(I|II|III|IV|V|VI|VII|VIII|IX|X)$`, s)
	if valid_dec {
		return "dec", nil
	}
	if valid_rom {
		return "rom", nil
	}
	return "", errors.New("input error")
}

func ParseInput() (int, int, string, string) {
	var firstArg, secondArg int
	input := readLn()
	format, err := typeInput(input)
	if err != nil {
		log.Fatal(err)
	}
	splited := strings.Split(input, " ")
	operand := splited[1]
	switch format {
	case "dec":
		firstArg, err = strconv.Atoi(splited[0])
		if err != nil {
			log.Fatal(err)
		}
		secondArg, err = strconv.Atoi(splited[2])
		if err != nil {
			log.Fatal(err)
		}
	case "rom":
		firstArg = romanconverter.RomToInt(splited[0])
		secondArg = romanconverter.RomToInt(splited[2])
	}
	return firstArg, secondArg, operand, format
}
