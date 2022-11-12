// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		pivot := arr[0]
		var left, right, equal []int

		for i := 0; i < len(arr); i++ {
			if arr[i] > pivot {
				right = append(right, arr[i])
			} else if arr[i] < pivot {
				left = append(left, arr[i])
			} else {
				equal = append(equal, arr[i])
			}
		}
		return append(append(quickSort(left), equal...), quickSort(right)...)
	}
}

func main() {
	var arr = []int{99, 72, 53, 18, 40, 91, 26, 31, 71, 23, 57, 32, 19, 51, 68, 41, 79, 39, 45, 67, 49, 100, 59, 52, 24}
	fmt.Printf("Unsorted array:\t%v", arr)
	sl := arr[:]
	fmt.Printf("\nSorted array:\t%v", quickSort(sl))
}
