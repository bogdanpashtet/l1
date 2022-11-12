// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func checkUnique(runeArr []rune) bool {
	var m = make(map[rune]struct{})

	for i := range runeArr {
		if _, ok := m[runeArr[i]]; !ok {
			m[runeArr[i]] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	var str = "abcdefjhigklmnoprstuvwxyz"
	str = strings.ToLower(str)
	fmt.Printf("Input string: %s\nResult: %v", str, checkUnique([]rune(str)))
}
