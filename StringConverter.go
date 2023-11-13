package main

import (
	"strconv"
	"strings"
)

var romeNumbers = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

func ConvertToExpression(input string) Expression {
	operationSymbols := [4]string{"+", "-", "*", "/"}
	var inputOperation string
	for i := 0; i < len(operationSymbols); i++ {
		if strings.Contains(input, operationSymbols[i]) {
			inputOperation = operationSymbols[i]
			break
		}
	}
	if inputOperation == "" {
		ShowPanicMessageAndCloseApplication(InvalidOperation)
	}

	var numbers = strings.Split(input, inputOperation)
	var firstNumber, firstNumberType = convertToInteger(numbers[0])
	var secondNumber, secondNumberType = convertToInteger(numbers[1])

	if firstNumberType != secondNumberType {
		ShowPanicMessageAndCloseApplication(DifferentNumbers)
	}

	var expression = Expression{
		firstNumber:  firstNumber,
		secondNumber: secondNumber,
		operation:    convertToOperation(inputOperation),
		numberType:   firstNumberType,
	}
	return expression
}

func convertToInteger(s string) (int8, NumberType) {

	numberType := Arabic
	var number int
	var err error
	var firstChar = s[0]
	if firstChar < 48 || firstChar > 57 {
		numberType = Rome
		number = romanToInteger(s)
	} else {
		number, err = strconv.Atoi(s)
	}

	if err != nil {
		ShowPanicMessageAndCloseApplication(InvalidNumber)
	}

	if number < 1 || number > 10 {
		ShowPanicMessageAndCloseApplication(OutOfARange)
	}
	return (int8)(number), numberType
}

func convertToOperation(s string) OperationType {
	var op OperationType
	switch s {
	case "+":
		op = Sum
	case "-":
		op = Sub
	case "*":
		op = Mul
	case "/":
		op = Div
	}
	return op
}

func romanToInteger(s string) int {
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := romeNumbers[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}

func integerToRoman(number int8) string {
	conversions := []struct {
		value int8
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
