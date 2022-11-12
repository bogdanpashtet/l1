// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {

	// первое число из строки
	firstBigNumber := new(big.Int)
	firstBigNumber.SetString("111111111111111111111111111111111111111111111111111111111111111111", 10)

	// второе число из строки
	secondBigNumber := new(big.Int)
	secondBigNumber.SetString("66666666666666666666666666666666666666666666666666666666", 10)

	// умножение
	multiplication := new(big.Int)
	multiplication.Mul(firstBigNumber, secondBigNumber)

	// деление
	division := new(big.Int)
	division.Div(firstBigNumber, secondBigNumber)

	// сложение
	addition := new(big.Int)
	addition.Add(firstBigNumber, secondBigNumber)

	// вычитание
	subtraction := new(big.Int)
	subtraction.Sub(firstBigNumber, secondBigNumber)

	fmt.Printf("First number:\t%s\nSecond number:\t%s\nMultiplication:\t%s\nDivision:\t%s\nAddition:\t%s\nSubtraction:\t%s",
		firstBigNumber, secondBigNumber, multiplication, division, addition, subtraction)
}
