// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	var sl = []int{1, 4, 8, 13, 16, 22, 27, 29, 32, 36, 37, 45, 53, 60, 64, 68, 73, 77, 82, 84, 86, 87, 90, 92, 93}
	deleteElemIndex := 5 // удаляем 22

	slResult := make([]int, len(sl))
	copy(slResult, sl)
	slResult = append(slResult[:deleteElemIndex], slResult[deleteElemIndex+1:]...)

	fmt.Printf("Input slice:\t%v\nNew slice:\t%v", sl, slResult)
}
