// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func main() {
	var val int64 = 333232
	numBit := 4
	var inBinary int64 = 1 << numBit
	fmt.Printf("Number: %v (%#b)\nNumber of bit: %v (%#b)\n", val, val, numBit, inBinary)

	res := val ^ inBinary
	fmt.Printf("Result: %v (%#b)\n", res, res)
}
