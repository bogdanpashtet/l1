// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Разработать программу, которая переворачивает слова в строке."
	strArray := strings.Fields(str)

	// reverse elements in string array
	for left, right := 0, len(strArray)-1; left < right; left, right = left+1, right-1 {
		strArray[left], strArray[right] = strArray[right], strArray[left]
	}

	fmt.Printf("Input string: %s\nResult string: %s", str, strings.Join(strArray, " "))
}
