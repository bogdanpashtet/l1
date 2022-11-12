// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup
func method1(mass []int) {
	var w sync.WaitGroup

	for i := range mass {
		w.Add(1)
		go func(value int) {
			fmt.Printf("%v, ", value*value)
			w.Done()
		}(mass[i])
	}
	w.Wait()
}

// Time sleep
func method2(mass []int) {
	for i := range mass {
		go func(value int) {
			fmt.Printf("%v, ", value*value)
		}(mass[i])
	}
	time.Sleep(100 * time.Millisecond)
}

// Buffered Channel
func method3(mass []int) {
	var ch = make(chan int, 5)

	defer close(ch)

	for i := range mass {
		go func(value int) {
			ch <- value * value
		}(mass[i])
	}
	for j := 0; j < len(mass); j++ {
		fmt.Printf("%v, ", <-ch)
	}
}

func main() {
	var arr = [...]int{2, 4, 6, 8, 10}
	sl := arr[:]
	fmt.Printf("WaitGroup: \t\t")
	method1(sl)
	fmt.Printf("\nTime Sleep 100ms: \t")
	method2(sl)
	fmt.Printf("\nBuffered channel: \t")
	method3(sl)
}
