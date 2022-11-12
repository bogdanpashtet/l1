// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
)

func writeToFirstChannel(arr []int, ch chan int) {
	defer close(ch)
	for i := range arr {
		ch <- arr[i]
	}
}

func writeToSecondChannel(ch1, ch2 chan int) {
	defer close(ch2)
	for i := range ch1 {
		ch2 <- i * 2
	}
}

func main() {
	var arr = [7]int{3, 4, 5, 6, 7, 8, 9}
	sl := arr[:]
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go writeToFirstChannel(sl, ch1)

	go writeToSecondChannel(ch1, ch2)

	for i := range ch2 {
		fmt.Printf("%v, ", i)
	}
}
