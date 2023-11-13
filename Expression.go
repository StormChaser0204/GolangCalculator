package main

type OperationType uint8

const (
	Sum OperationType = iota //Sum
	Sub                      //Subtract
	Mul                      //Multiply
	Div                      //Divide
)

type NumberType uint8

const (
	Arabic NumberType = iota
	Rome
)

type Expression struct {
	firstNumber, secondNumber int8
	operation                 OperationType
	numberType                NumberType
}
