package main

import (
	"fmt"
	"kacalc/calculator"
	"kacalc/formatter"
	"kacalc/parser"
	"log"
)

func main() {
	result, format, err := calculator.Calculate(parser.ParseInput())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(formatter.FormatResult(result, format))
}
