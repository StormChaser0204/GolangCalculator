package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput = getInputFromConsole()
	var expression = ConvertToExpression(userInput)
	var result = calculate(expression)
	printResult(result, expression.numberType)
}

func getInputFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter expression")
	text, err := reader.ReadString('\n')
	if err != nil {
		ShowPanicMessageAndCloseApplication(InvalidInput)
	}

	text = strings.TrimSpace(text)
	return text
}

func calculate(expression Expression) int8 {
	var result int8

	switch expression.operation {
	case Sum:
		result = expression.firstNumber + expression.secondNumber
	case Sub:
		result = expression.firstNumber - expression.secondNumber
	case Mul:
		result = expression.firstNumber * expression.secondNumber
	case Div:
		result = expression.firstNumber / expression.secondNumber
	}

	if expression.numberType == Rome && result <= 0 {
		ShowPanicMessageAndCloseApplication(InvalidCalculation)
	}

	return result
}

func printResult(result int8, numberType NumberType) {
	if numberType == Arabic {
		fmt.Println(result)
	} else {
		fmt.Println(integerToRoman(result))
	}
}
