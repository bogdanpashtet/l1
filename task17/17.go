// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func binarySearch(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if target == nums[middle] {
			return middle
		} else if right-left == 1 && target == nums[right] {
			return right
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	var sl = []int{1, 4, 8, 13, 16, 22, 27, 29, 32, 36, 37, 45, 53, 60, 64, 68, 73, 77, 82, 84, 86, 87, 90, 92, 93}
	findElem := 60
	res := binarySearch(sl, findElem)

	if res != -1 {
		fmt.Printf("Element %v exists in nums and its index is %v", findElem, res)
	} else {
		fmt.Printf("Element %v doesn't exists in nums.", findElem)
	}
}
