// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func worker(channel chan int) {
	for j := range channel {
		fmt.Printf("%v, ", j)
	}
}

func main() {
	var ch = make(chan int)
	N := 10

	defer close(ch)

	go worker(ch)

	go func(channel chan int) {
		for i := 1; ; i++ {
			channel <- i
			time.Sleep(time.Second)
		}
	}(ch)

	time.Sleep(time.Duration(N) * time.Second)
}
