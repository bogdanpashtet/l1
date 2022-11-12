// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	firstNumber, secondNumber := 4, 2
	fmt.Printf("First varriable value: %v\tSecond varriable value: %v\n", firstNumber, secondNumber)
	fmt.Println("Switch variables...")
	secondNumber, firstNumber = firstNumber, secondNumber
	fmt.Printf("First varriable value: %v\tSecond varriable value: %v", firstNumber, secondNumber)
}
