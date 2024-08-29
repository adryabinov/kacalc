package calculator

import "errors"

func Calculate(a int, b int, operand string, format string) (int, string, error) {
	var result int
	var err error
	switch operand {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	if format == "rom" {
		if result <= 0 {
			err = errors.New("rome must be >= 0")
		}
	}
	return result, format, err
}
