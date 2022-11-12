// Реализовать все возможные способы остановки выполнения горутины.

package main

import "fmt"

// Generator один канал для передачи данных и сигнализации
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {

	quit := make(chan struct{}) // сигнал остановки
	go func() {
		for {
			select {
			case <-quit:
				return
			default:

			}
		}
	}()
	close(quit)

	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)

	close(number) // close channel -> goroutine stops
}
