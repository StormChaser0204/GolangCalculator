package main

import (
	"fmt"
	"os"
	"time"
)

type PanicType uint8

const (
	InvalidNumber PanicType = iota
	InvalidOperation
	DifferentNumbers
	OutOfARange
	InvalidInput
	InvalidCalculation
)

func ShowPanicMessageAndCloseApplication(panicType PanicType) {
	var message = getMessage(panicType)
	fmt.Println(message)
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

func getMessage(panicType PanicType) string {
	var targetMessage string

	switch panicType {
	case InvalidNumber:
		targetMessage = "InvalidNumber: Can't convert input value to number"
	case InvalidOperation:
		targetMessage = "InvalidOperation: No expression"
	case DifferentNumbers:
		targetMessage = "DifferentNumbers: Input numbers have different type"
	case OutOfARange:
		targetMessage = "OutOfARange: Input numbers out of valid range"
	case InvalidInput:
		targetMessage = "InvalidInput: Wrong input number"
	case InvalidCalculation:
		targetMessage = "InvalidCalculation: Wrong result"
	}

	return targetMessage
}
