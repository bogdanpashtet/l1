// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// WaitGroup mutex
func method1(mass []int) int {
	var w sync.WaitGroup
	var m sync.Mutex
	var summ int

	for i := range mass {
		w.Add(1)
		go func(value int) {
			m.Lock()
			summ += value * value
			m.Unlock()
			w.Done()
		}(mass[i])
	}
	w.Wait()

	return summ
}

// Time sleep
func method2(mass []int) int {
	var m sync.Mutex
	var summ int
	for i := range mass {
		go func(value int) {
			m.Lock()
			summ += value * value
			m.Unlock()
		}(mass[i])
	}
	time.Sleep(100 * time.Millisecond)
	return summ
}

// Buffered Channel
func method3(mass []int) int {
	var ch = make(chan int, 5)
	var summ int

	defer close(ch)

	for i := range mass {
		go func(value int) {
			ch <- value * value
		}(mass[i])
		summ += <-ch
	}

	return summ
}

// WaitGroup atomic
func method4(mass []int) int {
	var w sync.WaitGroup
	var summ int64
	var summP *int64

	summP = &summ

	for i := range mass {
		w.Add(1)
		go func(value int) {
			atomic.AddInt64(summP, int64(value*value))
			w.Done()
		}(mass[i])
	}
	w.Wait()

	return int(summ)
}

// Time sleep
func method5(mass []int) int {
	var summ int64
	var summP *int64

	summP = &summ

	for i := range mass {
		go func(value int) {
			atomic.AddInt64(summP, int64(value*value))
		}(mass[i])
	}
	time.Sleep(100 * time.Millisecond)

	return int(summ)
}

func main() {
	var arr = [...]int{2, 4, 6, 8, 10}
	sl := arr[:]
	fmt.Printf("WaitGroup mutex: \t %v\n", method1(sl))
	fmt.Printf("TimeSleep 100ms: \t %v\n", method2(sl))
	fmt.Printf("Buffered channel: \t %v\n", method3(sl))
	fmt.Printf("WaitGroup atomic: \t %v\n", method4(sl))
	fmt.Printf("TimeSleep atomic: \t %v\n", method5(sl))
}
